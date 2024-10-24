package rate

import (
	"context"

	"github.com/VadimGossip/drs_data_loader/internal/converter"
	desc "github.com/VadimGossip/drs_data_loader/pkg/rate_v1"
)

func (i *Implementation) FindSupRates(ctx context.Context, req *desc.FindSupRatesRequest) (*desc.FindSupRatesResponse, error) {
	supRatesBase, err := i.rateService.FindSupRates(ctx, req.DateAt, req.ANumber, req.BNumber)
	if err != nil {
		return nil, err
	}

	return &desc.FindSupRatesResponse{SupRatesBase: converter.ToSupRatesBaseFromService(supRatesBase)}, nil
}
