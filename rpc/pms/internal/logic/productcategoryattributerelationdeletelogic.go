package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryAttributeRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationDeleteLogic {
	return &ProductCategoryAttributeRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryAttributeRelationDeleteLogic) ProductCategoryAttributeRelationDelete(in *pms.ProductCategoryAttributeRelationDeleteReq) (*pms.ProductCategoryAttributeRelationDeleteResp, error) {
	err := l.svcCtx.PmsProductCategoryAttributeRelationModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductCategoryAttributeRelationDeleteResp{}, nil
}
