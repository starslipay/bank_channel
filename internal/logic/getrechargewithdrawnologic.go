package logic

import (
	"context"

	"github.com/starslipay/bank_channel/bank_channel_pb"
	"github.com/starslipay/bank_channel/internal/svc"
	"github.com/starslipay/paycomm/xerror"
	"github.com/starslipay/trade_id_mgr/trade_id_mgr_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRechargeWithdrawNoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRechargeWithdrawNoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRechargeWithdrawNoLogic {
	return &GetRechargeWithdrawNoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRechargeWithdrawNoLogic) GetRechargeWithdrawNo(in *bank_channel_pb.GetRechargeWithdrawNoReq) (*bank_channel_pb.GetRechargeWithdrawNoRsp, error) {
	// 从 trade_id_mgr 获取单号, sp_id/scene_id 与 trade_itg 保持一致
	tradeIdRsp, err := l.svcCtx.TradeIdMgr.GenTradeId(l.ctx, &trade_id_mgr_pb.GenTradeIdReq{
		SpId:    "1000000000",
		Uid:     in.Uid,
		SceneId: 1,
	})
	if err != nil {
		return nil, xerror.HandleRPCError(err, "TradeIdMgr.GenTradeId")
	}

	return &bank_channel_pb.GetRechargeWithdrawNoRsp{
		TransactionId: tradeIdRsp.TradeId,
	}, nil
}
