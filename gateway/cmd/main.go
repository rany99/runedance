package cmd

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"runedance/common/config"
	"runedance/gateway/api/authorization"
	"runedance/gateway/rpc"
	"runedance/pkg/minio"
)

func Init() {
	config.InitConfig()
	authorization.Init()
	minio.Init()
	InitInjectModule()
	rpc.Init()
}

func main() {
	Init()
	r := gin.New()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	if config.Server.RunMode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	register(r)
	if config.Server.PprofSwitch == "on" {
		pprof.Register(r)
	}
	if err := r.Run(config.Server.HttpPort); err != nil {
		klog.Fatal(err)
	}
}
