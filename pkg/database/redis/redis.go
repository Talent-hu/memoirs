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
	PoolSize     int           `yaml:"poolSize"`
	Addr         []string      `yaml:"addr"`
	Pwd          string        `yaml:"pwd"`
	DialTimeout  time.Duration `yaml:"dialTimeout"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

func NewRedis(c Config) (client *Client, err error) {
	var redisCli redis.Cmdable
	if len(c.Addr) > 1 {
		redisCli = redis.NewClusterClient(
			&redis.ClusterOptions{
				Addrs:        c.Addr,
				PoolSize:     c.PoolSize,
				DialTimeout:  c.DialTimeout,
				ReadTimeout:  c.ReadTimeout,
				WriteTimeout: c.WriteTimeout,
				Password:     c.Pwd,
			})
	} else {
		redisCli = redis.NewClient(&redis.Options{
			Addr:         c.Addr[0],
			DialTimeout:  c.DialTimeout,
			ReadTimeout:  c.ReadTimeout,
			WriteTimeout: c.WriteTimeout,
			Password:     c.Pwd,
			PoolSize:     c.PoolSize,
			DB:           0,
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
