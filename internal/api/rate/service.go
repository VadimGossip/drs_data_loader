package rate

import (
	"github.com/VadimGossip/drs_data_loader/internal/service"
	desc "github.com/VadimGossip/drs_data_loader/pkg/rate_v1"
)

type Implementation struct {
	desc.UnimplementedRateV1Server
	rateService    service.RateService
	gatewayService service.GatewayService
}

func NewImplementation(rateService service.RateService, gatewayService service.GatewayService) *Implementation {
	return &Implementation{
		rateService:    rateService,
		gatewayService: gatewayService,
	}
}
