package main

import (
	"fmt"
	"sync"
)

func main()  {
	users := []string{"feixiang", "zhangsan", "lisi"}
	ages := []string{"19", "21", "25"}
	c1 := make(chan bool, 1)
	c2 := make(chan bool)
	res := make([]string, 0)
	wg := sync.WaitGroup{}
	wg.Add(2)
	done := make(chan int)
	go func() {
		defer wg.Done()
		for _, v := range users {
			<- c1
			res = append(res, v)
			c2 <- true
		}
	}()

	go func() {
		defer wg.Done()
		for _, v := range ages {
			<- c2
			res = append(res, v)
			c1 <- true
		}
	}()


	go func() {
		defer func() {
			done<-1
		}()
		wg.Wait()
		for _,item:=range res{
			fmt.Println(item)
		}
	}()

	c1 <- true
	<- done
}

