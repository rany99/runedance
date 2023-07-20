package user

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
	userproto "runedance/kitexGen/kitex_gen/userproto/userservice"
	"runedance/pkg/middleware"
	"runedance/pkg/tracer"
	"runedance/user/dao/dal"
	"runedance/user/dao/redis"
	"runedance/user/pulsar"
)

func Init() {
	config.InitConfig()
	dal.Init()
	redis.Init()
	pulsar.Init()
	tracer.InitJaeger(constant.UserDomainServiceName)
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.Server.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.Server.UserServiceAddr)
	if err != nil {
		panic(err)
	}

	svr := userproto.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.UserDomainServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                                  // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		//server.WithRegistryInfo(regInfo),
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
