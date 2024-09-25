package rate

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/VadimGossip/drs_data_loader/internal/model"
)

func (s *service) refreshARmsgs(ctx context.Context) error {
	ts := time.Now()
	aRmsgs, rows, err := s.srcRateRepo.GetAGroups(ctx)
	if err != nil {
		return fmt.Errorf("error while get aRmsgs %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", model.RAObjectKey, rows, time.Since(ts))

	ts = time.Now()
	if err = s.dstRateRepo.LoadAGroups(aRmsgs); err != nil {
		return fmt.Errorf("error while load aRmsgs %s", err)
	}
	logrus.Infof("Load %s. Duration: %s", model.RAObjectKey, time.Since(ts))

	return nil
}

func (s *service) refreshBRmsgs(ctx context.Context) error {
	ts := time.Now()
	bRmsgs, rows, err := s.srcRateRepo.GetBGroups(ctx)
	if err != nil {
		return fmt.Errorf("error while get bRmsgs %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", model.RBObjectKey, rows, time.Since(ts))

	ts = time.Now()
	if err = s.dstRateRepo.LoadBGroups(bRmsgs); err != nil {
		return fmt.Errorf("error while load bRmsgs %s", err)
	}
	logrus.Infof("Load %s. Duration: %s", model.RBObjectKey, time.Since(ts))
	return nil
}

func (s *service) refreshRates(ctx context.Context) error {
	ts := time.Now()
	rates, rows, err := s.srcRateRepo.GetRates(ctx)
	if err != nil {
		return fmt.Errorf("error while get rates %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", model.RTSObjectKey, rows, time.Since(ts))

	ts = time.Now()
	if err = s.dstRateRepo.LoadRates(rates); err != nil {
		return fmt.Errorf("error while load rates %s", err)
	}
	logrus.Infof("Load %s. Duration: %s", model.RTSObjectKey, time.Since(ts))

	return nil
}

func (s *service) refreshRateValues(ctx context.Context) error {
	ts := time.Now()
	rv, rows, err := s.srcRateRepo.GetRateValues(ctx)
	if err != nil {
		return fmt.Errorf("error while get rate values %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", model.RVObjectKey, rows, time.Since(ts))

	ts = time.Now()
	if err = s.dstRateRepo.LoadRateValues(rv); err != nil {
		return fmt.Errorf("error while load rate values %s", err)
	}
	logrus.Infof("Load %s. Duration: %s", model.RVObjectKey, time.Since(ts))

	return nil
}

func (s *service) refreshCurRates(ctx context.Context) error {
	ts := time.Now()
	cr, rows, err := s.srcRateRepo.GetCurrencyRates(ctx)
	if err != nil {
		return fmt.Errorf("error while get currency rates %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", model.CURRTSObjectKey, rows, time.Since(ts))

	ts = time.Now()
	if err = s.dstRateRepo.LoadCurrencyRates(cr); err != nil {
		return fmt.Errorf("error while load currency rates %s", err)
	}
	logrus.Infof("Load %s. Duration: %s", model.CURRTSObjectKey, time.Since(ts))

	return nil
}

func (s *service) Refresh(ctx context.Context) error {
	refreshList := []func(ctx context.Context) error{
		s.refreshARmsgs,
		s.refreshBRmsgs,
		s.refreshRates,
		s.refreshRateValues,
		s.refreshCurRates,
	}
	//make in TX

	if err := s.dstRateRepo.TruncateData(); err != nil {
		return fmt.Errorf("error while truncate data: %s", err)
	}

	for _, f := range refreshList {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}
