package rate

import (
	"github.com/VadimGossip/drs_data_loader/internal/service"
	desc "github.com/VadimGossip/drs_data_loader/pkg/rate_v1"
)

type Implementation struct {
	desc.UnimplementedRateV1Server
	rateService service.RateService
}

func NewImplementation(rateService service.RateService) *Implementation {
	return &Implementation{
		rateService: rateService,
	}
}
