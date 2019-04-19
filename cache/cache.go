package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Cache interface {
	Set(key string,value string,expire time.Duration) error
	Get(key string) (string,error)
}

type RedisCache struct {
	redisClient *redis.Client
}

func NewRedisCache(redisClient *redis.Client)  *RedisCache {

	return &RedisCache{redisClient:redisClient}
}

func (r *RedisCache) Set(key string,value string,expire time.Duration) error  {
	println(fmt.Sprintf("key--%s  value-%s  expire-%d",key,value,expire))
	err := r.redisClient.Set(key,value,expire).Err()

	return err
}

func (r *RedisCache) Get(key string) (string,error)  {
	println(fmt.Sprintf("key--%s",key))
	result,err := r.redisClient.Get(key).Result()
	if err!=nil && err != redis.Nil {
		println("err",err)
		return "",err
	}
	println("result",result)
	return  result,nil
}

type MemoryCache struct {
   cacheMap map[string]string
}

func NewMemoryCache() *MemoryCache  {

	return &MemoryCache{
		cacheMap: map[string]string{},
	}
}

func (m *MemoryCache) Set(key string,value string,expire time.Duration) error  {
	m.cacheMap[key] = value
	return nil
}

func (m *MemoryCache) Get(key string) (string,error)  {

	return  m.cacheMap[key],nil
}