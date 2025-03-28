package etcd

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"server/core"
)

func GetAddress(etcdAddr, serviceName string) string {
	client := core.InitEtcd(etcdAddr)
	res, err := client.Get(context.Background(), serviceName)
	if err != nil || res.Kvs == nil {
		logx.Errorf("%s 不存在", serviceName)
		return ""
	}
	return string(res.Kvs[0].Value)
}
