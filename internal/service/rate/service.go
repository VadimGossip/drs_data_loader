package rate

import (
	"drs_data_loader/internal/repository"
	def "drs_data_loader/internal/service"

	db "github.com/VadimGossip/platform_common/pkg/db/oracle"
)

var _ def.RateService = (*service)(nil)

type service struct {
	dstRateRepo repository.DstRatesRepository
	srcRateRepo repository.SrcRatesRepository
	txManager   db.TxManager
}

func NewService(dstRateRepo repository.DstRatesRepository,
	srcRateRepo repository.SrcRatesRepository,
	txManager db.TxManager) *service {
	return &service{
		dstRateRepo: dstRateRepo,
		srcRateRepo: srcRateRepo,
		txManager:   txManager,
	}
}
