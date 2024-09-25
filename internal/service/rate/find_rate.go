package rate

import (
	"context"
	"github.com/VadimGossip/drs_data_loader/internal/model"
)

func (s *service) FindRate(_ context.Context, gwgrId, dateAt int64, dir uint8, aNumber, bNumber string) (model.RateBase, error) {
	return s.dstRateRepo.FindRate(gwgrId, dateAt, dir, aNumber, bNumber)
}
