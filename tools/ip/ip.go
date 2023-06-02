package ip

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

// GetRequestIP .. 获取客户端 IP 地址
func GetRequestIP(r *http.Request) (ip string, err error) {
	// 从请求头部的 X-REAL-IP 获取 IP
	ip = r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	// 从请求头部的 X-FORWARDED-FOR 获取 IP
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	// 从请求头部的 RemoteAddr 获取 IP
	ip, _, err = net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	return "", fmt.Errorf("正确ip获取失败")
}
