package gateway

import (
	"github.com/VadimGossip/drs_data_loader/internal/repository"
	db "github.com/VadimGossip/platform_common/pkg/db/oracle"
)

type service struct {
	dstGatewayRepo repository.DstGatewayRepository
	srcGatewayRepo repository.SrcGatewayRepository
	txManager      db.TxManager
}

func NewService(dstGatewayRepo repository.DstGatewayRepository,
	srcGatewayRepo repository.SrcGatewayRepository,
	txManager db.TxManager) *service {
	return &service{
		dstGatewayRepo: dstGatewayRepo,
		srcGatewayRepo: srcGatewayRepo,
		txManager:      txManager,
	}
}
