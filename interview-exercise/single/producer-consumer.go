package main

import "fmt"

func Producer(ch chan int, producerFinishChan chan bool)  {
	for i := 0; i < 10;i ++ {
		ch <- i
	}
	// 生产完毕就下班 关掉通道
	//close(ch)
	producerFinishChan <- true
}

func Consumer(id int, ch chan int, done chan bool)  {
	for  {
		res, ok := <- ch
		if ok {
			fmt.Println(id, res)
		}else {
			fmt.Println("channel close")
			break
		}
	}
	done <- true
}

func main() {
	ch := make(chan int, 3)
	consumerNumber := 2
	done := make(chan bool, consumerNumber)
	for i := 0; i < consumerNumber; i ++ {
		go Consumer(i, ch, done)
	}
	
	producerNumber := 2
	producerDoneChan := make(chan bool, producerNumber)
	for i := 0; i < producerNumber; i ++ {
		go Producer(ch, producerDoneChan)
	}

	for i := 0; i <producerNumber; i ++ {
		<- producerDoneChan
	}
	close(ch)
	
	for i := 0; i < consumerNumber; i ++ {
		<- done
	}
	fmt.Println("finish")

}