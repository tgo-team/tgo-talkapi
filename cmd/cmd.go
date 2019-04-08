package cmd

import (
	"github.com/golang/protobuf/proto"
	"github.com/tgo-team/tgo-core/tgo"
	"github.com/tgo-team/tgo-core/tgo/packets"
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

}

func NewContext(cmd *CMD) *Context  {
	return &Context{cmd:cmd,Log:cmd.m}
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

const (
	SUCCESS uint16 = 1
	ERROR uint16 = 0
)

