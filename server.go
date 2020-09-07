package main

import (
	"domain_proxy/handler"
	"domain_proxy/handler/proxy"
	"domain_proxy/handler/redis"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/gorilla/mux"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8082, "指定服务监控的端口，例如 :8082 ，注意，需要带冒号")
}

func main() {
	redis.Init()

	r := mux.NewRouter()

	r.Handle("/proxy/add", handler.NewBaseHandler(proxy.NewAddHandler(), nil, nil))

	r.PathPrefix("/").Handler(handler.NewBaseHandler(proxy.NewProxyHandler(), nil, nil))

	fmt.Println("server start at :", port)
	http.Handle("/", r)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
