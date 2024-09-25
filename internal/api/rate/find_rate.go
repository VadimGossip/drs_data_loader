package rate

import (
	"context"

	desc "github.com/VadimGossip/drs_data_loader/pkg/rate_v1"
)

func (i *Implementation) FindRate(ctx context.Context, req *desc.FindRateRequest) (*desc.FindRateResponse, error) {
	rateBase, err := i.rateService.FindRate(ctx, req.GwgrId, req.DateAt, uint8(req.Dir), req.ANumber, req.BNumber)
	if err != nil {
		return nil, err
	}

	return &desc.FindRateResponse{Rate: &desc.RateBase{
		RmsrId:    rateBase.RmsrId,
		PriceBase: rateBase.PriceBase,
	}}, nil
}
