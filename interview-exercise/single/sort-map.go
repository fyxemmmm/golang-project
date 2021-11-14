package main

import (
	"fmt"
	"sort"
)

type DataSorter []Item  //构建一个排序器
type Item struct {  //为了把map转为切片，故意弄得一个struct 为了保存map的key和value
	Key string
	Val *Data
}
type Data struct {  //  map里面的 业务类 。模仿你的 代码
	ProcessName string
}
//这个函数把map转为 切片 --必须
func NewDataSorter(m map[string]*Data) DataSorter {
	ms := make(DataSorter, 0, len(m))
	for k, v := range m {
		ms = append(ms, Item{k, v})
	}
	return ms
}
func (ms DataSorter) Len() int {  //官方规定方法 必须实现,没有为什么
	return len(ms)
}
func (ms DataSorter) Less(i, j int) bool { //官方规定方法 必须实现 没有为什么
	return ms[i].Val.ProcessName > ms[j].Val.ProcessName  // ProcessName排序
	//仔细看这里。 这怎么是根据int排序呢。 我还是按照ProcessName排序啊
	//如果要按照key排序 那就是
	//return ms[i].Key<ms[j].Key// 看这
}
func (ms DataSorter) Swap(i, j int) {//官方规定方法 必须实现  没有为什么
	ms[i], ms[j] = ms[j], ms[i]
}
func main() {
	//构建测试数据
	test:=map[string]*Data{}
	test["java"]=&Data{ProcessName: "java进程"}
	test["php"]=&Data{ProcessName: "php进程"}
	test["go"]=&Data{ProcessName: "golang进程"}

	//开始排序
	sorter:=NewDataSorter(test)   //必须执行转换
	sort.Sort(sorter)  //执行排序

	for _,d:=range sorter{ //排序后就是一个 切片了 ,不是原来的map
		fmt.Println(d.Key)
	}

	// 最终你只要 遍历出来的值 是 按序的就可以了
}