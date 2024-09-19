package app

import (
	"context"
	"drs_data_loader/internal/client/db/tarantool"
	"drs_data_loader/internal/client/db/tarantool/tdb"
	"drs_data_loader/internal/repository"
	"drs_data_loader/internal/service"
	"drs_data_loader/internal/service/rate"
	"log"

	"github.com/VadimGossip/platform_common/pkg/db/oracle"
	"github.com/VadimGossip/platform_common/pkg/db/oracle/odb"
	"github.com/VadimGossip/platform_common/pkg/db/oracle/transaction"
	"github.com/sirupsen/logrus"

	"drs_data_loader/internal/closer"
	"drs_data_loader/internal/config"
	dbCfg "drs_data_loader/internal/config/db"
	serverCfg "drs_data_loader/internal/config/server"
	srcRateRepo "drs_data_loader/internal/repository/rate/oracle"
	dstRateRepo "drs_data_loader/internal/repository/rate/tarantool"
)

type serviceProvider struct {
	httpConfig   config.HTTPConfig
	oracleConfig config.OracleConfig

	oracleClient    oracle.Client
	txManager       oracle.TxManager
	tarantoolClient tarantool.Client
	srcRateRepo     repository.SrcRatesRepository
	dstRateRepo     repository.DstRatesRepository

	rateService service.RateService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := serverCfg.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get httpConfig: %s", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) OracleConfig() config.OracleConfig {
	if s.oracleConfig == nil {
		cfg, err := dbCfg.NewOracleConfig()
		if err != nil {
			log.Fatalf("failed to get pgConfig: %s", err)
		}

		s.oracleConfig = cfg
	}

	return s.oracleConfig
}

func (s *serviceProvider) OracleClient(ctx context.Context) oracle.Client {
	if s.oracleClient == nil {
		cl, err := odb.New(s.OracleConfig().DSN())
		if err != nil {
			logrus.Fatalf("failed to create oracle client: %s", err)
		}

		if err = cl.DB().Ping(ctx); err != nil {
			log.Fatalf("ping error: %s", err)
		}
		closer.Add(cl.Close)
		s.oracleClient = cl
	}

	return s.oracleClient
}

func (s *serviceProvider) TxManager(ctx context.Context) oracle.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.OracleClient(ctx).DB())
	}
	return s.txManager
}

func (s *serviceProvider) TarantoolClient(ctx context.Context) tarantool.Client {
	if s.tarantoolClient == nil {
		cl, err := tdb.New(ctx, "todo_config")
		if err != nil {
			logrus.Fatalf("failed to create tarantool client: %s", err)
		}
		closer.Add(cl.Close)
		s.tarantoolClient = cl
	}

	return s.tarantoolClient
}

func (s *serviceProvider) SrcRateRepo(ctx context.Context) repository.SrcRatesRepository {
	if s.srcRateRepo == nil {
		s.srcRateRepo = srcRateRepo.NewRepository(s.OracleClient(ctx))
	}
	return s.srcRateRepo
}

func (s *serviceProvider) DstRateRepo(ctx context.Context) repository.DstRatesRepository {
	if s.dstRateRepo == nil {
		s.dstRateRepo = dstRateRepo.NewRepository(s.TarantoolClient(ctx))
	}
	return s.dstRateRepo
}

func (s *serviceProvider) RateService(ctx context.Context) service.RateService {
	if s.rateService == nil {
		s.rateService = rate.NewService(s.DstRateRepo(ctx), s.SrcRateRepo(ctx), s.txManager)
	}
	return s.rateService
}
