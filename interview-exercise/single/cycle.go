package main
import (
	"math/rand"
	"sync"
	"time"
)
func myjob(){
	time.Sleep(time.Second*3)
	println(rand.Intn(100))
}
func setpool(ch chan struct{}){
	for i:=0;i<5;i++{
		ch<- struct{}{}
	}
}
func main() {
	pool:=make(chan struct{},5)
	setpool(pool)
	wg:=sync.WaitGroup{}
	wg.Add(5)
	go func() {
		for {
			wg.Wait()
			println("发放了5个任务")
			setpool(pool)
			wg.Add(5)
		}
	}()
	for {
		<-pool
		go func() {
			defer wg.Done()
			myjob()
		}()
	}
}