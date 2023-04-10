package main

import (
	"fmt"
	"strconv"
	"strings"
)

//做一个命令行管理--任务管理器
//用户管理

//1、函数，输入，输出，复合数据结构，基本数据类型
//2、了解流程（对数据的操作流程、增删改查）

//1、任务输出--添加任务
//2、任务列表--任务查询
//3、任务修改--
//4、任务删除
//5、详情

// 任务
// 任务名称，开始时间，结束时间，状态，负责人
// ID,name,star_time,end_time,status,owner
// []map[string][string]
var todos = []map[string]string{
	{"id": "1", "name": "复习", "startTime": "18:00", "endTime": "19:00", "status": "完成", "owner": "limi"},
	{"id": "3", "name": "打架", "startTime": "18:00", "endTime": "19:00", "status": "完成", "owner": "limi"},
}

const (
	id        = "id"
	name      = "name"
	startTime = "startTime"
	endTime   = "endTime"
	status    = "status"
	owner     = "owner"
)
const (
	statusNew     = "未执行"
	statusCompele = "完成"
)

func genId() int {
	var rt int
	for _, todo := range todos {
		todoId, _ := strconv.Atoi(todo["id"])
		if rt < todoId {
			rt = todoId
		}

	}
	return rt + 1
}

func newTask() map[string]string {
	//id生成（todolist 中最大的ID+1）
	task := make(map[string]string)
	task[id] = strconv.Itoa(genId())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = ""
	task[owner] = ""
	return task

}
func printTask(task map[string]string) {
	fmt.Println(strings.Repeat("+", 20))
	fmt.Println("ID:", task[id])
	fmt.Println("任务名称:", task[name])
	fmt.Println("开始时间:", task[startTime])
	fmt.Println("结束时间:", task[endTime])
	fmt.Println("状态:", task[status])
	fmt.Println("负责人信息:", task[owner])
	fmt.Println(strings.Repeat("+", 20))
}

func input(prompt string) string {
	var text string
	fmt.Printf("%v\n", prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

func makeTask() {
	task := newTask()
	fmt.Println("请输入任务信息：")
	task["name"] = input("请输入任务名称")
	task["startTime"] = input("请输入开始时间")
	task["endTime"] = input("请输入结束时间")
	task["status"] = input("请输入状态")
	task["owner"] = input("请输入负责人信息")
	//fmt.Println(todos)
	todos = append(todos, task)
}

func queryTask() {
	q := input("请输入查询信息")
	//all，显示所有信息

	for _, todo := range todos {
		if q == "all" || strings.Contains(todo[name], q) {
			printTask(todo)
		}
	}
}

func main() {
EXIT:
	for {
		switch input("请输入操作 (add /query / exit) ") {
		case "add":
			makeTask()
		case "query":
			queryTask()
		case "exit":
			//return  也可以
			break EXIT
			//break + 标签
		default:
			fmt.Println("指令错误")
		}
	}

}
