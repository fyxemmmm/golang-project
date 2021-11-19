package main

import (
	"fmt"
	"sync"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	return &Mutex{ch: make(chan struct{}, 1)}
}

func (this *Mutex) Lock() {
	this.ch <- struct{}{}
}

func (this *Mutex) Unlock() {
	select {
	case <- this.ch:
	default:
		panic("unlock error")
	}
}


type Job struct {
	Count int
	lock *Mutex
}

func (this *Job) Incr() {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.Count ++
}

func (this *Job) Decr() {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.Count --
}

// 可以实现超时控制等官方mutex没有实现的功能
func main()  {
	job := &Job{lock: NewMutex()}
	wg := sync.WaitGroup{}
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i ++ {
			job.Incr()
		}
	}()
	
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i ++{
			job.Decr()
		}
	}()
	
	wg.Wait()
	fmt.Println(job.Count)
	
}