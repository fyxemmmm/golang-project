package main

import "fmt"

type Iterator struct {
	index int
	length int
}

func (this *Iterator) HasNext() bool {
	if this.length == 0{
		return false
	}
	return this.index < this.length
}

type IntIterator struct {
	*Iterator
	data []int  // 多类型 int改成string或者别的即可
}

func IterForInt(data []int) *IntIterator {
	return &IntIterator{data: data, Iterator: &Iterator{
		index:  0,
		length: len(data),
	}}
}

func (this *IntIterator) Next() int {
	defer func() {
		this.index ++
	}()
	return this.data[this.index]
}

func main()  {
	ints := []int{1,2,3,4,5}
	iter := IterForInt(ints)
	for iter.HasNext() {
		r := iter.Next()
		fmt.Println(r)
	}
}