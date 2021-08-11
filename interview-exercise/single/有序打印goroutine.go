package main

import (
	"sync"
)

var ch chan struct{}

func job1() {
	println("job1")
	ch <- struct{}{}

}
func job2() {
	<-ch
	println("job2")
}
func job3() {
	<-ch
	println("job3")
}
func do(fs ...func()) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	for _, fn := range fs {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(fn)
	}
	return wg
}

func main() {
	ch = make(chan struct{})
	wg := do(job1, job2)
	wg.Wait()
}
