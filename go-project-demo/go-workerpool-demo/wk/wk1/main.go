package main

import (
	"fmt"
	"suki.fyxemmmm.cn/my_test/wk/wk1/model"
	"suki.fyxemmmm.cn/my_test/wk/wk1/worker"
)

// complete: https://github.com/Joker666/goworkerpool

func main() {
	// Prepare the data
	var allData []model.SimpleData
	for i := 0; i < 1000; i++ {
		data := model.SimpleData{ ID: i }
		allData = append(allData, data)
	}
	fmt.Printf("Start processing all work \n")

	// Process
	//basic.Work(allData)

	// Process
	worker.PooledWorkError(allData)
}