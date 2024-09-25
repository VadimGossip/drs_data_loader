package service

import (
	"context"
	"github.com/VadimGossip/drs_data_loader/internal/model"
)

type RateService interface {
	Refresh(ctx context.Context) error
	FindRate(ctx context.Context, gwgrId, dateAt int64, dir uint8, aNumber, bNumber string) (model.RateBase, error)
	FindSupRates(_ context.Context, gwgrIds []int64, dateAt int64, aNumber, bNumber string) (map[int64]model.RateBase, error)
}

type GatewayService interface {
	GetSupGwgrIds(ctx context.Context) ([]int64, error)
}
