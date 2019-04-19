package cmd

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/tgo-team/tgo-core/tgo"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"github.com/tgo-team/tgo-talkapi/cache"
	"github.com/tgo-team/tgo-talkapi/config"
	"github.com/tgo-team/tgo-talkapi/utils"
	"strings"
)

type CMD struct {
	Cmd   string
	Param []byte
	m *tgo.MContext
}

func NewCMD(cmd string,param []byte,m *tgo.MContext) *CMD  {
	c := &CMD{
		Cmd:cmd,
		Param:param,
		m: m,
	}
	return c
}

/// Manager 命令管理者
type Manager struct {
	cmdChan chan *CMD
}

func NewManager() *Manager {

	return &Manager{
		cmdChan: make(chan *CMD, 1024), //TODO 决定了系统的执行命令的最大并发数
	}
}

func (m *Manager) PushCMD(cmd *CMD) {
	m.cmdChan <- cmd
}

func (m *Manager) PopCMD() *CMD {
	cmd := <-m.cmdChan
	return cmd
}

// 命令处理者
type Handler func(ctx *Context)

type Context struct {
	cmd *CMD
	tgo.Log
	Cache cache.Cache
	cfg *config.Config
	currentOpenId string
	isAbort bool // 是否中断
}

func NewContext(cmd *CMD,cache cache.Cache,cfg *config.Config) *Context  {
	return &Context{cmd:cmd,Log:cmd.m,Cache:cache,cfg:cfg}
}

func (c *Context) Param() []byte {
	return c.cmd.Param
}

func (c *Context) ReplySuccess(data []byte)  {
	c.cmd.m.ReplyPacket(packets.NewCmdackPacket(c.cmd.Cmd,SUCCESS,data))
}

func (c *Context) ReplyError(status uint16,msg string)  {
	errObj := &Error{Status:int32(status),Message:msg}
	errorBytes,err := proto.Marshal(errObj)
	if err!=nil {
		panic(err)
	}
	c.cmd.m.ReplyPacket(packets.NewCmdackPacket(c.cmd.Cmd,ERROR,errorBytes))
}
func (c *Context) ReplyErrorMsg(msg string)  {
	c.ReplyError(ERROR,msg)
}

func (c *Context) Auth() (string,error)  {
	 token := c.cmd.m.CmdPacket().Token
	 if strings.TrimSpace(token) == "" {
	 	return "",errors.New("token不存在！")
	 }
	 if c.Cache==nil {
	 	return "",errors.New("不存在缓存对象！")
	 }
	 openId,err := c.Cache.Get(fmt.Sprintf("%s%s",c.cfg.CachePrefix.TokenPrefix,token))
	 if err!=nil {
	 	return "",fmt.Errorf("获取token出错！-> %v",err)
	 }
	 if openId == "" {
	 	return "",fmt.Errorf("token不存在或已失效！")
	 }
	 return openId,nil
}

func (c *Context) SetCurrentOpenId(openId string)   {
	c.currentOpenId = openId
}

func (c *Context) CurrentOpenId()  string {

	return c.currentOpenId
}

func (c *Context) Abort()  {
	c.isAbort = true
}

func (c *Context) IsAbort() bool  {
	return c.isAbort
}

func (c *Context) MarshalProto(message proto.Message) []byte  {
	respData,err := proto.Marshal(message)
	utils.CheckErr(err)
	return respData
}

func (c *Context) UnmarshalProto(data []byte,pb proto.Message) error {

	return proto.Unmarshal(data,pb)
}

const (
	SUCCESS uint16 = 1
	ERROR uint16 = 0
)

