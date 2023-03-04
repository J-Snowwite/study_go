package main

import (
	"fmt"
	"sort"
)

func main() {
	//char => counter
	//A => 65  B=>  66 C=> 67
	stats := [][]int{{'A', 3}, {'B', 2}, {'C', 1}}
	//使用出现次数进行排序
	//https://golang.google.cn/pkg/sort/#Slice

	//使用出现次数进行排序
	//{B,2},{D,2} => 稳定
	//{D,2},{B,2} => 不稳定
	//
	sort.Slice(stats, func(i, j int) bool { return stats[i][1] > stats[j][1] }) //不稳定排序
	//此时是降序，，如果需要升序，改变大于号方向即可
	fmt.Println(stats)
	//sort.SliceStable()   稳定排序

	//使用search方法，必须是升序，否则会出现问题
	index := sort.Search(len(stats), func(i int) bool { return stats[i][1] == 1 })
	fmt.Println(index)

}
