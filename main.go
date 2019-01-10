package main

import (
	"github.com/astaxie/beego"
	// 再初始化路由
	_ "hank.com/etcdrpcx/routers"
)

func main(){
	beego.Run()
}
