package user

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/tgo-team/tgo-talkapi/cmd"
	"github.com/tgo-team/tgo-talkapi/config"
	"github.com/tgo-team/tgo-talkapi/constant"
	"github.com/tgo-team/tgo-talkapi/ctrl"
	"github.com/tgo-team/tgo-talkapi/handler/db"
	"github.com/tgo-team/tgo-talkapi/utils"
	"github.com/tgo-team/tgo-talkapi/utils/network"
	"net/http"
	"strings"
)

type Handler struct {
	idFactory *utils.GuidFactory
	cfg       *config.Config
	db *db.DB
	dao Dao
}

func NewHandler(cfg *config.Config) *Handler {
	h := &Handler{
		cfg:       cfg,
		idFactory: utils.NewGUIDFactory(cfg.NodeId),
		dao: NewDefaultDao(),
	}
	return h
}

// Login 用户登录
func (h *Handler) Login(c *cmd.Context) {
	c.Info("this is login")

	loginReq := &UserLoginReq{}
	err := proto.Unmarshal(c.Param(), loginReq)
	if err != nil {
		panic(err)
	}
	// 查询登录用户信息
	userModel,err := h.dao.QueryUserWithUsername(loginReq.Username)
	if userModel==nil {
		c.Error("用户不存在！")
		c.ReplyErrorMsg("用户不存在！")
		return
	}

	if userModel.Password != loginReq.Password {
		c.Error("用户密码不正确！")
		c.ReplyErrorMsg("用户密码不正确！")
		return
	}

	token := utils.GenerUUId()

	// 请求talk服务器更新客户端信息
	if h.cfg.TalkHttpUrl != "" {
		talkClientUpdateUrl := fmt.Sprintf("%s%s", h.cfg.TalkHttpUrl, "/client/update")
		err = h.requestTalkUpdateClient(talkClientUpdateUrl, token, userModel.TalkId)
		if err != nil {
			c.Error("请求talk服务器出错！-> %v", err)
			c.ReplyErrorMsg("请求talk服务器出错！")
			return
		}
	}

	// 设置token
	err = c.Cache.Set(fmt.Sprintf("%s%s",h.cfg.CachePrefix.TokenPrefix,token),userModel.OpenId,h.cfg.TokenExpire.Duration)
	if err!=nil {
		c.Error("设置token失败！-> %v", err)
		c.ReplyErrorMsg("设置token失败！")
		return
	}
	loginResp := &UserLoginResp{
		OpenId:   userModel.OpenId,
		TalkId:   userModel.TalkId,
		Token:    token,
		Username: userModel.Username,
		Nickname: userModel.Nickname,
		Sex:      int32(userModel.Sex),
		Zone:     userModel.Zone,
		Mobile:   userModel.Mobile,
	}
	respData, _ := proto.Marshal(loginResp)

	c.ReplySuccess(respData)
}

// Register 注册用户
func (h *Handler) Register(c *cmd.Context) {
	registerReq := &RegisterReq{}
	err := proto.Unmarshal(c.Param(), registerReq)
	if err != nil {
		panic(err)
	}
	if err := h.checkRegister(registerReq); err != nil {
		c.ReplyErrorMsg(err.Error())
		return
	}
	// 查询登录用户信息
	userModel,err := h.dao.QueryUserWithUsername(registerReq.Username)
	if err!=nil {
		c.Error("查询用户数据出错！-> %v",err)
		c.ReplyErrorMsg("查询用户数据出错！")
		return
	}
	if userModel!=nil {
		c.Error("用户已存在！")
		c.ReplyErrorMsg("用户已存在！")
		return
	}
	talkID,err := h.idFactory.NewGUID()
	if err!=nil {
		c.Error("生成talkID出错！-> %v",err)
		c.ReplyErrorMsg("生成talkID出错！")
		return
	}

	userModel.OpenId = utils.GenerUUId()
	userModel.TalkId = uint64(talkID)
	userModel.Zone = registerReq.Zone
	userModel.Mobile = registerReq.Mobile
	userModel.Username = registerReq.Username
	userModel.Nickname = registerReq.Nickname
	userModel.Password = registerReq.Password
	userModel.Sex = int(registerReq.Sex)
	err = h.dao.InsertUser(userModel)
	if err!=nil {
		c.Error("添加用户信息出错！-> %v",err)
		c.ReplyErrorMsg("添加用户信息出错！")
		return
	}
}

// 获取单个用户信息
func (h *Handler) GetUser(c *cmd.Context) {
	userQueryReq := &UserQueryReq{}
	err := c.UnmarshalProto(c.Param(), userQueryReq)
	if err != nil {
		c.Error("数据格式有误！-> %v",err)
		c.ReplyErrorMsg("数据格式有误！")
		return
	}
	user,err := h.dao.QueryUserWithOption(userQueryReq.TalkId,userQueryReq.OpenId)
	if err!=nil {
		c.Error("查询用户数据出错！-> %v",err)
		c.ReplyErrorMsg("查询用户数据出错！")
		return
	}

	if user == nil {
		c.Error("用户不存在！-> %v",userQueryReq)
		c.ReplyError(constant.ERROR_USER_NO_EXIST,"用户不存在！")
		return
	}
	userVo := h.userToVo(user)
	c.ReplySuccess(c.MarshalProto(userVo))
}

func (h *Handler) checkRegister(registerReq *RegisterReq) error {
	if strings.TrimSpace(registerReq.Username) == "" && strings.TrimSpace(registerReq.Mobile) == "" {
		return fmt.Errorf("用户名不能为空！")
	}
	if strings.TrimSpace(registerReq.Password) == "" {
		return fmt.Errorf("密码不能为空！")
	}
	if len(strings.TrimSpace(registerReq.Password)) < 6 {
		return fmt.Errorf("密码长度不能小于6位！")
	}
	return nil
}

// 请求更新talk的客户端信息
func (h *Handler) requestTalkUpdateClient(talkClientUpdateUrl string, token string, talkID uint64) error {
	resp, err := network.Post(talkClientUpdateUrl, []byte(utils.ToJson(map[string]interface{}{
		"client_id": talkID,
		"password":  token,
	})), nil)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		var result struct {
			Status int    `json:"status"`
			Msg    string `json:"msg"`
		}
		err = utils.ReadJsonByByte([]byte(resp.Body), &result)
		if err != nil {
			return err
		}
		if result.Status != 200 {
			return fmt.Errorf(result.Msg)
		}
	} else {
		return fmt.Errorf("http状态码[%d]错误！", resp.StatusCode)
	}
	return nil
}

func (h *Handler) RegisterHandler(c *ctrl.Controller) {
	c.RegisterHandlerFuncs("login", h.Login)
	c.RegisterHandlerFuncs("register", h.Register)
	c.RegisterHandlerFuncs("get_user", h.GetUser)
}



func (h *Handler) userToVo(user *User) *UserVo {
	userVo := &UserVo{}
	userVo.OpenId = user.OpenId
	userVo.TalkId = user.TalkId
	userVo.Nickname = user.Nickname
	userVo.Sex = int32(user.Sex)
	return userVo
}
