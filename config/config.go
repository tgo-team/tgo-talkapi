package config

import (
	"github.com/BurntSushi/toml"
	"time"
)

type Config struct {
	TalkHttpUrl string `toml:"talk_http_url"` // talk服务器的http接口
	NodeId int64 `toml:"node_id"` // 节点唯一ID
	TokenExpire duration `toml:"token_expire"` // token失效时间
	Mysql MysqlConfig
	Redis RedisConfig
	CachePrefix CachePrefixConfig `toml:"cache_prefix"`

}

// mysql相关配置
type MysqlConfig struct {
	Addr string
	Db string
	User string
	Password string
}

// redis相关配置
type RedisConfig struct {
	Addr string
}

// 缓存相关配置
type CachePrefixConfig struct {
	TokenPrefix string `toml:"token_prefix"`
}

func New() *Config {
	var config Config
    _,err := toml.DecodeFile("config/config.toml",&config)
    if err!=nil {
    	panic(err)
	}
	return &config
}


type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}