package ctrl

import (
	"fmt"
	"github.com/tgo-team/tgo-core/tgo"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"github.com/tgo-team/tgo-talkapi/cmd"
	"github.com/tgo-team/tgo-talkapi/config"
)

type Controller struct {
	cmdManager *cmd.Manager
	handlerMap map[string]cmd.Handler
	tgo        *tgo.TGO
	waitGroup  tgo.WaitGroupWrapper
	config *config.Config
}

func New(tgo *tgo.TGO,config *config.Config) *Controller {
	c := &Controller{
		tgo:        tgo,
		cmdManager: cmd.NewManager(),
		handlerMap: map[string]cmd.Handler{},
		config: config,
	}
	c.waitGroup.Wrap(c.cmdLoop)
	return c
}

// loop队列里的cmd 一个个执行（TODO 为了提高性能这里可以开启多个loop，后面优化）
func (c *Controller) cmdLoop() {
	for {
		cmdObj := c.cmdManager.PopCMD()
		handler := c.handlerMap[cmdObj.Cmd]
		if handler != nil {
			c.waitGroup.Add(1)
			go func(cmdObj *cmd.CMD) {
				handler(cmd.NewContext(cmdObj))
				c.waitGroup.Done()
			}(cmdObj)
		}
	}
}

func (c *Controller) Start() {
	c.tgo.Match(fmt.Sprintf("type:%d", packets.Cmd), c.server)
}

// RegisterCMDHandler 注册cmd处理者
func (c *Controller) RegisterHandlerFunc(cmd string, handler cmd.Handler) {
	c.handlerMap[cmd] = handler
}

func (c *Controller) server(m *tgo.MContext) {
	cmdStr := m.CmdPacket().CMD
	c.cmdManager.PushCMD(cmd.NewCMD(cmdStr,m.CmdPacket().Payload,m))
}


type Handlers interface {
	RegisterHandler(c *Controller)
}