package main

import "fmt"

func main() {
	fmt.Println("打印数组")
	//数组  长度固定的，相同数据类型的一组值  的   序列
	//类型  【length】type
	//填充值都是对应变量的 0 值
	//长度一旦固定就无法修改

	var names [5]string
	fmt.Printf("%T\n", names)
	fmt.Printf("%#v\n\n", names)

	//确定一组真假
	var yon [5]bool
	fmt.Printf("%T\n", yon)
	fmt.Printf("%#v\n\n", yon)

	//赋值
	names = [5]string{"我", "爱", "中", "国", "呀"}
	fmt.Printf("%#v\n\n", names)

	//推断赋值,根据填充的数据，确定数组的长度
	names = [...]string{"推", "断", "赋", "值", "呀"}
	fmt.Printf("%#v\n\n", names)

	//初始化赋值
	var echoDate [5]string = [5]string{"我", "爱", "中", "国", "吗"}
	fmt.Printf("%#v\n\n", echoDate)

	//常用的方式
	testName := [...]string{"赋", "值", "呀"}
	fmt.Printf("%#v\n\n", testName)

	//部分赋值，其他不变，通过索引赋值,通过最后一个值确认数组长度
	testDate := [...]string{5: "好"}
	fmt.Printf("%#v\n\n", testDate)

	//指定数组长度
	testLen := [6]string{5: "好"}
	fmt.Printf("%#v\n\n", testLen)

	//操作
	//关系运算 ==   ！=   首先进行运算的数组 长度和类型需要保持一致
	fmt.Println(testDate != [6]string{})
	fmt.Println(testDate == testLen)
	fmt.Printf("\n\n")

	//元素
	//访问  &  修改  &  索引（0,1,2,3,4）
	fmt.Println(testDate[5])
	testDate[5] = "坏"
	fmt.Println(testDate[5])
	fmt.Printf("\n\n")

	//遍历
	for i := 0; i < len(echoDate); i++ {
		fmt.Println(echoDate[i])
	}
	fmt.Printf("\n\n")
	//上述方式不推荐，主要是不灵活
	for _, date1 := range echoDate {
		fmt.Println(date1)
	}
	fmt.Printf("\n\n")
	//上述方式非常推荐

	//多维数组
	dateList2 := [2][2]int{}
	//[0]dateList2==> [2]int ==> int{0,0}
	//[1]dateList2==> [2]int ==> int{0,0}
	fmt.Printf("%#v\n\n", dateList2)
	//赋值--索引
	dateList2[1][1] = 5
	fmt.Printf("%#v\n\n", dateList2[1][1])
	//赋值-初始化

	/*
		dateList3 := [2][2]int{0: [2]int{22, 22}, 1: [2]int{33, 33}}
		fmt.Printf("%#v\n\n", dateList3)
		fmt.Printf("%#v\n\n", dateList3[0])
		fmt.Printf("%#v\n\n", dateList3[1][1])

		dateList4 := [1][1][1]int{0: [1][1]int{0: [1]int{99}}}
		fmt.Printf("%#v\n\n", dateList4[0][0][0])
	*/
	//切片--nil切片，没有进行初始化
	var names2 []string
	//类型
	fmt.Printf("%T\n\n", names2)
	//零值 就是 nil
	fmt.Printf("%#v\n\n", names2)

	//空切片，已经初始化，长度是0，所以值为空
	names2 = []string{}
	fmt.Printf("%#v\n\n", names2)

	//切片无需指定长度
	names3 := []string{"我", "爱", "中", "国", "吗"}
	fmt.Printf("%#v\n\n", names3)

	//切片初始化值,这时候长度确定，值也不会出现nil
	names4 := []string{1: "一", 5: "伍"}
	fmt.Printf("%#v\n\n", names4)

	//定义切片的第三种方式  make
	//make() 2个参数  指定切片类型，指定长度
	//make() 3个参数
	/*
		切片底层使用的是数组
		切片是一个结构器
		struct {
			array   内存中数组指针
			length  数组长度，存储元素的数量
			cap     数组容量
		}
		先申请10个字符串长度的元素切片，如果后面存储的元素个数超过10个，那么切片会重新申请一个新的数组
		新的数组的增长是依靠增长因子来袭实现，如果字节小于1024 会增以2的倍数进行增长，超过1024 的话增长因子为 0.25
		以10个元素为例，添加第11 个参数，实际切片长度为 11，容量为20
	*/

	names5 := make([]string, 5) //申请一个长度为5的切片，并且用0值来初始化
	fmt.Printf("%#v\n\n", names5)

	names5 = make([]string, 1, 10) //申请一个长度为1,容量为10 的切片，并且用0值来初始化
	fmt.Printf("%#v\n\n", names5)

	//切片访问,修改---通过索引
	names5[0] = "aa"
	fmt.Println(names5[0])

	//names5[2] = "aa"
	//fmt.Println(names5[2])
	//报错panic: runtime error: index out of range [2] with length 1
	//存在一定限制，如果在穿件之前指定了长度，那么访问长度之外的数据就会报错

	//查看切片长度len（）,容量 cap（）
	fmt.Printf("%#v   %#v \n\n", len(names5), cap(names5))

	//添加元素,就可以访问，添加的是长度，如果长度超过存储，就会复制转移，导致数据转移
	names5 = append(names5, "d")
	fmt.Printf("%#v   %#v \n\n", len(names5), cap(names5))
	names5[1] = "bb"
	fmt.Println(names5[1])
	fmt.Printf("添加元素\n\n")

	//切片遍历
	for m := 0; m < len(names5); m++ {
		fmt.Println(names5[m])
	}
	fmt.Printf("\n\n")
	//第二种方式
	for _, n := range names5 {
		fmt.Println(n)
	}
	fmt.Printf("\n\n")
	//切片复制
	/*
		深浅拷贝都是进行复制，区别在于复制出来的新对象与原来的对象在它们发生改变时，是否会相互影响，
		本质区别就是复制出来的对象与原对象是否会指向同一个地址。
		在Go语言，切片拷贝有三种方式：
			使用=操作符拷贝切片，这种就是浅拷贝
			使用[:]下标的方式赋值切片，这种也是浅拷贝
			使用Go语言的内置函数copy()进行切片拷贝，这种就是深拷贝，
	*/
	//长度相等，将后面的数据复制到前面，，，，copy（目的，原）
	//复制只会替换目的对应索引位置的数据
	aSlice := []string{"a", "b", "c"}
	bSlice := []string{"x", "y", "z", "w"}
	//copy(aSlice, bSlice)
	copy(bSlice, aSlice)
	fmt.Printf("%#v   \n%#v\n\n", aSlice, bSlice)

	//切片的切片操作   names5[start:end]
	cSlice := bSlice[2:3]
	fmt.Printf("%T,%#v,%v\n\n", cSlice, cSlice, cap(cSlice))

	//切片副作用    start <= end <= cap  ; new_cap = cap - start
	//如果start = end ，则是一个空切片
	//新生成的切片容量为  切片容量，减去切片的切片start的索引值
	//切片的切片，不新生成数组，共用原来的数组，初始地址，切片新指向第一个切片元素的地址，其他后面的不变
	//由于公用数组，后续对切片的操作，切片的切片也会继承操作结果
	bSlice[2] = "XXX"
	//上述只操作bSlice切片，但是cSlice也会随之变化
	fmt.Printf("%T,%#v,%v\n\n", cSlice, cSlice, cap(cSlice))
	//cSlice，变化，bSclice也会随之变化
	cSlice = append(cSlice, "YYY")
	fmt.Printf("%T,%#v,%v\n\n", bSlice, bSlice, cap(bSlice))

	//容量和存储位置共享，但是各个切片的长度不共享
	face1 := make([]string, 5, 10)
	face2 := face1[3:4]
	face3 := face1[3:4]
	fmt.Printf("%#v  %#v\n\n", len(face2), len(face3))
	face2 = append(face2, "d")
	fmt.Printf("%#v  %#v\n\n", len(face2), len(face3))

	//第二种切片的切片,使用限定容量的，
	//如果追加元素，会生成新数组，所以后续操作无影响
	//names5[start:end:max],max 是指定新切片的结束位置
	//cap = max - start
	face5 := []int{0, 1, 2, 3, 4, 5}
	face6 := face5[3:4:4]
	face6 = append(face6, 100)
	fmt.Printf("%#v\n%#v\n\n", face5, face6)

	//数组切片
	//start <= end <= len
	//因为数组指向的是同一个数组，添加元素，会导致原始数组值改变
	beta1 := [6]int{0, 1, 2, 3, 4, 5}
	beta2 := beta1[3:4]
	beta2 = append(beta2, 100)
	fmt.Printf("%#v\n%#v\n\n", beta1, beta2)

	//同样使用max，同样就会新建数组去复制操作了，不影响原始数据
	beta1 = [6]int{0, 1, 2, 3, 4, 5}
	beta2 = beta1[3:4:4]
	beta2 = append(beta2, 100)
	fmt.Printf("%#v\n%#v\n\n", beta1, beta2)

	//移除切片的数据--默认不支持
	//移除第一个或者最后一个元素，使用切片去操作
	good1 := []int{1, 2, 3, 4}
	//移除第一个元素
	good2 := good1[1:]
	fmt.Println(good2)
	//移除最后一个元素
	good3 := good1[:len(good1)-1]
	fmt.Println(good3)
	//移除中间一个元素
	//需要使用copy+切片,就是将数据整体往前提一格
	//用后切出的数据覆盖前切出的数据，然后去掉最后一个元素
	//例如，移除3
	copy(good1[2:], good1[3:])
	//copy ([3,4],[4])=> [4,4] => [1,2,4,4]
	fmt.Print(good1[:len(good1)-1])
	//[1,2,4,4]  => [1,2,4]
	fmt.Printf("\n\n")
	//移除多个元素也可以
	good1 = []int{1, 2, 3, 4}
	//[2,3,4]  [4]  => [4,3,4] => [1,4,3,4]
	copy(good1[1:], good1[3:])
	//[1,4,3,4] => [1,4]
	fmt.Print(good1[:len(good1)-2])

}
