package main
import "fmt"

func main(){
	var y string
	fmt.Println("有卖西瓜的那？")
	fmt.Scan(&y)

	//值比较
	switch y {
	case "yes","Y","y":
			fmt.Println("一个包子")
	case "no","N","n":
			fmt.Println("十个包子")
	default:
			fmt.Println("请输入yes或者是no")
	}

	//表达式
	var x int
	fmt.Println("输入一个数字")
	fmt.Scan(&x)
	switch {
	case x >= 90:
		fmt.Println("优秀")
	case x >= 60:
		fmt.Println("还行")
	default:
		fmt.Println("差")
	}
}