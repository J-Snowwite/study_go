package main

import "fmt"

// n 个盘子   star（开始）    end（终点）    temp（借助）

// n  start   ->temp    ->end
// 把n个盘子从start上移动到 end 上需要 借助temp

// n-1  start  ->end  ->temp
//把 n-1 个盘子移动到  temp上

// start -> end
//然后第n个盘子，从start  移动到 end 上

// n-1   temp  -> start  -> end
//最后将 n-1个盘子 ，从temp上移动到 end上

// 终止条件   n=1 直接返回

func hnt(start, end, temp string, layer int) {
	//表示从 start 移动到end 上，需要借助 temp
	if layer == 1 {
		fmt.Println(start, "->", end)
		return
		//无返回值，表示函数结束
	}
	hnt(start, temp, end, layer-1)
	//第一步，将n-1层，从start 上移动到temp上，需要借助end
	fmt.Println(start, "->", end)
	hnt(temp, end, start, layer-1)
	//第二步，将n从 start 上移动到 end上
}

func main() {
	hnt("塔1", "塔3", "塔2", 4)
}
