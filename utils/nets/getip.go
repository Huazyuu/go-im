package nets

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net"
)

func GetLocalIP() []string {
	interfaces, err := net.Interfaces()
	if err != nil {
		logx.Error("获取网卡信息错误")
		return nil
	}

	res := make([]string, 0)

	for _, inter := range interfaces {
		addrs, err := inter.Addrs()
		if err != nil {
			logx.Error("获取ip出错,err:", err)
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					res = append(res, ipNet.IP.String()+"\n")
				}
			}
		}
	}
	return res
}
