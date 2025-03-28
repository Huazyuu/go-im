package etcd

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/netx"
	"server/core"
	"strings"
)

// PutAddress 地址上送服务
func PutAddress(etcdAddr, serviceName, serviceAddr string) {
	list := strings.Split(serviceAddr, ":")
	if len(list) != 2 {
		logx.Errorf("%s 错误的地址", serviceAddr)
		return
	}
	if list[0] == "0.0.0.0" {
		ip := netx.InternalIp()
		serviceAddr = strings.ReplaceAll(serviceAddr, "0.0.0.0", ip)
	}

	client := core.InitEtcd(etcdAddr)
	_, err := client.Put(context.Background(), serviceName, serviceAddr)
	if err != nil {
		logx.Errorf("地址上送失败 %s", err.Error())
		return
	}
	logx.Infof("地址上送成功 %s  %s", serviceName, serviceAddr)
}
