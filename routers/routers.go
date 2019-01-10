package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"hank.com/etcdrpcx/rpcx/client"
)

func init() {
	// api 服务
	beego.Any("/api/*", func(c *context.Context) {
		client.ServerJSON(c)
	})
}
