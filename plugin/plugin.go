package plugin

import (
	"fmt"
	"github.com/tgo-team/tgo-core/tgo"
	"github.com/tgo-team/tgo-talkapi/cmd"
	"github.com/tgo-team/tgo-talkapi/ctrl"
	"reflect"
)

type Context struct {
	Tgo *tgo.TGO
	controller *ctrl.Controller
}

func NewContext(controller *ctrl.Controller,tgo *tgo.TGO) *Context {

	return &Context{controller:controller,Tgo:tgo}
}

func (c *Context) RegisterHandler(handler interface{})  {

	var objType = reflect.TypeOf(handler)
	var objValue = reflect.ValueOf(handler)

	for i:=0;i<objType.NumMethod();i++ {
		//m := objType.Method(i)
		//fmt.Printf("%s: %v\n", m.Name, m.Type)
		var m = objValue.Method(i)
		args := []reflect.Value{reflect.ValueOf(&cmd.Context{})}
		m.Call(args)

		//c.controller.RegisterHandlerFuncs(m.Name,)
	}

	fmt.Println("get all Fields is:", objType)
}