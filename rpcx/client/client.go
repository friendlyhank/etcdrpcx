// go run -tags etcd client.go
package client

import (
	"encoding/json"
	"errors"
	"flag"
	"strings"
	"hank.com/etcdrpcx/example"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/smallnest/rpcx/client"
)

var (
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
)

//ServerJSON -
func ServerJSON(c *context.Context) {
	logs.Info("%v", c.Request.URL.Path)
	NewTtcDiscovery(c)
}

//NewTtcDiscovery -
func NewTtcDiscovery(c *context.Context) {
	flag.Parse()

	var (
		servername string
		methodname string
		err        error
	)
	//GetServerName
	servername, err = FindServer(c.Request.URL.Path)
	if nil != err {
		logs.Info("%v", err)
		c.ResponseWriter.Write([]byte("服务名获取出错"))
		return
	}

	methodname, err = FindMethod(c.Request.URL.Path)
	if nil != err {
		logs.Info("%v", err)
		c.ResponseWriter.Write([]byte("方法名获取出错"))
		return
	}

	logs.Info("SercerName|%v|MethodName|%v|", servername, methodname)

	d := client.NewEtcdDiscovery(example.NewPreKey(servername), servername, []string{*etcdAddr}, nil)
	xclient := client.NewXClient(servername, client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	req := &example.Req{}

	reply := &example.Res{}
	err = xclient.Call(c.Request.Context(), methodname, req, reply)
	if err != nil {
		c.ResponseWriter.Write([]byte("Fail to call"))
		return
	}

	logs.Info("%v", reply)

	rebyte, _ := json.Marshal(reply)

	c.ResponseWriter.Write(rebyte)
}

//FindServer -
func FindServer(api string) (servername string, err error) {
	logs.Info("|api|%v", api)
	if sint := strings.Index(api, "api/goods"); sint != -1 {
		return "goods", nil
	}

	if sint := strings.Index(api, "api/order"); sint != -1 {
		return "order", nil
	}

	return "", errors.New("No Found")
}


//FindMethod -
func FindMethod(api string) (method string, err error) {
	arrs := strings.Split(api, "/")
	if len(arrs) != 4 {
		return "", errors.New("Not The Current API")
	}

	return arrs[3], nil
}
