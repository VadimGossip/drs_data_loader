package service

import (
	"context"
	"github.com/VadimGossip/drs_data_loader/internal/model"
)

type RateService interface {
	Refresh(ctx context.Context) error
	FindRate(ctx context.Context, gwgrId, dateAt int64, dir uint8, aNumber, bNumber uint64) (model.RateBase, error)
	FindSupRates(_ context.Context, dateAt int64, aNumber, bNumber uint64) (map[int64]model.RateBase, error)
}

type GatewayService interface {
	Refresh(ctx context.Context) error
	GetSupGwgrIds(ctx context.Context) ([]int64, error)
}
