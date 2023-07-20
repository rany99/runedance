package comment

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
	"runedance/comment/dao/dal"
	"runedance/comment/dao/redis"
	"runedance/comment/pulsar"
	"runedance/common/config"
	"runedance/common/constant"
	commentproto "runedance/kitexGen/kitex_gen/commentproto/commentservice"
	filter "runedance/pkg/ditryWords"
	"runedance/pkg/middleware"
	"runedance/pkg/tracer"
)

func Init() {
	config.InitConfig()
	dal.Init()
	redis.Init()
	pulsar.Init()
	filter.Init()
	tracer.InitJaeger(constant.CommentDomainServiceName)
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.Server.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.Server.CommentServiceAddr)
	if err != nil {
		panic(err)
	}

	svr := commentproto.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.CommentDomainServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                                     // middleWare
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
