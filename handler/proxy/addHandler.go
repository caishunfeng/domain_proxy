package proxy

import (
	"domain_proxy/handler/redis"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type AddHandler struct {
}

func NewAddHandler() *AddHandler {
	return &AddHandler{}
}

func (this *AddHandler) Handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	domain := r.Form.Get("domain")
	port := r.Form.Get("port")
	expireStr := r.Form.Get("expire")

	ip := clientIP(r)
	if ip == "" {
		w.Write([]byte("获取ip失败"))
		return
	}

	if port == "" {
		w.Write([]byte("port参数为空"))
		return
	}
	if domain == "" {
		w.Write([]byte("domain参数为空"))
		return
	}
	if !strings.HasSuffix(domain, r.Host) {
		w.Write([]byte("domain形式应该为xxx." + r.Host))
		return
	}
	if expireStr == "" {
		expireStr = "12"
	}
	expire, err := strconv.Atoi(expireStr)
	if err != nil {
		w.Write([]byte("expire参数应该为整数"))
		return
	}

	redisClient, err := redis.CreateClient()
	if err != nil {
		w.Write([]byte("连接redis异常"))
		return
	}

	boolCmd := redisClient.Set(DOMAIN_PROXY_PREFIX+domain, ip+":"+port, time.Duration(expire)*time.Hour)
	if boolCmd.Err() != nil {
		w.Write([]byte("redis set 异常"))
		return
	}

	fmt.Printf("set success domain:%s,port:%s,ip:%s\n", domain, port, ip)

	w.Header().Add("Access-Control-Allow-Origin", "*")

	w.Write([]byte(fmt.Sprintf("设置成功，内网域名：%v,有效期为%d小时", domain, expire)))
}

func clientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
