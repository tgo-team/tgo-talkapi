package main

import (
	"github.com/go-redis/redis"
	"github.com/judwhite/go-svc/svc"
	"github.com/rubenv/sql-migrate"
	"github.com/tgo-team/tgo-core/tgo"
	"github.com/tgo-team/tgo-talkapi/cache"
	"github.com/tgo-team/tgo-talkapi/config"
	"github.com/tgo-team/tgo-talkapi/ctrl"
	"github.com/tgo-team/tgo-talkapi/handler/contacts"
	"github.com/tgo-team/tgo-talkapi/handler/db"
	"github.com/tgo-team/tgo-talkapi/handler/user"
	_ "github.com/tgo-team/tgo-talkapi/tgo/log"
	_ "github.com/tgo-team/tgo-talkapi/tgo/protocol/mqtt"
	_ "github.com/tgo-team/tgo-talkapi/tgo/server/tcp"
	_ "github.com/tgo-team/tgo-talkapi/tgo/storage/memory"
	"github.com/tgo-team/tgo-talkapi/utils"
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

	cfg := config.New()
	// 初始化DB
	initDB(cfg)
	redisClient := initRedisClient(cfg)

	ce := cache.NewRedisCache(redisClient)

	opts := tgo.NewOptions()
	opts.TCPAddress = "0.0.0.0:6667"
	t := tgo.New(opts)
	p.t = t
	// 开启控制器
	controller := ctrl.New(t, cfg,ce)
	registerHandlers(controller, cfg)
	controller.Start()
	// 开启TGO
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

func registerHandlers(c *ctrl.Controller, cfg *config.Config) {
	// 用户
	user.NewHandler(cfg).RegisterHandler(c)
	// 联系人
	contacts.NewHandler().RegisterHandler(c)
}

func initDB(cfg *config.Config) {
	db.Init(cfg.Mysql)

	migrations := &migrate.FileMigrationSource{
		Dir: "config/sql",
	}
	_, err := migrate.Exec(db.NewSession().DB, "mysql", migrations, migrate.Up)
	utils.CheckErr(err)
}

func initRedisClient(cfg *config.Config) *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
