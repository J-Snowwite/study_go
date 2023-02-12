package main

import (
	"bytes"
	"fmt"
)

func main() {

	//  网址：https://golang.google.cn/pkg/bytes/

	//前面的切片是否包含后面的切片
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))
}
