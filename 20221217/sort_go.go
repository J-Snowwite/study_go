package main

import (
	"fmt"
	"sort"
)

// 直接访问  golang.google.cn/pkg/sort
func main() {
	nums := []int{3, 7, 5, 4, 9, 2, 4, 1}
	sort.Ints(nums)
	//使用sort，对int类型的切片进行排序
	fmt.Println("正序排序", nums)
	//打印排序结果

	//倒叙
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	//使用sort，对int类型的切片进行倒叙排序
	fmt.Println("倒序排序", nums)
	//打印排序结果

	names := []string{"A", "b", "C", "e", "h"}
	sort.Strings(names)
	//使用sort，对string类型的切片进行排序
	fmt.Println(names)

	//查找---go中存在的二分法查找，局限，只能是顺序的数列,切必须是从小到大
	//[0,100)
	nums2 := []int{2, 3, 4, 5, 6, 8, 9}
	fmt.Println(sort.SearchInts(nums2, 8))
	fmt.Println(nums[sort.SearchInts(nums2, 8)] == 8)

}
