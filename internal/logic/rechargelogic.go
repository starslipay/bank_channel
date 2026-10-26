package logic

import (
	"context"

	"github.com/starslipay/bank_channel/bank_channel_pb"
	"github.com/starslipay/bank_channel/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RechargeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRechargeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RechargeLogic {
	return &RechargeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RechargeLogic) Recharge(in *bank_channel_pb.RechargeReq) (*bank_channel_pb.RechargeRsp, error) {
	// mock接口: 不真实调用银行, 直接返回成功
	l.Infof("mock recharge success, transaction_id=%s, amount=%d, bind_serial_no=%s",
		in.TransactionId, in.Amount, in.BindSerialNo)

	return &bank_channel_pb.RechargeRsp{
		TransactionId: in.TransactionId,
	}, nil
}
