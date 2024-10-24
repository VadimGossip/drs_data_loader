package repository

import (
	"context"
	"github.com/VadimGossip/drs_data_loader/internal/model"
)

type DstRatesRepository interface {
	TruncateData() error
	LoadTermAGroups(aRmsgs map[model.ARmsgKey][]model.IdHistItem) error
	LoadTermBGroups(bRmsgs map[model.BRmsgKey][]model.IdHistItem) error
	LoadOrigAGroups(aRmsgs map[uint64]map[model.GwgrRmsgKey][]model.IdHistItem) error
	LoadOrigBGroups(bRmsgs map[uint64]map[int64][]model.IdHistItem) error
	LoadRates(data map[model.RateKey][]model.RmsRateHistItem) error
	LoadRateValues(data map[int64]model.Rate) error
	LoadCurrencyRates(data map[int64][]model.CurrencyRateHist) error
	FindRate(gwgrId, dateAt int64, dir uint8, aNumber, bNumber uint64) (model.RateBase, error)
	FindSupRates(dateAt int64, aNumber, bNumber uint64) (map[int64]model.RateBase, error)
}

type SrcRatesRepository interface {
	GetTermAGroups(ctx context.Context) (map[model.ARmsgKey][]model.IdHistItem, int, error)
	GetTermBGroups(ctx context.Context) (map[model.BRmsgKey][]model.IdHistItem, int, error)
	GetOrigAGroups(ctx context.Context) (map[uint64]map[model.GwgrRmsgKey][]model.IdHistItem, int, error)
	GetOrigBGroups(ctx context.Context) (map[uint64]map[int64][]model.IdHistItem, int, error)
	GetRates(ctx context.Context) (map[model.RateKey][]model.RmsRateHistItem, int, error)
	GetRateValues(ctx context.Context) (map[int64]model.Rate, int, error)
	GetCurrencyRates(ctx context.Context) (map[int64][]model.CurrencyRateHist, int, error)
}

type SrcGatewayRepository interface {
	GetSupGwgrIds(ctx context.Context) ([]int64, error)
}

type DstGatewayRepository interface {
	TruncateData() error
	LoadSupGwgrIds(data []int64) error
	GetSupGwgrIds(ctx context.Context) ([]int64, error)
}
