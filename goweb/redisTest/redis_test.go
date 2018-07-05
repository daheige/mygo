package redisTest

import (
	"flag"
	"goweb/util"
	"testing"
)

var redisAddress string
var redisPwd string
var redisDefaultDb int = 0

func init() {
	flag.StringVar(&redisAddress, "redis_address", "127.0.0.1:6379", "redis address")
	flag.StringVar(&redisPwd, "redis_pwd", "", "redis pwd")
}

func Test_redis(t *testing.T) {
	var opts = util.GetOption()
	opts.Addr = redisAddress
	opts.Password = redisPwd
	opts.DB = redisDefaultDb
	t.Log(opts)
	var obj = util.NewRedis(opts)

	err := obj.Client.Set("name", "daheige", 0).Err()
	if err != nil {
		t.Fatal(err)
	}

	var r = &util.RedisHandler{
		Opts: opts,
	}

	r.GetClient()
	t.Log(r)

	err = r.Client.Set("user", "heige", 0).Err()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("test success")
}
