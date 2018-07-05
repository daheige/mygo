package util

import (
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

//redis opts类型
type RedisHandler struct {
	Client *redis.Client
	Opts   *redis.Options
	once   sync.Once
}

func (this *RedisHandler) GetClient() {
	this.once.Do(func() {
		client, err := InitClient(this.Opts)
		if err != nil {
			log.Fatalf("redis connect error: %s", err)
			return
		}
		this.Client = client
	})
}

func GetOption() *redis.Options {
	return &redis.Options{}
}

//通过工厂模式创建一个redis obj
func NewRedis(opts *redis.Options) *RedisHandler {
	var obj = &RedisHandler{
		Opts: opts,
	}

	obj.GetClient()
	return obj
}

//获得一个redis client
//默认db是0
func InitClient(opts *redis.Options) (client *redis.Client, err error) {
	//redis options设置
	opts.DialTimeout = time.Second * 2
	opts.ReadTimeout = time.Second * 2
	opts.WriteTimeout = time.Second * 1
	opts.PoolSize = 16 //最大pool个数
	opts.IdleTimeout = 10 * time.Minute

	client = redis.NewClient(opts)
	pong, err := client.Ping().Result() //PONG <nil>
	if err != nil {
		log.Fatalf("redis connect error: %s", err)
		return
	}

	log.Println(pong)

	return
}
