package main

import (
	"fmt"
	"strings"
)

var todos = []map[string]string{
	{"id": "1", "name": "复习", "start_time": "18:00", "end_time": "19:00", "status": "完成", "user": "limi"},
	{"id": "3", "name": "打架", "start_time": "18:00", "end_time": "19:00", "status": "完成", "user": "limi"},
}

const (
	id        = "id"
	name      = "name"
	startTime = "start_time"
	endTime   = "end_time"
	status    = "status"
	owner     = "user"
)

func printTask(task map[string]string) {
	fmt.Println(strings.Repeat("+", 20))
	fmt.Println("+++++++++++++++++++++++++++++")
	fmt.Println("任务名称:", task[name])
	fmt.Println("开始时间:", task[startTime])
	fmt.Println("+++++++++++++++++++++++++++++")
}

func main() {
	var text string
	fmt.Println("请输入查询信息")
	fmt.Scan(&text)

	for _, todo := range todos {
		if strings.Contains(todo[name], text) {
			printTask(todo)
		}
	}
}
