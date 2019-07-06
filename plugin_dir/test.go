package main

import (
	"github.com/tgo-team/tgo-talkapi/plugin"
	"github.com/tgo-team/tgo-talkapi/utils"
	pn "plugin"
)

func main()  {
	p,err := pn.Open("./plugin_dir/plugin.so")
	utils.CheckErr(err)

	sb, err := p.Lookup("Setup")

	sbs := sb.(func(ctx *plugin.Context))

	sbs(&plugin.Context{Name:"ddddd"})
}
