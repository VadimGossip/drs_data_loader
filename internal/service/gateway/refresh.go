package gateway

import (
	"context"
	"fmt"
	"github.com/VadimGossip/drs_data_loader/internal/model"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *service) refreshSupGwgrIds(ctx context.Context) error {
	ts := time.Now()
	supGwgrIds, err := s.srcGatewayRepo.GetSupGwgrIds(ctx)
	if err != nil {
		return fmt.Errorf("error while get supGwgrIds %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", model.SUPGWObjectKey, len(supGwgrIds), time.Since(ts))

	ts = time.Now()
	if err = s.dstGatewayRepo.LoadSupGwgrIds(supGwgrIds); err != nil {
		return fmt.Errorf("error while load supGwgrIds %s", err)
	}
	logrus.Infof("Load %s. Duration: %s", model.SUPGWObjectKey, time.Since(ts))

	return nil
}

func (s *service) Refresh(ctx context.Context) error {
	refreshList := []func(ctx context.Context) error{
		s.refreshSupGwgrIds,
	}
	//make in TX

	if err := s.dstGatewayRepo.TruncateData(); err != nil {
		return fmt.Errorf("error while truncate data: %s", err)
	}

	for _, f := range refreshList {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}
