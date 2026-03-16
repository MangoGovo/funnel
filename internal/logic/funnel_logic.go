// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"funnel/internal/svc"
	"funnel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FunnelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFunnelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FunnelLogic {
	return &FunnelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FunnelLogic) Funnel(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
