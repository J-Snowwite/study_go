//一个冒泡排序法，如果不会的话直接使用 sort

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 7, 5, 4, 9, 2, 4, 1}
	sort.Ints(nums)
	//使用sort，对int类型的切片进行排序
	fmt.Println("正序排序", nums)

	//使用sort，对int类型的切片进行排序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("倒叙排序", nums)

	//冒泡排序
	height := []int{333, 222, 111, 444, 555, 999, 666, 777, 888}

	//数据交换--方式1
	a, b, tmp := 1, 2, 0

	if a < b {
		tmp = b
		b = a
		a = tmp
	}

	//数据交换--方式2--直接交换赋值--python也支持，其他语言不支持

	c, d := 55, 66
	fmt.Println(c, d)
	c, d = d, c
	fmt.Println(c, d)
	fmt.Printf("\n\n\n")
	//开始排序
	for k, v := range height {
		fmt.Printf("%T,%d,%T,%d \n", k, k, v, v)
	}

	//第二种
	for j := 0; j < len(height)-1; j++ {
		for i := 0; i < len(height)-1; i++ {
			if height[i] > height[i+1] {
				//fmt.Println(height[i])
				//类似 	c, d = d, c  同时赋值--交换值
				height[i], height[i+1] = height[i+1], height[i]
			}
			//fmt.Println(height)
		}
		//第一次循环是  0 和 1 、  1 和 2 、2 和 3
		//排序第一次后完成了相邻两个值的比较，
		//需要重复 len（height） -1 次才能一个字符和其他所有字符的比较
		//fmt.Println(height)

	}
	fmt.Println(height)

	//切片，获取欺骗中第二个最大元素
	//  1,2,3,4,5,5    并排第一   4
	// 				   不并排第一  5
	//插入排序算法
	//二分查找算法

}
