package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Resolve func(rsp interface{})
type Reject func(err error)
type PromiseFunc func(resolve Resolve, reject Reject)
type PromiseOpt func(promise *Promise)
type PromiseOpts []PromiseOpt

func (this PromiseOpts) apply(p *Promise) {
	for _, opt := range this{
		opt(p)
	}
}

func WithTimeout(t time.Duration) PromiseOpt {
	return func(promise *Promise) {
		promise.timeout = t
	}
}

type Promise struct {
	f PromiseFunc
	resolve Resolve
	reject Reject
	wg sync.WaitGroup
	timeout time.Duration
}

func NewPromise(f PromiseFunc) *Promise {
	return &Promise{f: f, timeout: time.Second * 10}
}

func (this *Promise) Then(resolve Resolve) *Promise {
	this.resolve = resolve
	return this
}

func (this *Promise) Catch(reject Reject) *Promise {
	this.reject = reject
	return this
}

func (this *Promise) Done(opts ...PromiseOpt) {
	defer func() {
		if e:=recover();e!= nil {
			this.reject(fmt.Errorf(e.(string)))
		}
	}()
	PromiseOpts(opts).apply(this)
	this.wg.Add(1)
	timeoutCtx, _ := context.WithTimeout(context.Background(), this.timeout)
	done := make(chan struct{})
	go func() {
		defer this.wg.Done()
		this.f(this.resolve, this.reject)
	}()

	go func() {
		defer func() {
			done <- struct{}{}
		}()
		this.wg.Wait()
	}()

	select {
	case <- timeoutCtx.Done():
		panic("超时")
	case <- done:
		log.Println("正常执行")
	}
}

func main() {
	NewPromise(func(resolve Resolve, reject Reject) {
		time.Sleep(time.Second*2)
		if time.Now().Unix() % 2 == 0 {
			resolve("ok")
		}else {
			reject(fmt.Errorf("my err"))
		}
	}).Then(func(rsp interface{}) {
		fmt.Println(rsp)
	}).Catch(func(err error) {
		log.Println(err)
	}).Done(WithTimeout(time.Second*1))
}

