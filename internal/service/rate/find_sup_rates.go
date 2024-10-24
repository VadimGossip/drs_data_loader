package rate

import (
	"context"

	"github.com/VadimGossip/drs_data_loader/internal/model"
)

func (s *service) FindSupRates(_ context.Context, dateAt int64, aNumber, bNumber uint64) (map[int64]model.RateBase, error) {
	return s.dstRateRepo.FindSupRates(dateAt, aNumber, bNumber)
}
