package oracle

import (
	"context"
	"database/sql"
	"drs_data_loader/internal/model"
	def "drs_data_loader/internal/repository"
	"fmt"
	"time"

	db "github.com/VadimGossip/platform_common/pkg/db/oracle"
	"github.com/godror/godror"
	"github.com/sirupsen/logrus"
)

type repository struct {
	db db.Client
}

var _ def.SrcRatesRepository = (*repository)(nil)

var defaultFetchSize int = 1_000_000

func NewRepository(db db.Client) *repository {
	return &repository{db: db}
}

func (r *repository) GetBGroups(ctx context.Context) (map[model.BRmsgKey][]model.IdHistItem, int, error) {
	var err error
	var rows *sql.Rows
	var actualRows, expectRows int
	if err = r.db.DB().QueryRowContext(ctx, sqlRBCountQuery).Scan(&expectRows); err != nil {
		return nil, 0, err
	}

	if expectRows == 0 {
		return nil, 0, nil
	}

	preFetchSize, fetchSize := defaultFetchSize, defaultFetchSize
	if expectRows <= defaultFetchSize {
		preFetchSize = expectRows + 1
		fetchSize = expectRows
	}
	rows, err = r.db.DB().QueryContext(ctx, sqlRBQuery, godror.PrefetchCount(preFetchSize), godror.FetchArraySize(fetchSize))
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"handler": "GetBGroups",
				"problem": "rows close",
			}).Error(err)
		}
	}()

	var bRmsgId, gwgrId int64
	var code string
	var dBegin, dEnd time.Time
	var direction uint8
	result := make(map[model.BRmsgKey][]model.IdHistItem)

	for rows.Next() {
		if err = rows.Scan(&bRmsgId, &gwgrId, &direction, &code, &dBegin, &dEnd); err != nil {
			return nil, 0, err
		}
		key := model.BRmsgKey{
			GwgrId:    gwgrId,
			Direction: direction,
			Code:      code,
		}

		rbg := model.IdHistItem{
			Id:     bRmsgId,
			DBegin: dBegin.Unix(),
			DEnd:   dEnd.Unix(),
		}
		result[key] = append(result[key], rbg)
		actualRows++
	}
	if expectRows != actualRows {
		return nil, actualRows, fmt.Errorf("GetGroups. expectRows %d != actualRows %d", expectRows, actualRows)
	}

	return result, actualRows, nil
}

func (r *repository) GetAGroups(ctx context.Context) (map[model.ARmsgKey][]model.IdHistItem, int, error) {
	var err error
	var rows *sql.Rows
	var actualRows, expectRows int
	if err = r.db.DB().QueryRowContext(ctx, sqlRACountQuery).Scan(&expectRows); err != nil {
		return nil, actualRows, err
	}

	if expectRows == 0 {
		return nil, 0, nil
	}

	preFetchSize, fetchSize := defaultFetchSize, defaultFetchSize
	if expectRows <= defaultFetchSize {
		preFetchSize = expectRows + 1
		fetchSize = expectRows
	}
	rows, err = r.db.DB().QueryContext(ctx, sqlRAQuery, godror.PrefetchCount(preFetchSize), godror.FetchArraySize(fetchSize))
	if err != nil {
		return nil, actualRows, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"handler": "GetAGroups",
				"problem": "rows close",
			}).Error(err)
		}
	}()

	var bRmsgId, aRmsgId, gwgrId int64
	var code string
	var dBegin, dEnd time.Time
	var direction uint8
	result := make(map[model.ARmsgKey][]model.IdHistItem)
	for rows.Next() {
		if err = rows.Scan(&bRmsgId, &aRmsgId, &gwgrId, &direction, &code, &dBegin, &dEnd); err != nil {
			return nil, actualRows, err
		}
		key := model.ARmsgKey{
			GwgrId:    gwgrId,
			Direction: direction,
			BRmsgId:   bRmsgId,
			Code:      code,
		}

		rag := model.IdHistItem{
			Id:     aRmsgId,
			DBegin: dBegin.Unix(),
			DEnd:   dEnd.Unix(),
		}
		result[key] = append(result[key], rag)
		actualRows++
	}
	if expectRows != actualRows {
		return nil, actualRows, fmt.Errorf("GetAGroups. expectRows %d != actualRows %d", expectRows, actualRows)
	}

	return result, actualRows, nil
}

func (r *repository) GetRates(ctx context.Context) (map[model.RateKey][]model.RmsRateHistItem, int, error) {
	var err error
	var rows *sql.Rows
	var actualRows, expectRows int
	if err = r.db.DB().QueryRowContext(ctx, sqlRTSCountQuery).Scan(&expectRows); err != nil {
		return nil, actualRows, err
	}

	if expectRows == 0 {
		return nil, 0, nil
	}

	preFetchSize, fetchSize := defaultFetchSize, defaultFetchSize
	if expectRows <= defaultFetchSize {
		preFetchSize = expectRows + 1
		fetchSize = expectRows
	}
	rows, err = r.db.DB().QueryContext(ctx, sqlRTSQuery, godror.PrefetchCount(preFetchSize), godror.FetchArraySize(fetchSize))
	if err != nil {
		return nil, actualRows, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"handler": "GetRates",
				"problem": "rows close",
			}).Error(err)
		}
	}()

	var gwgrId, rmsrId, rmsvId, aRmsgId, bRmsgId int64
	var dBegin, dEnd time.Time
	var direction uint8
	result := make(map[model.RateKey][]model.RmsRateHistItem, expectRows)
	for rows.Next() {
		if err := rows.Scan(&gwgrId, &direction, &aRmsgId, &bRmsgId, &rmsrId, &rmsvId, &dBegin, &dEnd); err != nil {
			return nil, actualRows, err
		}
		key := model.RateKey{
			GwgrId:    gwgrId,
			Direction: direction,
			ARmsgId:   aRmsgId,
			BRmsgId:   bRmsgId,
		}
		rt := model.RmsRateHistItem{
			RmsrId: rmsrId,
			RmsvId: rmsvId,
			DBegin: dBegin.Unix(),
			DEnd:   dEnd.Unix(),
		}
		result[key] = append(result[key], rt)
		actualRows++
	}
	if expectRows != actualRows {
		return nil, actualRows, fmt.Errorf("GetRates. expectRows %d != actualRows %d", expectRows, actualRows)
	}

	return result, actualRows, nil
}

func (r *repository) GetRateValues(ctx context.Context) (map[int64]model.Rate, int, error) {
	var err error
	var rows *sql.Rows
	var actualRows, expectRows int
	if err = r.db.DB().QueryRowContext(ctx, sqlRVCountQuery).Scan(&expectRows); err != nil {
		return nil, actualRows, err
	}

	if expectRows == 0 {
		return nil, 0, nil
	}

	preFetchSize, fetchSize := defaultFetchSize, defaultFetchSize
	if expectRows <= defaultFetchSize {
		preFetchSize = expectRows + 1
		fetchSize = expectRows
	}
	rows, err = r.db.DB().QueryContext(ctx, sqlRVQuery, godror.PrefetchCount(preFetchSize), godror.FetchArraySize(fetchSize))
	if err != nil {
		return nil, actualRows, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"handler": "GetRateValues",
				"problem": "rows close",
			}).Error(err)
		}
	}()

	var rmsvId, currencyId int64
	var price float64
	result := make(map[int64]model.Rate, expectRows)
	for rows.Next() {
		if err = rows.Scan(&rmsvId, &currencyId, &price); err != nil {
			return nil, actualRows, err
		}
		rv := model.Rate{
			Price:      price,
			CurrencyId: currencyId,
		}
		result[rmsvId] = rv
		actualRows++
	}
	if expectRows != actualRows {
		return nil, actualRows, fmt.Errorf("GetRateValues. expectRows %d != actualRows %d", expectRows, actualRows)
	}

	return result, actualRows, nil
}

func (r *repository) GetCurrencyRates(ctx context.Context) (map[int64][]model.CurrencyRateHist, int, error) {
	var err error
	var rows *sql.Rows
	var actualRows, expectRows int
	if err = r.db.DB().QueryRowContext(ctx, sqlCURRTSCountQuery).Scan(&expectRows); err != nil {
		return nil, actualRows, err
	}

	if expectRows == 0 {
		return nil, 0, nil
	}
	preFetchSize, fetchSize := defaultFetchSize, defaultFetchSize
	if expectRows <= defaultFetchSize {
		preFetchSize = expectRows + 1
		fetchSize = expectRows
	}
	rows, err = r.db.DB().QueryContext(ctx, sqlCURRTSQuery, godror.PrefetchCount(preFetchSize), godror.FetchArraySize(fetchSize))
	if err != nil {
		return nil, actualRows, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"handler": "GetCurrencyRates",
				"problem": "rows close",
			}).Error(err)
		}
	}()

	var currencyId int64
	var curRate float64
	var dBegin, dEnd time.Time
	result := make(map[int64][]model.CurrencyRateHist, expectRows)
	for rows.Next() {
		if err = rows.Scan(&currencyId, &curRate, &dBegin, &dEnd); err != nil {
			return nil, actualRows, err
		}
		currencyRate := model.CurrencyRateHist{
			CurrencyRate: curRate,
			DBegin:       dBegin.Unix(),
			DEnd:         dEnd.Unix(),
		}
		result[currencyId] = append(result[currencyId], currencyRate)
		actualRows++
	}

	if expectRows != actualRows {
		return nil, actualRows, fmt.Errorf("GetCurrencyRates. expectRows %d != actualRows %d", expectRows, actualRows)
	}
	return result, actualRows, nil
}
