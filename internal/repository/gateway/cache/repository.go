package cache

import (
	"context"
	def "github.com/VadimGossip/drs_data_loader/internal/repository"
)

var _ def.DstGatewayRepository = (*repository)(nil)

type data struct {
	supGwgrIds []int64
}

type repository struct {
	data data
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) TruncateData() error {
	r.data = data{
		supGwgrIds: make([]int64, 0, len(r.data.supGwgrIds)),
	}
	return nil
}

func (r *repository) LoadSupGwgrIds(supGwgrIds []int64) error {
	r.data.supGwgrIds = supGwgrIds
	return nil
}

func (r *repository) GetSupGwgrIds(_ context.Context) ([]int64, error) {
	return r.data.supGwgrIds, nil
}
