package tarantool

import (
	"fmt"

	db "github.com/VadimGossip/platform_common/pkg/db/tarantool"
	"github.com/tarantool/go-tarantool/v2"

	"github.com/VadimGossip/drs_data_loader/internal/model"
	"github.com/VadimGossip/drs_data_loader/internal/repository/rate/tarantool/converter"
)

// Need refactor like cache

const (
	insertBRmsgGroupsFunc   string = "rates.insert_b_rmsg_groups"
	insertARmsgGroupsFunc   string = "rates.insert_a_rmsg_groups"
	insertRatesFunc         string = "rates.insert_rates"
	insertRateValuesFunc    string = "rates.insert_rate_values"
	insertCurrencyRatesFunc string = "rates.insert_currency_rates"
	truncateData            string = "rates.truncate_data"
	chunkSize               int    = 50_000
)

type repository struct {
	db db.Client
}

func NewRepository(db db.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) TruncateData() error {
	_, err := r.db.DB().Do(tarantool.NewCallRequest(truncateData)).Get()
	return err
}

func (r *repository) LoadAGroups(aRmsgs map[model.ARmsgKey][]model.IdHistItem) error {
	chunk := make(map[model.ARmsgKey][]model.IdHistItem)
	counter := 0
	for key, value := range aRmsgs {
		counter++
		chunk[key] = value
		if len(chunk) == chunkSize || counter == len(aRmsgs) {
			_, err := r.db.DB().Do(tarantool.NewCallRequest(insertARmsgGroupsFunc).Args([]interface{}{converter.ToRepoFromARmsgGroups(chunk)})).Get()
			if err != nil {
				return err
			}
			chunk = make(map[model.ARmsgKey][]model.IdHistItem)
		}
	}
	return nil
}

func (r *repository) LoadBGroups(bRmsgs map[model.BRmsgKey][]model.IdHistItem) error {
	chunk := make(map[model.BRmsgKey][]model.IdHistItem)
	counter := 0
	for key, value := range bRmsgs {
		counter++
		chunk[key] = value
		if len(chunk) == chunkSize || counter == len(bRmsgs) {
			_, err := r.db.DB().Do(tarantool.NewCallRequest(insertBRmsgGroupsFunc).Args([]interface{}{converter.ToRepoFromBRmsgGroups(chunk)})).Get()
			if err != nil {
				return err
			}
			chunk = make(map[model.BRmsgKey][]model.IdHistItem)
		}
	}
	return nil
}

func (r *repository) LoadRates(rates map[model.RateKey][]model.RmsRateHistItem) error {
	chunk := make(map[model.RateKey][]model.RmsRateHistItem)
	counter := 0
	for key, value := range rates {
		counter++
		chunk[key] = value
		if len(chunk) == chunkSize || counter == len(rates) {
			_, err := r.db.DB().Do(tarantool.NewCallRequest(insertRatesFunc).Args([]interface{}{converter.ToRepoFromRmsRates(chunk)})).Get()
			if err != nil {
				return err
			}
			chunk = make(map[model.RateKey][]model.RmsRateHistItem)
		}
	}
	return nil
}

func (r *repository) LoadRateValues(rv map[int64]model.Rate) error {
	chunk := make(map[int64]model.Rate)
	counter := 0
	for key, value := range rv {
		counter++
		chunk[key] = value
		if len(chunk) == chunkSize || counter == len(rv) {
			_, err := r.db.DB().Do(tarantool.NewCallRequest(insertRateValuesFunc).Args([]interface{}{converter.ToRepoFromRateValues(chunk)})).Get()
			if err != nil {
				return err
			}
			chunk = make(map[int64]model.Rate)
		}
	}
	return nil
}

func (r *repository) LoadCurrencyRates(curRates map[int64][]model.CurrencyRateHist) error {
	chunk := make(map[int64][]model.CurrencyRateHist)
	counter := 0
	for key, value := range curRates {
		counter++
		chunk[key] = value
		if len(chunk) == chunkSize || counter == len(curRates) {
			_, err := r.db.DB().Do(tarantool.NewCallRequest(insertCurrencyRatesFunc).Args([]interface{}{converter.ToRepoFromCurrencyRates(chunk)})).Get()
			if err != nil {
				return err
			}
			chunk = make(map[int64][]model.CurrencyRateHist)
		}
	}
	return nil
}

func (r *repository) FindRate(gwgrId, dateAt int64, dir uint8, aNumber, bNumber string) (model.RateBase, error) {
	return model.RateBase{}, fmt.Errorf("not implemented")
}

func (r *repository) FindSupRates(gwgrIds []int64, dateAt int64, aNumber, bNumber string) (map[int64]model.RateBase, error) {
	return nil, fmt.Errorf("not implemented")
}
