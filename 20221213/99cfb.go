package main

import "fmt"

func main(){
	var (
		i = 1
		j = 1
	)
	for i=1;i <=9;i++ {
		for j=1;j <=i;j++ {
			fmt.Printf("%2d * %2d = %2d ",i,j,i*j)	
		//如果变量超过2个字符那么按照变量去打印
		}
		fmt.Printf("\n")
	}
}