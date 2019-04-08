package config

import "github.com/BurntSushi/toml"

type Config struct {
	TalkHttpUrl string `toml:"talk_http_url"`
	NodeId int64 `toml:"node_id"`
	Mysql MysqlConfig

}

type MysqlConfig struct {
	Addr string
	Db string
	User string
	Password string
}

func New() *Config {
	var config Config
    _,err := toml.DecodeFile("config/config.toml",&config)
    if err!=nil {
    	panic(err)
	}
	return &config
}
