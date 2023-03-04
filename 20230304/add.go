package main

import (
	"fmt"
)

var todos = make([]map[string]string, 0)

const (
	name      = "name"
	startTime = "start_time"
	endTime   = "end_time"
	status    = "status"
	owner     = "user"
)

func newTask() map[string]string {
	//id生成（todolist 中最大的ID+1）
	task := make(map[string]string)
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = ""
	task[owner] = ""
	return task

}

func main() {
	task := newTask()
	var text string
	fmt.Println("请输入任务信息：")
	fmt.Println("任务名称：")
	fmt.Scan(&text)
	task["name"] = text
	fmt.Println("开始时间：")
	fmt.Scan(&text)
	task["startTime"] = text
	fmt.Println("结束时间：")
	fmt.Scan(&text)
	task["endTime"] = text
	fmt.Println("状态")
	fmt.Scan(&text)
	task["status"] = text
	fmt.Println("负责人信息")
	fmt.Scan(&text)
	task["owner"] = text

	todos = append(todos, task)
	fmt.Println(todos)
}
