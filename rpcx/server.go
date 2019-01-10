package main

import (
	"flag"
	"log"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"hank.com/etcdrpcx/example"
)

var (
	addr1    = flag.String("addr1", "localhost:8973", "server address1")
	addr2    = flag.String("addr2", "localhost:8972", "server address2")
	addr3    = flag.String("addr3", "localhost:8971", "server address3")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
)


func main() {
	flag.Parse()

	go createGoodsServer(*addr1, *etcdAddr, example.NewPreKey("goods"), "goods", new(example.Goods))
	go createGoodsServer(*addr2, *etcdAddr, example.NewPreKey("goods"), "goods", new(example.Goods))
	go createGoodsServer(*addr3, *etcdAddr, example.NewPreKey("order"), "order", new(example.Order))

	select {}
}

func createGoodsServer(addr string, etcdAddr string, basepath string, registerName string, regidterdoor interface{}) {
	s := server.NewServer()

	r := NewEtcRegistryPlugin(addr, etcdAddr, basepath)

	//启动
	err := r.Start()

	if err != nil {
		log.Fatal(err)
	}

	//加入插件
	s.Plugins.Add(r)

	s.RegisterName(registerName, regidterdoor, "")
	s.Serve("tcp", addr)
}

//NewEtcRegistryPlugin -
func NewEtcRegistryPlugin(addr string, etcaddr string, basepath string) *serverplugin.EtcdRegisterPlugin {

	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		EtcdServers:    []string{etcaddr},
		BasePath:       basepath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}

	return r
}
