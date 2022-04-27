package base

import (
	"context"
	"mathapi/api"
	"mathapi/model"
)

type Service interface {
	CalcMin(ctx context.Context, req model.MathRequest) ([]int, error)
	CalcMax(ctx context.Context, req model.MathRequest) ([]int, error)
	CalcAvg(ctx context.Context, req model.MathRequest) (float64, error)
	CalcMedian(ctx context.Context, req model.MathRequest) (float64, error)
	CalcPercentile(ctx context.Context, req model.MathRequest) (int, error)
}
type baseService struct {
}

func NewService() Service {
	return baseService{}
}

func (s baseService) CalcMin(ctx context.Context, req model.MathRequest) ([]int, error) {
	return api.CalcMin(req.List, req.Quantifier)
}

func (s baseService) CalcMax(ctx context.Context, req model.MathRequest) ([]int, error) {
	return api.CalcMax(req.List, req.Quantifier)
}
func (s baseService) CalcAvg(ctx context.Context, req model.MathRequest) (float64, error) {
	return api.CalcAvg(req.List)
}
func (s baseService) CalcMedian(ctx context.Context, req model.MathRequest) (float64, error) {
	return api.CalcMedian(req.List)
}
func (s baseService) CalcPercentile(ctx context.Context, req model.MathRequest) (int, error) {
	return api.CalcPercentile(req.List, req.Quantifier)
}
