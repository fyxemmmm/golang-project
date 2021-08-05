package main

import (
	"fmt"
	"time"
)

func main() {
	mysqlData := make(chan interface{})
	fanOut(mysqlData, redisJob())

	for i := 0; i < 10;i ++{
		mysqlData <- i
	}

	time.Sleep(time.Hour)
}

func redisJob() chan interface{} {
	redisCh := make(chan interface{})
	go func() {
		for v := range redisCh {
			time.Sleep(time.Millisecond * 500)
			fmt.Println("数据", v, "写入了redis")
		}
	}()
	return redisCh
}

func fanOut(data <-chan interface{}, inCh ...chan interface{})  {
	go func() {
		defer func() {
			for i := 0; i < len(inCh); i ++ {
				close(inCh[i])
			}
		}()
		// 遍历数据源的每一个值
		for v := range data {
			for i :=0 ; i <len(inCh); i ++{
				inCh[i] <- v
			}
		}
	}()
}
