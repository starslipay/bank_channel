package main

import (
	"flag"
	"fmt"

	"github.com/starslipay/bank_channel/bank_channel_pb"
	"github.com/starslipay/bank_channel/internal/config"
	"github.com/starslipay/bank_channel/internal/metrics"
	"github.com/starslipay/bank_channel/internal/middleware"
	"github.com/starslipay/bank_channel/internal/server"
	"github.com/starslipay/bank_channel/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/bankchannel.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		bank_channel_pb.RegisterBankChannelServer(grpcServer, server.NewBankChannelServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(metrics.UnaryMetricInterceptor)
	// gRPC 访问日志拦截器(按字段名脱敏)
	if c.AccessLog.Enable {
		s.AddUnaryInterceptors(middleware.AccessLogUnaryInterceptor(
			middleware.NewSensitiveFields(c.AccessLog.SensitiveFields),
		))
	}
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
