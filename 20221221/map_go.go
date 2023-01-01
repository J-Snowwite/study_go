package main

import "fmt"

func main() {

	//map
	//映射：由键值对组成的一种数据结构，通过key来对value进行操作
	//一种无序的数据结构,实现方式hashtable

	//每个同学的成绩
	// key => ID value => 成绩

	var scores map[string]float64

	fmt.Printf("%T   %#v\n\n", scores, scores)

	//初始化--字面量
	scores = map[string]float64{} //空的map
	fmt.Printf("%T %#v\n\n", scores, scores)
	//赋值
	scores = map[string]float64{"A": 88, "B": 99, "C": 77}
	fmt.Printf("%T %#v\n\n", scores, scores)

	//make
	scores = make(map[string]float64) //==map[string]float64{}
	//make不能初始化值，也不能设置容量
	fmt.Printf("%T %#v\n\n", scores, scores)

	//操作
	scores = map[string]float64{"A": 88, "B": 99, "C": 77}
	//添删改查
	fmt.Println(len(scores))
	//1、查,查找失败会返回对应值类型的 零 值
	fmt.Println(scores["B"])
	//判断键值对是否存在
	v1, v2 := scores["B"]
	fmt.Println(v1, v2)
	//2、改
	scores["C"] = 55
	fmt.Println(scores["C"])
	//3、添,如果D 这个键存在的话就是修改，如果不存在就是添加
	scores["D"] = 44
	v4, _ := scores["D"]
	fmt.Println(v4)
	//4、删,删除一个不存在的键值对，没有任何反应
	delete(scores, "A")
	_, v5 := scores["A"]
	fmt.Println(v5)

	//遍历,不能使用len，因为无法使用索引
	for k1, k2 := range scores {
		fmt.Println(k1, k2)
	}

	var nilMap map[string]string
	fmt.Println(len(nilMap))
	//下面会报错，nil映射不能直接添加元素，需要先赋值初始化在添加元素
	//尽可能使用    空切片  或者  空映射  不要使用nil切片nil映射
	nilMap["kk"] = "XX"
	fmt.Println(nilMap["kk"])

}
