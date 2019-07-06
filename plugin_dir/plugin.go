package main

import (
	"fmt"
	"github.com/tgo-team/tgo-talkapi/cmd"
	"github.com/tgo-team/tgo-talkapi/plugin"
)

func Setup(context *plugin.Context) {
	fmt.Println("testzzz=",context)

	context.RegisterHandler(Test{})

}

type Test struct {

}


func (t Test) SayHello(context *cmd.Context)  {
     fmt.Println("SayHello")
	fmt.Println(context)
}
