package ctrl

import (
	"fmt"
	"github.com/tgo-team/tgo-core/tgo"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"github.com/tgo-team/tgo-talkapi/cache"
	"github.com/tgo-team/tgo-talkapi/cmd"
	"github.com/tgo-team/tgo-talkapi/config"
)

type Controller struct {
	cmdManager *cmd.Manager
	handlerMap map[string][]cmd.Handler
	tgo        *tgo.TGO
	waitGroup  tgo.WaitGroupWrapper
	config *config.Config
	Cache cache.Cache
}

func New(tgo *tgo.TGO,config *config.Config,cache cache.Cache) *Controller {
	c := &Controller{
		tgo:        tgo,
		cmdManager: cmd.NewManager(),
		handlerMap: map[string][]cmd.Handler{},
		config: config,
		Cache: cache,
	}
	c.waitGroup.Wrap(c.cmdLoop)
	return c
}

// loop队列里的cmd 一个个执行（TODO 为了提高性能这里可以开启多个loop，后面优化）
func (c *Controller) cmdLoop() {
	for {
		cmdObj := c.cmdManager.PopCMD()
		handlers := c.handlerMap[cmdObj.Cmd]
		if handlers != nil && len(handlers)>0 {
			c.waitGroup.Add(1)
			go func(cmdObj *cmd.CMD) {
				cmdContext := cmd.NewContext(cmdObj,c.Cache,c.config)
				for _,handler :=range handlers {
					if !cmdContext.IsAbort() {
						handler(cmdContext)
					}
				}
				c.waitGroup.Done()
			}(cmdObj)
		}
	}
}

func (c *Controller) Start() {
	c.tgo.Match(fmt.Sprintf("type:%d", packets.Cmd), c.server)
}

// RegisterCMDHandler 注册cmd处理者
func (c *Controller) RegisterHandlerFuncs(cmd string, handlers... cmd.Handler) {
	c.handlerMap[cmd] = handlers
}

func (c *Controller) server(m *tgo.MContext) {
	cmdStr := m.CmdPacket().CMD
	c.cmdManager.PushCMD(cmd.NewCMD(cmdStr,m.CmdPacket().Payload,m))
}


type Handlers interface {
	RegisterHandler(c *Controller)
}