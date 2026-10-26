package svc

import (
	"github.com/starslipay/bank_channel/internal/config"
	"github.com/starslipay/trade_id_mgr/trade_id_mgr_pb"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	TradeIdMgr trade_id_mgr_pb.TradeIdMgrClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		TradeIdMgr: trade_id_mgr_pb.NewTradeIdMgrClient(zrpc.MustNewClient(c.TradeIdMgrRpcConfig).Conn()),
	}
}
