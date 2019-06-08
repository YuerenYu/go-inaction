package main

import (
	_ "go-inaction/src/chapter2/matchers"
	"go-inaction/src/chapter2/search"
	"log"
	"os"
)

// init在main之前调用
func init(){
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main(){
	search.Run("球队")
}
