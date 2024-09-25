package gateway

import "context"

func (s *service) GetSupGwgrIds(ctx context.Context) ([]int64, error) {
	return s.dstGatewayRepo.GetSupGwgrIds(ctx)
}
