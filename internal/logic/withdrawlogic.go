package logic

import (
	"context"

	"github.com/starslipay/bank_channel/bank_channel_pb"
	"github.com/starslipay/bank_channel/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawLogic {
	return &WithdrawLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WithdrawLogic) Withdraw(in *bank_channel_pb.WithdrawReq) (*bank_channel_pb.WithdrawRsp, error) {
	// mock接口: 不真实调用银行, 直接返回成功
	l.Infof("mock withdraw success, transaction_id=%s, amount=%d, bind_serial_no=%s",
		in.TransactionId, in.Amount, in.BindSerialNo)

	return &bank_channel_pb.WithdrawRsp{
		TransactionId: in.TransactionId,
	}, nil
}
