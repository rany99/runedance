package message

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
	messageproto "runedance/kitexGen/kitex_gen/messageproto/messageservice"
	"runedance/message/dao/dal"
	"runedance/message/dao/redis"
	"runedance/message/pulsar"
	"runedance/pkg/middleware"
	"runedance/pkg/tracer"
)

func Init() {
	config.InitConfig()
	redis.Init()
	dal.Init()
	pulsar.Init()
	tracer.InitJaeger(constant.MessageDomainServiceName)
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.Server.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.Server.MessageServiceAddr)
	if err != nil {
		panic(err)
	}

	svr := messageproto.NewServer(new(MessageServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.MessageDomainServiceName}), // server name
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
