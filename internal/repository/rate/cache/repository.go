package cache

import (
	"fmt"
	"github.com/VadimGossip/drs_data_loader/internal/model"
	def "github.com/VadimGossip/drs_data_loader/internal/repository"
	"github.com/VadimGossip/drs_data_loader/pkg/util"
)

var _ def.DstRatesRepository = (*repository)(nil)

type data struct {
	aRmsgs   map[model.ARmsgKey][]model.IdHistItem
	bRmsgs   map[model.BRmsgKey][]model.IdHistItem
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
		aRmsgs:   make(map[model.ARmsgKey][]model.IdHistItem),
		bRmsgs:   make(map[model.BRmsgKey][]model.IdHistItem),
		rmsvs:    make(map[model.RateKey][]model.RmsRateHistItem),
		rates:    make(map[int64]model.Rate),
		curRates: make(map[int64][]model.CurrencyRateHist),
	}
	return nil
}

func (r *repository) LoadBGroups(bRmsgs map[model.BRmsgKey][]model.IdHistItem) error {
	for key, val := range bRmsgs {
		r.data.bRmsgs[key] = val
	}
	return nil
}

func (r *repository) LoadAGroups(aRmsgs map[model.ARmsgKey][]model.IdHistItem) error {
	for key, val := range aRmsgs {
		r.data.aRmsgs[key] = val
	}
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
	for i := len(key.Code); i > 0; i-- {
		if h, ok := r.data.bRmsgs[key]; ok {
			for _, item := range h {
				if item.DBegin <= dateAt && item.DEnd > dateAt {
					return item.Id, nil
				}
			}
		}
		key.Code = key.Code[:i-1]
	}
	return 0, fmt.Errorf("can't find B-code rate group")
}

func (r *repository) getARmsg(key model.ARmsgKey, dateAt int64) int64 {
	for i := len(key.Code); i > 0; i-- {
		if h, ok := r.data.aRmsgs[key]; ok {
			for _, item := range h {
				if item.DBegin <= dateAt && item.DEnd > dateAt {
					return item.Id
				}
			}
		}
		key.Code = key.Code[:i-1]
	}
	return -2
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

func (r *repository) FindRate(gwgrId, dateAt int64, dir uint8, aNumber, bNumber string) (model.RateBase, error) {
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

	rmsrId, rmsvId, err := r.getRmsrRmsvPair(model.RateKey{
		GwgrId:    gwgrId,
		Direction: dir,
		ARmsgId:   aRmsgId,
		BRmsgId:   bRmsgId,
	}, dateAt)
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

func (r *repository) FindSupRates(gwgrIds []int64, dateAt int64, aNumber, bNumber string) (map[int64]model.RateBase, error) {
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
