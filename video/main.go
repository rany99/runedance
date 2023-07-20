package video

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
	"runedance/common/config"
	"runedance/common/constant"
	videoproto "runedance/kitexGen/kitex_gen/videoproto/videoservice"
	"runedance/pkg/middleware"
	"runedance/pkg/tracer"
	"runedance/video/dao/dal"
	"runedance/video/dao/redis"
	"runedance/video/pulsar"
	"runedance/video/service"
)

func Init() {
	config.InitConfig()
	dal.Init()
	redis.Init()
	pulsar.Init()
	service.Tracer, service.Closer = tracer.InitJaeger(constant.VideoDomainServiceName)
}

func main() {
	Init()
	defer service.Closer.Close()
	r, err := etcd.NewEtcdRegistry([]string{config.Server.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", config.Server.VideoServiceAddr)
	if err != nil {
		panic(err)
	}
	svr := videoproto.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.VideoDomainServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                                   // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithRegistry(r),                                             // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
