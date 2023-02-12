package main

import (
	"fmt"
	"strings"
)

func main() {
	//查看官网
	//https://golang.google.cn/pkg/strings/   查看字符串相关
	//https://golang.google.cn/pkg/fmt/   查看fmt模块相关的

	fmt.Println(strings.Compare("a", "b"))
	//比较两个字符  如果相等 输出0 如果小于输出 -1 如果大于输出 1
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("c", "b"))

	//判断第一个参数是否包含第二个参数。是返回true ，不是返回false
	fmt.Println(strings.Contains("我是 AAA", "A"))
	fmt.Println(strings.Contains("我是 AAA", "B"))

	//任意包含  即 第一个和第二个是否存在重叠字符，存在返回true，不存在返回false
	fmt.Println(strings.ContainsAny("我是 AAA", "AB"))
	fmt.Println(strings.ContainsAny("我是 AAA", "B"))

	//计算字符出现的次数，统计单字符出现的次数
	fmt.Println(strings.Count("我是AA,AA是我", "AA"))
	fmt.Println(strings.Count("我是AA,AA是我", "A"))
	fmt.Println(strings.Count("我是AA,AA是我", "B"))

	//比较--字符串之间的比较--不区分大小写，正确输出ture
	fmt.Println(strings.EqualFold("abc", "ABC"))
	fmt.Println(strings.EqualFold("abc", "abc"))
	fmt.Println(strings.EqualFold("abc", "xyz"))

	//空白符--遇到一个空白符就变成一个字符串放置到切片里面----
	//空格  tab  回车  换行  换页  。。。。都属于空白符
	fmt.Printf("%#v\n", strings.Fields("aafds d\tc\nd\re\ff"))

	//开头(HasPrefix)——结尾(HasSuffix)
	fmt.Println(strings.HasPrefix("abc", "a"))
	fmt.Println(strings.HasPrefix("abc", "b"))

	fmt.Println(strings.HasSuffix("abc", "c"))
	fmt.Println(strings.HasSuffix("abc", "b"))

	//字符串出现的位置
	//只会寻找第一个,成功返回字符索引值，否则返回-1
	fmt.Println(strings.Index("abcdabcd", "c"))
	fmt.Println(strings.Index("abcdabcd", "cc"))

	//字符串出现的位置-最后一次出现
	//只会寻找倒数第一个,成功返回字符索引值，否则返回-1
	fmt.Println(strings.LastIndex("abcdabcd", "c"))

	//链接---返回字符串  分割---返回的是一个切片
	fmt.Printf("%#v\n", strings.Join([]string{"ab", "cd", "ef"}, "-"))
	fmt.Printf("%#v\n", strings.Split("aa-bb-cc", "-"))

	//重复--l连续重复打印某个字符
	fmt.Println(strings.Repeat("#", 6))

	//替换---替换字符,参数-1表示替换所有，1表示只替换一个
	fmt.Println(strings.Replace("abcdabcdabcdabcd", "ab", "mm", -1))
	fmt.Println(strings.Replace("abcdabcdabcdabcd", "ab", "mm", 1))
	fmt.Println(strings.Replace("abcdabcdabcdabcd", "ab", "mm", 2))
	fmt.Println(strings.ReplaceAll("abcdabcdabcdabcd", "ab", "mm"))

	//首字母大写
	fmt.Println(strings.Title("abcd"))

	//全部转成大写、小写
	fmt.Println(strings.ToLower("abcABC"))
	fmt.Println(strings.ToUpper("abcABC"))

	//trim  剔除
	fmt.Println(strings.Trim("abcde", "ab"))

	//剔除空白符
	//空格  tab  回车  换行  换页  。。。。都属于空白符
	fmt.Println(strings.TrimSpace("  \n \f\t abc"))

	//从左边剔除，从右边开始剔除
	fmt.Println(strings.TrimLeft("abcdefedcba", "a"))
	fmt.Println(strings.TrimRight("abcdefedcba", "a"))

	//从左边剔除和右边剔除，可以和上述的互等
	fmt.Println(strings.TrimPrefix("abcdefedcba", "a"))
	fmt.Println(strings.TrimSuffix("abcdefedcba", "a"))

}
