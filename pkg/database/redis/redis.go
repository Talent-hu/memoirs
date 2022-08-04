package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Client struct {
	redis.Cmdable
}

type Config struct {
	Addr         []string      `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password     string        `mapstructure:"password" json:"password" yaml:"password"`
	DB           int           `mapstructure:"db" json:"db" yaml:"db"`
	MaxIdle      int           `mapstructure:"max-idle" json:"max-idle" yaml:"max-idle"`
	MaxActive    int           `mapstructure:"max-active" json:"max-active" yaml:"max-active"`
	MaxTimeout   int           `mapstructure:"max-timeout" json:"max-timeout" yaml:"max-timeout"`
	PoolSize     int           `mapstructure:"pool-size" json:"pool-size" yaml:"pool-size"`
	DialTimeout  time.Duration `mapstructure:"dial-timeout" json:"dial-timeout" yaml:"dial-timeout"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout" json:"read-timeout" yaml:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout" json:"write-timeout" yaml:"write-timeout"`
}

func NewRedis(c *Config) (client *Client, err error) {
	var redisCli redis.Cmdable
	if len(c.Addr) > 1 {
		redisCli = redis.NewClusterClient(
			&redis.ClusterOptions{
				Addrs:        c.Addr,
				PoolSize:     c.PoolSize,
				DialTimeout:  c.DialTimeout,
				ReadTimeout:  c.ReadTimeout,
				WriteTimeout: c.WriteTimeout,
				Password:     c.Password,
			})
	} else {
		redisCli = redis.NewClient(&redis.Options{
			Addr:         c.Addr[0],
			DialTimeout:  c.DialTimeout,
			ReadTimeout:  c.ReadTimeout,
			WriteTimeout: c.WriteTimeout,
			Password:     c.Password,
			PoolSize:     c.PoolSize,
			DB:           c.DB,
		})
	}
	err = redisCli.Ping(context.Background()).Err()
	if nil != err {
		panic(err)
	}

	client = new(Client)
	client.Cmdable = redisCli
	return client, nil
}

func (c *Client) Process(cmd redis.Cmder) error {
	switch redisCli := c.Cmdable.(type) {
	case *redis.ClusterClient:
		return redisCli.Process(context.Background(), cmd)
	case *redis.Client:
		return redisCli.Process(context.Background(), cmd)
	default:
		return nil
	}
}

func (c *Client) Close() error {
	switch redisCli := c.Cmdable.(type) {
	case *redis.ClusterClient:
		return redisCli.Close()
	case *redis.Client:
		return redisCli.Close()
	default:
		return nil
	}
}
