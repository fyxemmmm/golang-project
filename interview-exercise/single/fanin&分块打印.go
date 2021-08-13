package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		// 获取用户信息和几分和别的xxx
		c.Writer.Header().Add("Transfer-Encoding", "chunked")
		res := fanIn(getUserInfo(), getUserScore())
		c.Writer.WriteHeader(200)
		for v := range res {
			c.Writer.Write([]byte(v.(string)))
			c.Writer.(http.Flusher).Flush()
		}
	})
	r.Run(":8080")
}


func getUserInfo() <-chan interface{} {
	infoChan := make(chan interface{})
	go func() {
		defer close(infoChan)
		time.Sleep(time.Second * 2) // 模拟业务很耗时
		infoChan <- "<p>get user info</p>"
	}()

	return infoChan
}


func getUserScore() <-chan interface{} {
	infoChan := make(chan interface{})
	go func() {
		defer close(infoChan)
		time.Sleep(time.Second * 3) // 模拟业务很耗时
		infoChan <- "<p>get user score: 10</p>"
	}()
	return infoChan
}


func fanIn(chs ...<- chan interface{}) <- chan interface{}{
	res := make(chan interface{})
	wg := sync.WaitGroup{}
	for _, ch := range chs {
		wg.Add(1)
		go func(c <- chan interface{}) {
			defer wg.Done()
			for v := range c {
				res <- v
			}
		}(ch)
	}
	go func() {
		defer close(res)
		wg.Wait()
	}()
	return res
}