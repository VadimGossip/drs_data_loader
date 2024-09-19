package service

import (
	"context"
)

type RateService interface {
	Refresh(ctx context.Context) error
}
