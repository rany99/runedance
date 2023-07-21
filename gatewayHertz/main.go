// Code generated by hertz generator.

package main

import (
	"runedance/gatewayHertz/biz/rpc"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	rpc.Init()
}
func main() {
	Init()
	// 127.0.0.1

	h := server.Default(server.WithHostPorts("0.0.0.0:8080"), server.WithMaxRequestBodySize(233333333))
	//h.Use(mw.MyJWT())

	register(h)
	h.Spin()
}
