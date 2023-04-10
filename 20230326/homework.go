package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// 生成一个切片
func GenSlice(n int) []int {
	sou1 := rand.NewSource(22)
	rand1 := rand.New(sou1)
	arr := make([]int, 10)
	for i := 0; i < n; i++ {
		arr = append(arr, rand1.Intn(128))
	}
	return arr
}

// 统计切片不同的值的数量
func CountUniq(arr []int) int {
	m := make(map[int]bool, len(arr))
	//创建一个map，map会对相同的key值进行排除
	//直接将遍历出来的值作为map的key就能实现去重
	for _, v2 := range arr {
		m[v2] = true
	}
	return len(m)
}

// 字符拼接
func Concat(arr []int) string {
	var sb strings.Builder
	for _, ele := range arr {
		//将数据遍历出来，转换成字符串类型，写入sb中
		sb.WriteString(strconv.Itoa(ele))
		sb.WriteString("")
	}
	//返回信息去除空格
	//return strings.Trim(sb.String(), " ")
	return sb.String()[0 : len(sb.String())-1]
}

func main() {
	arr := GenSlice(100)
	cnt := CountUniq(arr)
	fmt.Println(cnt)

	s := Concat(arr)
	fmt.Println(s)
}
