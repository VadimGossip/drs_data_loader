package converter

import (
	"github.com/VadimGossip/drs_data_loader/internal/model"
	desc "github.com/VadimGossip/drs_data_loader/pkg/rate_v1"
)

func ToSupRatesBaseFromService(supRatesBase map[int64]model.RateBase) []*desc.SupRateBase {
	result := make([]*desc.SupRateBase, 0, len(supRatesBase))
	for gwgrId, rateBase := range supRatesBase {
		result = append(result, &desc.SupRateBase{
			GwgrId: gwgrId,
			Rate: &desc.RateBase{
				RmsrId:    rateBase.RmsrId,
				PriceBase: rateBase.PriceBase,
			},
		})
	}
	return result
}
