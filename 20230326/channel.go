package main

import "fmt"

//channel  是一个环形队列---先进先出
//send(插入)
//recv(取出)
//从同一个位置沿同一个方向顺序执行
//sendX表示最后一次插入元素的位置
//recvx表示最后一次取出元素的位置
func main() {
	var ch2 chan int
	fmt.Printf("%t\n\n", ch2 == nil)
	//channel空值是 nil
	fmt.Printf("%d\n\n", len(ch2))
	//默认长度是 0
	ch2 = make(chan int, 10)
	fmt.Printf("%d\n\n", len(ch2))
	//初始化后长度依旧是10
	fmt.Printf("%d\n\n", cap(ch2))
	//初始化后容量是  10

	var ch1 chan int        //声明一个channel变量 ch1
	ch1 = make(chan int, 8) //用make 初始化 ch1，容量是8
	ch1 <- 1                //往队列里面写数据
	ch1 <- 2
	ch1 <- 3
	fmt.Println(ch1)
	v := <-ch1 //从队列里面取出数据
	//当管道存储满了之后，在插入就会报错，死锁
	//当管道存储空了之后，在取出就会报错，死锁
	fmt.Println(v)

	/*
		//下述方式无需关闭管道，即可遍历
		//下述方式需要将 channel的长度固定为数值，
		//因为len(ch1)一致在变，每取出一个，数值减少1
		lenght1 ：= len(ch1)
		for i :=0 i< lenght1; i++{
			ele := <-ch1
			fmt.Println(ele)
		}
	*/

	//遍历管道之前需要先关闭通道，禁止写入数据
	close(ch1)
	//遍历管道----只能遍历一次，遍历完成后数据消失
	for ele := range ch1 {
		fmt.Println(ele)
	}

}
