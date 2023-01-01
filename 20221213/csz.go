package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var (
		number1 int
		number2 int
		number3 int
		number4 int
		restart string
	)

	//生成一个随机数
	rand.Seed(time.Now().Unix())
	//生成随机数种子---只需要运行一次,
	//为什么使用时间作为种子，是因为种子设置一样的话，生成的值也是一样的
	number2 = rand.Int()
	fmt.Println(number2)

	number3 = rand.Int()%100
	//生成0-100的数字,int类型值取余即可
	fmt.Println(number3)

	for {
		number4 = rand.Intn(100)
		//生成0-100的数字
		fmt.Println(number4)
		//判断结果
		for i :=1;i <= 6;i++ {
			if i == 6 {
				fmt.Println("太笨了，游戏结束")
				break
			}
			fmt.Println("请输入一个0-100的数字")
			fmt.Scan(&number1)
			if number1 > number4 {
				fmt.Println("输入值大了")
			}else if number1 < number4 {
				fmt.Println("输入值小了")
			}else{
				fmt.Println("答对了，真聪明")
			break
			}
		}
		fmt.Println("是否继续玩？，输入y或者n")
		fmt.Scan(&restart)
		if restart == "n"{
			break
		}
	}
}