package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// 没有nginx环境可以使用我的 go gateway

var ServiceMap = map[string]string{
	"auth":  "http://127.0.0.1:20261",
	"user":  "http://127.0.0.1:20262",
	"chat":  "http://127.0.0.1:20263",
	"group": "http://127.0.0.1:20264",
}

type Data struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (d Data) toJson() []byte {
	b, _ := json.Marshal(d)
	return b
}

func gateway(res http.ResponseWriter, req *http.Request) {
	p := req.URL.Path
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	list := regex.FindStringSubmatch(p)
	if len(list) != 2 {
		_, _ = res.Write(Data{Code: 7, Msg: "服务错误"}.toJson())
		return
	}
	addr, ok := ServiceMap[list[1]]
	if !ok {
		_, _ = res.Write(Data{Code: 7, Msg: "服务错误"}.toJson())
		return
	}

	url := addr + req.URL.String()

	// 读取请求体
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		_, _ = res.Write(Data{Code: 7, Msg: "服务错误"}.toJson())
		return
	}
	// 关闭原请求体
	req.Body.Close()

	// 创建新的请求体
	proxyRequest, _ := http.NewRequest(req.Method, url, io.NopCloser(strings.NewReader(string(bodyBytes))))
	proxyRequest.Header = req.Header
	remoteAddr := strings.Split(req.RemoteAddr, ":")
	if len(remoteAddr) != 2 {
		_, _ = res.Write(Data{Code: 7, Msg: "服务错误"}.toJson())
		return
	}
	ip := remoteAddr[0]
	port := remoteAddr[1]
	fmt.Printf("%s %s =>  %s\n", ip, port, url)
	proxyRequest.Header.Set("X-Forwarded-For", ip)

	proxyResponse, err := http.DefaultClient.Do(proxyRequest)
	if err != nil {
		_, _ = res.Write(Data{Code: 7, Msg: "服务错误"}.toJson())
		return
	}
	defer proxyResponse.Body.Close()

	io.Copy(res, proxyResponse.Body)
	return
}
func main() {
	// 回调函数
	http.HandleFunc("/", gateway)
	// 绑定服务
	fmt.Printf("im_gateway 运行在：%s\n", "http://127.0.0.1:9001")
	http.ListenAndServe(":9001", nil)
}
