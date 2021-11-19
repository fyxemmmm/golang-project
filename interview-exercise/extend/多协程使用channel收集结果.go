package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	num := 5
	resCh := make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < num; i ++ {
		wg.Add(1)
		go func(param int) {
			defer wg.Done()
			resCh <- job(param)			
		}(i)
	}
	
	go func() {
		defer close(resCh)
		wg.Wait()
	}()
	
	for item := range resCh {
		fmt.Println("收到结果: ", item)
	}
}

func job(index int) int {
	time.Sleep(time.Millisecond * 500)
	return index
}