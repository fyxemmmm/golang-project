package main

import (
	"fmt"
	"sync"
)

type Cmd func(list []int) chan int
type PipeCmd func(in chan int) chan int // 支持管道的函数
func Evens(list []int) chan int{
	c := make(chan int)
	go func() {
		defer close(c)
		for _,num:=range list {
			if num%2==0{ // 业务流程  所谓的异步
				c <- num
			}
		}
	}()
	return c
}

// 可能两个都执行完成了
func M2(in chan int) chan int{ // 支持管道的 参数一律是channel
	out := make(chan int)
	go func() {
		defer close(out)
		for num:=range in {
			out <- num * 2
		}
	}()
	return out
}

// 可能两个都执行完成了
func M10(in chan int) chan int{ // 支持管道的 参数一律是channel
	out := make(chan int)
	go func() {
		defer close(out)
		for num:=range in {
			out <- num * 10
		}
	}()
	return out
}


// 管道函数
func Pipe(args []int, c1 Cmd, cs ...PipeCmd) chan int {
	ret := c1(args)  // 同步的, 拿到结果才能交给下一个
	if len(cs) == 0 {
		return ret
	}
	retlist := make([]chan int, 0)

	for index, c := range cs {
		if index == 0 {
			retlist = append(retlist, c(ret))
		}else {
			getChan := retlist[len(retlist)-1] // 取最后一个作为结果
			retlist = append(retlist, c(getChan))
		}
	}
	return retlist[len(retlist)-1]
}



//多路复用
func Pipe2(args []int ,c1 Cmd,cs ...PipeCmd) chan int{
	ret:=c1(args) //找偶数
	out:=make(chan int)
	wg:=sync.WaitGroup{}
	for _,c:=range cs{
		getChan:=c(ret)
		wg.Add(1)
		go func(input chan int) {
			defer wg.Done()
			for v:=range input{
				out<-v
			}
		}(getChan)
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}

func main()  {
	nums:=[]int{2,3,6,12,22,16,4,9,23,64,62}
	ret := Pipe(nums, Evens, M10, M10)
	for r := range ret {
		fmt.Printf("%d ", r)
	}
}

