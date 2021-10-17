package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

var (
	c *redis.ClusterClient
)

func init()  {
	ctx := context.Background()
	c = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              []string{"xx.fyxemmmm.cn:32056"},
	})

	if err := c.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	//defer c.Close()
}

func TestGet(t *testing.T)  {
	s, err := c.Get(context.Background(), "k4").Result()
	if err != nil {
		t.Log(err.Error())
		return
	}
	fmt.Println(s)
}

func TestSet (t *testing.T)  {
	err := c.Set(context.Background(), "k4", "v4", 0).Err()
	if err != nil {
		panic(err)
	}
}

func TestDel (t *testing.T)  {
	err := c.Del(context.Background(), "k4").Err()
	if err != nil {
		panic(err)
	}
}

func TestZAdd (t *testing.T)  {
	vals := []*redis.Z{
		{
			Score: 10,
			Member: "one",
		},
		{
			Score: 3,
			Member: "two",
		},
	}
	addCount, err := c.ZAdd(context.Background(), "papersocre", vals...).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(addCount)
}

func TestZRange(t *testing.T)  {
	//c.ZRevRange() 从大到小
	res, err := c.ZRange(context.Background(), "papersocre", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)	
}

func TestHMSet(t *testing.T)  {
	b, err := c.HMSet(context.Background(), "cart", "key1", "value2", "key2", "value3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}