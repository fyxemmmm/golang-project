package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func job1(id int) (string, error) {
	if rand.Intn(1000) % 2 == 0 {
		return "", fmt.Errorf("发生错误, ID=%d", id)
	}
	return fmt.Sprintf("成功结果, ID=%d", id), nil
}
func main()  {
	resChan := make(chan interface{}) // 或者这里写100, 消除一个go func
	pool := make(chan struct{}, 10)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i ++ {
			pool <- struct{}{}
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				defer func() {
					<- pool
				}()
				res, err := job1(index)
				if err != nil {
					resChan <- err
				}else {
					resChan <- res
				}
			}(i)
		}
	}()
	
	
	go func() {
		defer close(resChan)
		wg.Wait()
	}()
	
	errCnt := 0
	for item := range resChan {
		if err, ok := item.(error); ok {
			fmt.Println("ErrorInfo: ", err.Error())
			errCnt ++
		}else {
			fmt.Println("SuccessInfo: ", item)
		}
		if errCnt == 2 {
			break
		}
	}
	
}