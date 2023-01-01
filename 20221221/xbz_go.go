// 选班长，统计 小红，小李，小新 的得票次数
package main

import "fmt"

func main() {
	var xbz map[string]int
	xbz = map[string]int{"小红": 0, "小李": 0, "小新": 0}
	xp := []string{"小红", "小李", "小新", "小新", "小新", "小新", "小红", "小李"}
	for _, v := range xp {
		switch v {
		case "小红":
			xbz["小红"]++
		case "小新":
			xbz["小新"]++
		case "小李":
			xbz["小李"]++
		}
	}
	fmt.Printf("小红的选票：%v\n", xbz["小红"])
	fmt.Printf("小新的选票：%v\n", xbz["小新"])
	fmt.Printf("小李的选票：%v\n", xbz["小李"])

}
