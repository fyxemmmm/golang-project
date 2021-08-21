package util

import (
	"context"
	"io/ioutil"
	"net/http"

)

type Client struct {
	Services []*ServiceInfo
}

type ServiceInfo struct {
	ServiceID string
	ServiceName string
	ServiceAddr string
}


func (this *Client) GetService (sname string, method string, encodeFunc EncodeRequestFunc) Endpoint {
	for _, service := range this.Services{
		if service.ServiceName == sname{
			return func(ctx context.Context, requestParam interface{}) (responseResult interface{}, err error) {
				httpClient := &http.Client{}
				// url仅仅是ip
				httpRequest,err := http.NewRequest(method, "http://" + service.ServiceAddr, nil)
				//
				err = encodeFunc(ctx, httpRequest, requestParam) // 这一步是灵活的
				if err != nil{
					return nil, err
				}
				res, err := httpClient.Do(httpRequest)
				defer res.Body.Close()
				if err != nil{
					return nil, err
				}
				body,err := ioutil.ReadAll(res.Body)
				if err != nil{
					return nil, err
				}
				return string(body), nil
			}
		}
	}
	return nil
}
