package rate

import (
	"context"

	"github.com/VadimGossip/drs_data_loader/internal/model"
)

func (s *service) FindSupRates(_ context.Context, gwgrIds []int64, dateAt int64, aNumber, bNumber string) (map[int64]model.RateBase, error) {
	return s.dstRateRepo.FindSupRates(gwgrIds, dateAt, aNumber, bNumber)
}
