package tarantool

import (
	db "github.com/VadimGossip/drs_data_loader/internal/client/db/tarantool"
	"github.com/VadimGossip/drs_data_loader/internal/model"
	"github.com/VadimGossip/drs_data_loader/internal/repository/rate/tarantool/converter"

	def "github.com/VadimGossip/drs_data_loader/internal/repository"

	"github.com/tarantool/go-tarantool/v2"
)

var _ def.DstRatesRepository = (*repository)(nil)

const (
	insertBRmsgGroupsFunc   string = "rates.insert_b_rmsg_groups"
	insertARmsgGroupsFunc   string = "rates.insert_a_rmsg_groups"
	insertRatesFunc         string = "rates.insert_rates"
	insertRateValuesFunc    string = "rates.insert_rate_values"
	insertCurrencyRatesFunc string = "rates.insert_currency_rates"
	truncateData            string = "rates.truncate_data"
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

func (r *repository) LoadBGroups(data map[model.BRmsgKey][]model.IdHistItem) error {
	_, err := r.db.DB().Do(tarantool.NewCallRequest(insertBRmsgGroupsFunc).Args([]interface{}{converter.ToRepoFromBRmsgGroups(data)})).Get()
	return err
}

func (r *repository) LoadAGroups(data map[model.ARmsgKey][]model.IdHistItem) error {
	_, err := r.db.DB().Do(tarantool.NewCallRequest(insertARmsgGroupsFunc).Args([]interface{}{converter.ToRepoFromARmsgGroups(data)})).Get()
	return err
}

func (r *repository) LoadRates(data map[model.RateKey][]model.RmsRateHistItem) error {
	_, err := r.db.DB().Do(tarantool.NewCallRequest(insertRatesFunc).Args([]interface{}{converter.ToRepoFromRmsRates(data)})).Get()
	return err
}

func (r *repository) LoadRateValues(data map[int64]model.Rate) error {
	_, err := r.db.DB().Do(tarantool.NewCallRequest(insertRateValuesFunc).Args([]interface{}{converter.ToRepoFromRateValues(data)})).Get()
	return err
}

func (r *repository) LoadCurrencyRates(data map[int64][]model.CurrencyRateHist) error {
	_, err := r.db.DB().Do(tarantool.NewCallRequest(insertCurrencyRatesFunc).Args([]interface{}{converter.ToRepoFromCurrencyRates(data)})).Get()
	return err
}
