package services

import (
	"context"
	"net/http"
	"strconv"
)

func ProdEncodeFunc(ctx context.Context,httpRequest *http.Request,requestParams interface{}) error {
	prodr := requestParams.(ProdRequest)
	httpRequest.URL.Path += "/product/" + strconv.Itoa(prodr.ProdId)  // 访问哪个path 参数是什么这里决定了
	return nil
}



