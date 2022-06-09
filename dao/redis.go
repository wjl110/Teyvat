package dao

import (
	"context"
	"douying/setting"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var CTX context.Context
func InitRedis(config *setting.RedisConfig) {

	CTX = context.Background()
	ip := fmt.Sprintf("%s:%d", config.Host, config.Port)
	RDB = redis.NewClient(&redis.Options{
		Addr:	 ip ,
		Password: "", // no password set
		DB:		  config.DB,  // use default DB
	})

}

func RedisGet(key string,s interface{})  {
	val, _ := RDB.Get(CTX, key).Result()
	_ = json.Unmarshal([]byte(val), s)
}

func RedisSet(key string, s interface{}) error {
	val, err := json.Marshal(s)
	RDB.Set(CTX, key, val,0)
	return err
}

