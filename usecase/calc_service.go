package usecase

import (
	"context"
	"go-onion-sample/domain/entity"
)

type ICalcService interface {
	Add(ctx context.Context, params *entity.Params) (int, error)
	Subtract(ctx context.Context, params *entity.Params) (int, error)
}

type CalcService struct{}

func NewCalcService() *CalcService {
	return &CalcService{}
}

func (self *CalcService) Add(ctx context.Context, params *entity.Params) (int, error) {
	return params.A + params.B, nil
}

func (self *CalcService) Subtract(ctx context.Context, params *entity.Params) (int, error) {
	return params.A - params.B, nil
}
