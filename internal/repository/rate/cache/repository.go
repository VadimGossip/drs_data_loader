package cache

import (
	"fmt"
	"github.com/VadimGossip/drs_data_loader/internal/model"
	def "github.com/VadimGossip/drs_data_loader/internal/repository"
	"github.com/VadimGossip/drs_data_loader/pkg/util"
)

var _ def.DstRatesRepository = (*repository)(nil)

type termData struct {
	aRmsgs map[model.ARmsgKey][]model.IdHistItem
	bRmsgs map[model.BRmsgKey][]model.IdHistItem
}

type origData struct {
	aRmsgs map[uint64]map[model.GwgrRmsgKey][]model.IdHistItem
	bRmsgs map[uint64]map[int64][]model.IdHistItem
}

type data struct {
	term     termData
	orig     origData
	rmsvs    map[model.RateKey][]model.RmsRateHistItem
	rates    map[int64]model.Rate
	curRates map[int64][]model.CurrencyRateHist
}

type repository struct {
	data data
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) TruncateData() error {
	r.data = data{
		term: termData{
			aRmsgs: make(map[model.ARmsgKey][]model.IdHistItem),
			bRmsgs: make(map[model.BRmsgKey][]model.IdHistItem),
		},
		orig: origData{
			aRmsgs: make(map[uint64]map[model.GwgrRmsgKey][]model.IdHistItem),
			bRmsgs: make(map[uint64]map[int64][]model.IdHistItem),
		},
		rmsvs:    make(map[model.RateKey][]model.RmsRateHistItem),
		rates:    make(map[int64]model.Rate),
		curRates: make(map[int64][]model.CurrencyRateHist),
	}
	return nil
}

func (r *repository) LoadTermAGroups(aRmsgs map[model.ARmsgKey][]model.IdHistItem) error {
	r.data.term.aRmsgs = aRmsgs
	return nil
}

func (r *repository) LoadTermBGroups(bRmsgs map[model.BRmsgKey][]model.IdHistItem) error {
	r.data.term.bRmsgs = bRmsgs
	return nil
}

func (r *repository) LoadOrigAGroups(aRmsgs map[uint64]map[model.GwgrRmsgKey][]model.IdHistItem) error {
	r.data.orig.aRmsgs = aRmsgs
	return nil
}

func (r *repository) LoadOrigBGroups(bRmsgs map[uint64]map[int64][]model.IdHistItem) error {
	r.data.orig.bRmsgs = bRmsgs
	return nil
}

func (r *repository) LoadRates(rmsvs map[model.RateKey][]model.RmsRateHistItem) error {
	for key, val := range rmsvs {
		r.data.rmsvs[key] = val
	}
	return nil
}

func (r *repository) LoadRateValues(rates map[int64]model.Rate) error {
	for key, val := range rates {
		r.data.rates[key] = val
	}
	return nil
}

func (r *repository) LoadCurrencyRates(curRates map[int64][]model.CurrencyRateHist) error {
	r.data.curRates = curRates
	return nil
}

func (r *repository) getBRmsg(key model.BRmsgKey, dateAt int64) (int64, error) {
	for key.Code > 0 {
		if h, ok := r.data.term.bRmsgs[key]; ok {
			for _, item := range h {
				if item.DBegin <= dateAt && item.DEnd > dateAt {
					return item.Id, nil
				}
			}
		}
		key.Code /= 10
	}
	return 0, fmt.Errorf("can't find B-code rate group")
}

func (r *repository) getARmsg(key model.ARmsgKey, dateAt int64) int64 {
	for key.Code > 0 {
		if h, ok := r.data.term.aRmsgs[key]; ok {
			for _, item := range h {
				if item.DBegin <= dateAt && item.DEnd > dateAt {
					return item.Id
				}
			}
		}
		key.Code /= 10
	}
	return -2
}

func (r *repository) buildOrigARmsgsKeys(aNumber, bNumber uint64, dateAt int64) map[model.ARmsgShortKey]struct{} {
	result := make(map[model.ARmsgShortKey]struct{})
	addedSup := make(map[int64]struct{})
	for bNumber > 0 {
		if hList, ok := r.data.orig.bRmsgs[bNumber]; ok {
			for gwgrId, h := range hList {
				if _, rmsgFound := addedSup[gwgrId]; rmsgFound {
					continue
				}
				for _, item := range h {
					if item.DBegin <= dateAt && item.DEnd > dateAt {
						addedSup[gwgrId] = struct{}{}
						result[model.ARmsgShortKey{
							GwgrId:  gwgrId,
							BRmsgId: item.Id,
							Code:    aNumber,
						}] = struct{}{}
					}
				}
			}
		}
		bNumber /= 10
	}
	return result
}

func (r *repository) buildOrigRateKeys(aNumber uint64, keys map[model.ARmsgShortKey]struct{}, dateAt int64) map[model.RateKey]struct{} {
	result := make(map[model.RateKey]struct{}, len(keys))
	for key := range keys {
		result[model.RateKey{
			GwgrId:    key.GwgrId,
			Direction: 1,
			ARmsgId:   -2,
			BRmsgId:   key.BRmsgId,
		}] = struct{}{}
	}

	for aNumber > 0 {
		if hList, ok := r.data.orig.aRmsgs[aNumber]; ok {
			for gwgrRmsgKey, h := range hList {
				if _, rmsgNotFound := result[model.RateKey{
					GwgrId:    gwgrRmsgKey.GwgrId,
					Direction: 1,
					ARmsgId:   -2,
					BRmsgId:   gwgrRmsgKey.RmsgId,
				}]; !rmsgNotFound {
					continue
				}

				for _, item := range h {
					if item.DBegin <= dateAt && item.DEnd > dateAt {
						delete(result, model.RateKey{
							GwgrId:    gwgrRmsgKey.GwgrId,
							Direction: 1,
							ARmsgId:   -2,
							BRmsgId:   gwgrRmsgKey.RmsgId,
						})

						result[model.RateKey{
							GwgrId:    gwgrRmsgKey.GwgrId,
							Direction: 1,
							ARmsgId:   item.Id,
							BRmsgId:   gwgrRmsgKey.RmsgId,
						}] = struct{}{}
					}
				}
			}
		}
		aNumber /= 10
	}
	return result
}

func (r *repository) getRmsrRmsvPair(key model.RateKey, dateAt int64) (int64, int64, error) {
	if h, ok := r.data.rmsvs[key]; ok {
		for _, item := range h {
			if item.DBegin <= dateAt && item.DEnd > dateAt {
				return item.RmsrId, item.RmsvId, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("can't find rmsv_id")
}

func (r *repository) getRateValue(rmsvId int64) (model.Rate, error) {
	if rv, ok := r.data.rates[rmsvId]; ok {
		return rv, nil
	}
	return model.Rate{}, fmt.Errorf("can't find rate value")
}

func (r *repository) getCurrencyRate(currencyId int64, dateAt int64) (float64, error) {
	if hist, ok := r.data.curRates[currencyId]; ok {
		for _, item := range hist {
			if dateAt >= item.DBegin && dateAt < item.DEnd {
				return item.CurrencyRate, nil
			}
		}
	}
	return 0, fmt.Errorf("can't find currency rate")
}

func (r *repository) findRateByRateKey(key model.RateKey, dateAt int64) (model.RateBase, error) {
	rmsrId, rmsvId, err := r.getRmsrRmsvPair(key, dateAt)
	if err != nil {
		return model.RateBase{}, err
	}

	rv, err := r.getRateValue(rmsvId)
	if err != nil {
		return model.RateBase{}, err
	}

	currencyRate, err := r.getCurrencyRate(rv.CurrencyId, dateAt)
	if err != nil {
		return model.RateBase{}, err
	}

	return model.RateBase{
		RmsrId:    rmsrId,
		PriceBase: util.RoundFloat(rv.Price*currencyRate, 7),
	}, nil
}

func (r *repository) FindRate(gwgrId, dateAt int64, dir uint8, aNumber, bNumber uint64) (model.RateBase, error) {
	bRmsgId, err := r.getBRmsg(model.BRmsgKey{
		GwgrId:    gwgrId,
		Direction: dir,
		Code:      bNumber,
	}, dateAt)
	if err != nil {
		return model.RateBase{}, err
	}

	aRmsgId := r.getARmsg(model.ARmsgKey{
		GwgrId:    gwgrId,
		Direction: dir,
		BRmsgId:   bRmsgId,
		Code:      aNumber,
	}, dateAt)

	return r.findRateByRateKey(model.RateKey{
		GwgrId:    gwgrId,
		Direction: dir,
		ARmsgId:   aRmsgId,
		BRmsgId:   bRmsgId,
	}, dateAt)
}

func (r *repository) FindSupRates(dateAt int64, aNumber, bNumber uint64) (map[int64]model.RateBase, error) {
	result := make(map[int64]model.RateBase)
	aRmsgShortKeys := r.buildOrigARmsgsKeys(aNumber, bNumber, dateAt)
	if len(aRmsgShortKeys) == 0 {
		return result, nil
	}

	rateKeys := r.buildOrigRateKeys(aNumber, aRmsgShortKeys, dateAt)
	if len(rateKeys) == 0 {
		return result, nil
	}

	for key := range rateKeys {
		rate, err := r.findRateByRateKey(key, dateAt)
		if err != nil {
			continue
		}
		result[key.GwgrId] = rate
	}
	return result, nil
}

func (r *repository) FindSupRatesOld(gwgrIds []int64, dateAt int64, aNumber, bNumber uint64) (map[int64]model.RateBase, error) {
	result := make(map[int64]model.RateBase)
	for _, gwgrId := range gwgrIds {
		rateBase, err := r.FindRate(gwgrId, dateAt, 1, aNumber, bNumber)
		if err != nil {
			continue
		}
		result[gwgrId] = rateBase
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no sup rates found")
	}
	return result, nil
}
