package converter

import "drs_data_loader/internal/model"

func ToRepoFromBRmsgGroups(groups map[model.BRmsgKey][]model.IdHistItem) [][]interface{} {
	result := make([][]interface{}, 0, len(groups))
	for key, value := range groups {
		ids := make([]interface{}, 0, len(value))
		dBegins := make([]interface{}, 0, len(value))
		dEnds := make([]interface{}, 0, len(value))
		for _, v := range value {
			ids = append(ids, v.Id)
			dBegins = append(dBegins, v.DBegin)
			dEnds = append(dEnds, v.DEnd)
		}
		result = append(result, []interface{}{
			key.GwgrId,
			key.Direction,
			key.Code,
			ids,
			dBegins,
			dEnds,
		})
	}
	return result
}

func ToRepoFromARmsgGroups(groups map[model.ARmsgKey][]model.IdHistItem) [][]interface{} {
	result := make([][]interface{}, 0, len(groups))
	for key, value := range groups {
		ids := make([]interface{}, 0, len(value))
		dBegins := make([]interface{}, 0, len(value))
		dEnds := make([]interface{}, 0, len(value))
		for _, v := range value {
			ids = append(ids, v.Id)
			dBegins = append(dBegins, v.DBegin)
			dEnds = append(dEnds, v.DEnd)
		}
		result = append(result, []interface{}{
			key.GwgrId,
			key.Direction,
			key.BRmsgId,
			key.Code,
			ids,
			dBegins,
			dEnds,
		})
	}
	return result
}

func ToRepoFromRmsRates(groups map[model.RateKey][]model.RmsRateHistItem) [][]interface{} {
	result := make([][]interface{}, 0, len(groups))
	for key, value := range groups {
		rmsrIds := make([]interface{}, 0, len(value))
		rmsvIds := make([]interface{}, 0, len(value))
		dBegins := make([]interface{}, 0, len(value))
		dEnds := make([]interface{}, 0, len(value))
		for _, v := range value {
			rmsrIds = append(rmsrIds, v.RmsrId)
			rmsvIds = append(rmsvIds, v.RmsvId)
			dBegins = append(dBegins, v.DBegin)
			dEnds = append(dEnds, v.DEnd)
		}
		result = append(result, []interface{}{
			key.GwgrId,
			key.Direction,
			key.ARmsgId,
			key.BRmsgId,
			rmsrIds,
			rmsvIds,
			dBegins,
			dEnds,
		})
	}
	return result
}

func ToRepoFromRateValues(groups map[int64]model.Rate) [][]interface{} {
	result := make([][]interface{}, 0, len(groups))
	for key, value := range groups {
		result = append(result, []interface{}{
			key,
			value.Price,
			value.CurrencyId,
		})
	}
	return result
}

func ToRepoFromCurrencyRates(groups map[int64][]model.CurrencyRateHist) [][]interface{} {
	result := make([][]interface{}, 0, len(groups))
	for key, value := range groups {
		curRates := make([]interface{}, 0, len(value))
		dBegins := make([]interface{}, 0, len(value))
		dEnds := make([]interface{}, 0, len(value))
		for _, v := range value {
			curRates = append(curRates, v.CurrencyRate)
			dBegins = append(dBegins, v.DBegin)
			dEnds = append(dEnds, v.DEnd)
		}
		result = append(result, []interface{}{
			key,
			curRates,
			dBegins,
			dEnds,
		})
	}
	return result
}
