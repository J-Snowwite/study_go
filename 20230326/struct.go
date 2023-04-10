package main

import (
	"fmt"
	"time"
)

// 声明一个结构体 struct(结构体)
// 结构体名称是大写的话，可以在本package外访问。如果是小写，只能在内部访问
type user struct {
	//内部字段，大写可以跨package，小写不可以
	id, id2 int     //多个相同的变量可以写在一行中
	score   float32 //如果没有值，那么每个变量取类型的零值
}

type user2 struct {
	num1       int
	enrpllment time.Time //time本身也是一个结构体---结构中存在函数全为0打印出现的值
}

type user3 struct {
	Char1 string
	Char2 string
}

type People struct {
	Name string
	Sex  byte
}
type UserMap map[int]People

//嵌套，结构体里面的数据是映射，映射的值是结构体

func (um UserMap) Get(id int) People {
	return um[id]
}
func (um UserMap) Say(id int) {
	delete(um, id)
}

var stu struct {
	//创建一个匿名结构体  stu是变量名
	Id int
}

//访问方式stu.Id

func (user1 user3) hello(man string) string {
	return "hellow" + man + ", i m" + user1.Char1
}

//传参--参数是结构体---两个hello 方法等价

func hello(man string, user1 user3) string {
	return "hellow " + man + ", i m " + user1.Char1
}

func main() {
	//结构体短声明--赋值为空
	u := user{}
	fmt.Println(u)

	//结构体赋值，使用定义字段赋值时，可以不按照顺序来，也可以不赋值
	//结构体赋值，如果不写字段，但是需要和结构体定义里字段顺序一致
	u2 := user{score: 4, id: 3}
	fmt.Println(u2)
	u3 := user{7, 8, 9}
	fmt.Println(u3)

	u4 := user2{}
	fmt.Println(u4)
	//单独给结构体内的某个字段赋值
	u2.id2 = 5
	fmt.Println(u2)
	//单独访问结构体内的某个值
	fmt.Println(u2.id)

	ws := user3{Char1: "xjt"}
	fmt.Println(ws.hello("xjt"))
	hello("xjt", ws)

}
