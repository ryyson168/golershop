package pt

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ProductBase = cProductBase{}
)

type cProductBase struct{}

// =================== 管理端使用 =========================

// List 商品列表
func (c *cProductBase) List(ctx context.Context, req *pt.ProductListReq) (res *pt.ProductListRes, err error) {
	input := do.ProductIndexListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	/*
		whereExt := []*ml.WhereExt{}

		if !g.IsEmpty(req.ProductName) {
			whereExt = append(whereExt, &ml.WhereExt{
				Column: dao.ProductIndex.Columns().ProductName,
				Val:    "%" + req.ProductName + "%",
				Symbol: ml.LIKE,
			})
		}
		if !g.IsEmpty(req.ProductId) {
			whereExt = append(whereExt, &ml.WhereExt{
				Column: dao.ProductIndex.Columns().ProductId,
				Val:    req.ProductId,
				Symbol: model.EQ,
			})
		}
		if !g.IsEmpty(req.ProductNumber) {
			whereExt = append(whereExt, &ml.WhereExt{
				Column: dao.ProductIndex.Columns().ProductNumber,
				Val:    req.ProductNumber,
				Symbol: model.EQ,
			})
		}

		input.WhereExt = whereExt
	*/

	var result, error = service.ProductIndex().GetList(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// SaveProduct 保存商品
func (c *cProductBase) SaveProduct(ctx context.Context, req *pt.ProductSaveReq) (res *pt.ProductSaveRes, err error) {
	//初始化请求默认值
	input := model.SaveProductInput{}
	gconv.Scan(req, &input)

	var result, error = service.ProductBase().SaveProdcut(ctx, &input)

	if error != nil {
		err = error
	}

	//新增成功后
	//商品统计

	//店铺统计

	//发送通知消息

	res = &pt.ProductSaveRes{
		ProductId: result,
	}

	return
}

// EditState 编辑状态
func (c *cProductBase) EditState(ctx context.Context, req *pt.ProductEditStateReq) (res *pt.ProductEditStateRes, err error) {
	input := do.ProductIndex{}
	gconv.Scan(req, &input)

	var result, error = service.ProductIndex().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductEditStateRes{
		ProductId: result,
	}

	return
}

// RemoveProduct 删除商品
func (c *cProductBase) RemoveProduct(ctx context.Context, req *pt.ProductRemoveReq) (res *pt.ProductRemoveRes, err error) {
	var _, error = service.ProductBase().RemoveProdcut(ctx, req.ProductId)

	if error != nil {
		err = error
	}

	res = &pt.ProductRemoveRes{}

	//todo 发送通知消息

	return
}

// GetProduct 商品信息
func (c *cProductBase) GetProduct(ctx context.Context, req *pt.ProductDateReq) (res *pt.ProductDateRes, err error) {
	data, err := service.ProductBase().GetProduct(ctx, req.ProductId)

	gconv.Scan(data, &res)

	return res, err
}
