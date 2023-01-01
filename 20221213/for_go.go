package main
import "fmt"

func main(){
	letters := "abcdefghi"
/*
	//很少使用
	for i :=0;i < len(letters);i++ {
		fmt.Println(i)
		fmt.Printf("%c\n",letters[i])
	}

//一般为了保证打印不报错，遍历字符串的时候都会使用range 方式他会自动匹配相关类型	
	letters = "我爱中华人民共和国"
	for k,v := range letters {
	//这里K是索引，v是值
		fmt.Printf("%T,%#v,%T,%q \n",k,k,v,v)
		//%U是指打印unicode格式
		//%q是打印字符串格式
	}
*/
	letters = "我爱中华人民共和国"
//	for k,v := range letters {
	//这样会出现问题，因为K没有被引用，如果
	for _,v := range letters {
		fmt.Printf("%q \n",v)
		//使用下划线 代替 K ，这样就不会报错了
		//_是go里面的一个内置变量，当不处理变量的时候可以使用——
	}

	var (
		index = 0
		total = 0
	)
	//for不加上判断调价是死循环，可以使用if break 跳出循环
	for {
		total += index
		index++
		if index > 100 {
			fmt.Println(total)
			break
		}
	}
}