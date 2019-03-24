package main

import (
	"github.com/judwhite/go-svc/svc"
	"github.com/tgo-team/tgo-core/tgo"
	_ "github.com/tgo-team/tgo-talkapi/tgo/log"
	_ "github.com/tgo-team/tgo-talkapi/tgo/server/tcp"
	_ "github.com/tgo-team/tgo-talkapi/tgo/storage/memory"
	"github.com/tgo-team/tgo-talkapi/user"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		panic(err)
	}
}

type program struct {
	t *tgo.TGO
}

func (p *program) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func (p *program) Start() error {

	t := tgo.New(tgo.NewOptions())
	user.StartUser(t)
	p.t = t
	err := t.Start()
	if err != nil {
		panic(err)
	}

	return nil
}

func (p *program) Stop() error {
	if p.t != nil {
		return p.t.Stop()

	}
	return nil
}
