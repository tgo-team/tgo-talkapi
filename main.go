package main

import (
	"fmt"
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
	"github.com/tgo-team/tgo-talkapi/plugin"
	_ "github.com/tgo-team/tgo-talkapi/tgo/log"
	_ "github.com/tgo-team/tgo-talkapi/tgo/protocol/mqtt"
	_ "github.com/tgo-team/tgo-talkapi/tgo/server/tcp"
	_ "github.com/tgo-team/tgo-talkapi/tgo/storage/memory"
	"github.com/tgo-team/tgo-talkapi/utils"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	pn "plugin"
	"strings"
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
	//initDB(cfg)
	redisClient := initRedisClient(cfg)

	// 初始化缓存
	cache := cache.NewRedisCache(redisClient)

	/// tgo初始化
	opts := tgo.NewOptions()
	opts.TCPAddress = "0.0.0.0:6667"
	t := tgo.New(opts)
	p.t = t
	// 开启控制器
	controller := ctrl.New(t, cfg, cache)
	registerHandlers(controller, cfg)
	controller.Start()
	// 开启TGO
	err := t.Start()
	if err != nil {
		panic(err)
	}
	// 初始化插件
	initPlugins(plugin.NewContext(controller,t))
	return nil
}

/// 初始化插件
func initPlugins(context *plugin.Context) {
	var pluginRoot = "./plugin_dir"
	files, err := ioutil.ReadDir(pluginRoot)
	utils.CheckErr(err)
	if files != nil {
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".so") {
				initPlugin(path.Join(pluginRoot, "/", file.Name()), context)
			}
		}
	}
}

func initPlugin(filePath string, context *plugin.Context) {
	fmt.Println(filePath)
	p, err := pn.Open(filePath)
	utils.CheckErr(err)
	setup, err := p.Lookup("Setup")
	utils.CheckErr(err)
	setupFunc := setup.(func(ctx *plugin.Context))
	if setupFunc != nil {
		setupFunc(context)
	}
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

func initRedisClient(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
