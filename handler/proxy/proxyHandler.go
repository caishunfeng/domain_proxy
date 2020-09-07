package proxy

import (
	"domain_proxy/handler/redis"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyHandler struct {
}

func NewProxyHandler() *ProxyHandler {
	return &ProxyHandler{}
}

func (this *ProxyHandler) Handle(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	host := r.Host

	redisClient, err := redis.CreateClient()
	if err != nil {
		w.Write([]byte("连接redis异常"))
		return
	}

	stringCmd := redisClient.Get(DOMAIN_PROXY_PREFIX + host)
	if stringCmd.Err() != nil {
		w.Write([]byte("redis get 异常"))
		return
	}

	fmt.Printf("get domain success,domain:%s,ip:%s\n", host, stringCmd.String())

	remoteUrl := "http://" + stringCmd.Val()

	remote, err := url.Parse(remoteUrl)
	if err != nil {
		w.Write([]byte("转发异常,remoteUrl:" + remoteUrl))
		return
	}

	// 转发
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}
