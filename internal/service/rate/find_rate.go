package rate

import "context"

func (s *service) FindRate(_ context.Context, gwgrId, dateAt int64, dir uint8, aNumber, bNumber string) (int64, float64, error) {
	return s.dstRateRepo.FindRate(gwgrId, dateAt, dir, aNumber, bNumber)
}
