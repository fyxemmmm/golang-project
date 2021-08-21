package main

import (
	"context"
	"fmt"
	"log"
	"suki.fyxemmmm.cn/myetcd/services"
	"suki.fyxemmmm.cn/myetcd/util"
)

func main()  {
	client := &util.Client{}
	endpoint := client.GetService("productservice", "GET", services.ProdEncodeFunc)
	res,err := endpoint(context.Background(), services.ProdRequest{ProdId: 106})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(res)
}

