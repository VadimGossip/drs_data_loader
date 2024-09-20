package service

import (
	"context"
)

type RateService interface {
	Refresh(ctx context.Context) error
	FindRate(ctx context.Context, gwgrId, dateAt int64, dir uint8, aNumber, bNumber string) (int64, float64, error)
}
