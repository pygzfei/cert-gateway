package cert

import (
	"context"
	"github.com/pygzfei/issuer-gateway/issuer/internal/database/entity"
	"github.com/pygzfei/issuer-gateway/issuer/internal/svc"
	"github.com/pygzfei/issuer-gateway/issuer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDomainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDomainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDomainLogic {
	return &AddDomainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDomainLogic) AddDomain(req *types.AddDomainReq) (resp *types.AddOrRenewCertificateResp, err error) {

	db := l.svcCtx.DB.Create(&entity.Cert{
		Email:  req.Email,
		Domain: req.Domain,
		Target: req.Target,
	})

	if db.RowsAffected == 0 || db.Error != nil {
		return nil, db.Error
	}
	return &types.AddOrRenewCertificateResp{}, nil
}
