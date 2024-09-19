package rate

import (
	"context"
	"fmt"
	"github.com/VadimGossip/drs_data_loader/internal/model"
	"time"

	"github.com/sirupsen/logrus"
)

var chunkSize int = 100_000

func (s *service) refreshBRmsgs(ctx context.Context) error {
	ts := time.Now()
	bRmsgs, rows, err := s.srcRateRepo.GetBGroups(ctx)
	if err != nil {
		return fmt.Errorf("error while get bRmsgs %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", "RB", rows, time.Since(ts))

	ts = time.Now()
	chunk := make(map[model.BRmsgKey][]model.IdHistItem)
	counter := 0
	for key, value := range bRmsgs {
		counter++
		chunk[key] = value
		if len(chunk) == chunkSize || counter == len(bRmsgs) {
			err = s.dstRateRepo.LoadBGroups(chunk)
			if err != nil {
				return fmt.Errorf("error while load bRmsgs %s", err)
			}
			chunk = make(map[model.BRmsgKey][]model.IdHistItem)
		}
	}

	logrus.Infof("Load %s. Duration: %s", "RB", time.Since(ts))
	return nil
}

func (s *service) refreshARmsgs(ctx context.Context) error {
	ts := time.Now()
	aRmsgs, rows, err := s.srcRateRepo.GetAGroups(ctx)
	if err != nil {
		return fmt.Errorf("error while get aRmsgs %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", "RA", rows, time.Since(ts))

	ts = time.Now()
	chunk := make(map[model.ARmsgKey][]model.IdHistItem)
	counter := 0
	for key, value := range aRmsgs {
		counter++
		chunk[key] = value
		if len(chunk) == chunkSize || counter == len(aRmsgs) {
			err = s.dstRateRepo.LoadAGroups(chunk)
			if err != nil {
				return fmt.Errorf("error while load aRmsgs %s", err)
			}
			chunk = make(map[model.ARmsgKey][]model.IdHistItem)
		}
	}

	logrus.Infof("Load %s. Duration: %s", "RA", time.Since(ts))
	return nil
}

func (s *service) refreshRates(ctx context.Context) error {
	ts := time.Now()
	rates, rows, err := s.srcRateRepo.GetRates(ctx)
	if err != nil {
		return fmt.Errorf("error while get rates %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", "RTS", rows, time.Since(ts))

	ts = time.Now()
	chunk := make(map[model.RateKey][]model.RmsRateHistItem)
	counter := 0
	for key, value := range rates {
		counter++
		chunk[key] = value
		if len(chunk) == chunkSize || counter == len(rates) {
			err = s.dstRateRepo.LoadRates(chunk)
			if err != nil {
				return fmt.Errorf("error while load rates %s", err)
			}
			chunk = make(map[model.RateKey][]model.RmsRateHistItem)
		}
	}

	logrus.Infof("Load %s. Duration: %s", "RTS", time.Since(ts))
	return nil
}

func (s *service) refreshRateValues(ctx context.Context) error {
	ts := time.Now()
	rv, rows, err := s.srcRateRepo.GetRateValues(ctx)
	if err != nil {
		return fmt.Errorf("error while get rate values %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", "RV", rows, time.Since(ts))

	ts = time.Now()
	chunk := make(map[int64]model.Rate)
	counter := 0
	for key, value := range rv {
		counter++
		chunk[key] = value
		if len(chunk) == chunkSize || counter == len(rv) {
			err = s.dstRateRepo.LoadRateValues(chunk)
			if err != nil {
				return fmt.Errorf("error while load rate values %s", err)
			}
			chunk = make(map[int64]model.Rate)
		}
	}

	logrus.Infof("Load %s. Duration: %s", "RV", time.Since(ts))
	return nil
}

func (s *service) refreshCurRates(ctx context.Context) error {
	ts := time.Now()
	cr, rows, err := s.srcRateRepo.GetCurrencyRates(ctx)
	if err != nil {
		return fmt.Errorf("error while get currency rates %s", err)
	}
	logrus.Infof("Get %s. Rows read: %d Duration: %s", "CURRRTS", rows, time.Since(ts))

	ts = time.Now()
	if err = s.dstRateRepo.LoadCurrencyRates(cr); err != nil {
		return fmt.Errorf("error while load currency rates %s", err)
	}
	logrus.Infof("Load %s. Duration: %s", "CURRRTS", time.Since(ts))
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
