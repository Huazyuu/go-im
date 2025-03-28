package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"server/common/etcd"
)

// BaseResponse 基础响应结构体，用于统一接口响应格式
type BaseResponse struct {
	Code int    `json:"code"` // 响应状态码（0表示成功，非0表示失败）
	Msg  string `json:"msg"`  // 响应消息描述
	Data any    `json:"data"` // 响应数据内容
}

// FilResponse 快速构造并发送错误响应
// msg: 错误信息内容
// res: HTTP响应写入器
func FilResponse(msg string, res http.ResponseWriter) {
	response := BaseResponse{Code: 7, Msg: msg}
	byteData, _ := json.Marshal(response)
	res.Write(byteData)
}

// auth 调用认证服务进行请求鉴权
// authAddr: 认证服务地址
// res: HTTP响应写入器，用于直接返回错误
// req: HTTP请求对象
// 返回值: 认证通过返回true，否则false
func auth(authAddr string, res http.ResponseWriter, req *http.Request) (ok bool) {
	// 构造认证请求
	authReq, _ := http.NewRequest("POST", authAddr, nil)
	authReq.Header = req.Header // 复制原始请求头

	// 从URL参数获取token并添加到请求头（兼容URL参数传递token的场景）
	token := req.URL.Query().Get("token")
	if token != "" {
		authReq.Header.Set("token", token)
	}

	// 将当前请求路径添加到请求头供认证服务校验权限
	authReq.Header.Set("validpath", req.URL.Path)

	// 发送认证请求
	authRes, err := http.DefaultClient.Do(authReq)
	if err != nil {
		logx.Error("认证服务请求失败: ", err)
		FilResponse("认证服务错误", res)
		return
	}
	defer authRes.Body.Close()

	// 解析认证响应
	type AuthResponse struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data *struct {
			UserID uint `json:"userID"` // 用户ID
			Role   int  `json:"role"`   // 用户角色
		} `json:"data"`
	}
	var authResponse AuthResponse

	byteData, _ := io.ReadAll(authRes.Body)
	if err := json.Unmarshal(byteData, &authResponse); err != nil {
		logx.Error("认证响应解析失败: ", err)
		FilResponse("认证服务错误", res)
		return
	}

	// 认证失败直接返回错误信息
	if authResponse.Code != 0 {
		res.Write(byteData)
		return false
	}

	// 认证成功时设置用户信息到请求头
	if authResponse.Data != nil {
		req.Header.Set("User-ID", fmt.Sprintf("%d", authResponse.Data.UserID))
		req.Header.Set("Role", fmt.Sprintf("%d", authResponse.Data.Role))
	}
	return true
}

// Proxy 反向代理结构体，实现http.Handler接口
type Proxy struct{}

// ServeHTTP 处理所有HTTP请求的核心方法
func (Proxy) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// 使用正则表达式从URL路径中提取服务名称
	// 示例路径: /api/user/profile -> 提取"user"
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if len(matches) != 2 {
		FilResponse("无效的请求路径", res)
		return
	}
	service := matches[1]

	// 从ETCD获取目标服务的真实地址
	addr := etcd.GetAddress(config.Etcd, service+"_api")
	if addr == "" {
		logx.Errorf("服务发现失败: %s", service)
		FilResponse("服务不可用", res)
		return
	}

	// 获取认证服务地址并构造认证URL
	authAddr := etcd.GetAddress(config.Etcd, "auth_api")
	authUrl := fmt.Sprintf("http://%s/api/auth/authentication", authAddr)

	// 打印访问日志
	logx.Infof("客户端地址请求(%s)经过api网关:  代理到: http://%s%s", req.RemoteAddr, addr, req.URL.Path)

	// 进行请求鉴权
	if !auth(authUrl, res, req) {
		return // 认证失败已通过res返回响应
	}

	// 创建反向代理并转发请求
	remote, _ := url.Parse(fmt.Sprintf("http://%s", addr))
	reverseProxy := httputil.NewSingleHostReverseProxy(remote)
	reverseProxy.ServeHTTP(res, req) // 执行反向代理
}

var configFile = flag.String("f", "gateway-go.yaml", "配置文件路径") // 命令行参数-f指定配置文件

// Config 服务配置结构
type Config struct {
	Addr string       `json:"addr"` // 网关监听地址
	Etcd string       `json:"etcd"` // ETCD连接地址
	Log  logx.LogConf `json:"log"`  // 日志配置
}

var config Config // 全局配置实例

func main() {
	flag.Parse()

	// 加载配置文件
	conf.MustLoad(*configFile, &config)

	// 初始化日志系统
	logx.SetUp(config.Log)

	// 启动HTTP服务
	fmt.Printf("API网关启动中，监听地址: %s\n", config.Addr)
	proxy := Proxy{}
	if err := http.ListenAndServe(config.Addr, proxy); err != nil {
		logx.Errorf("网关启动失败: %v", err)
		panic(err)
	}
}
