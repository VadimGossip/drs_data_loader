package app

import (
	"context"
	"log"

	"github.com/VadimGossip/drs_data_loader/internal/client/db/tarantool"
	"github.com/VadimGossip/drs_data_loader/internal/client/db/tarantool/tdb"
	"github.com/VadimGossip/platform_common/pkg/db/oracle"
	"github.com/VadimGossip/platform_common/pkg/db/oracle/odb"
	"github.com/VadimGossip/platform_common/pkg/db/oracle/transaction"
	"github.com/sirupsen/logrus"

	"github.com/VadimGossip/drs_data_loader/internal/api/rate"
	"github.com/VadimGossip/drs_data_loader/internal/closer"
	"github.com/VadimGossip/drs_data_loader/internal/config"
	dbCfg "github.com/VadimGossip/drs_data_loader/internal/config/db"
	serverCfg "github.com/VadimGossip/drs_data_loader/internal/config/server"
	"github.com/VadimGossip/drs_data_loader/internal/repository"
	dstRateCacheRepo "github.com/VadimGossip/drs_data_loader/internal/repository/rate/cache"
	srcRateRepo "github.com/VadimGossip/drs_data_loader/internal/repository/rate/oracle"
	"github.com/VadimGossip/drs_data_loader/internal/service"
	rateService "github.com/VadimGossip/drs_data_loader/internal/service/rate"
)

type serviceProvider struct {
	httpConfig   config.HTTPConfig
	grpcConfig   config.GRPCConfig
	oracleConfig config.OracleConfig

	oracleClient    oracle.Client
	txManager       oracle.TxManager
	tarantoolClient tarantool.Client
	srcRateRepo     repository.SrcRatesRepository
	dstRateRepo     repository.DstRatesRepository

	rateService service.RateService

	rateImpl *rate.Implementation
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

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := serverCfg.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpcConfig: %s", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
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

func (s *serviceProvider) DstRateRepo(_ context.Context) repository.DstRatesRepository {
	if s.dstRateRepo == nil {
		s.dstRateRepo = dstRateCacheRepo.NewRepository()
	}
	return s.dstRateRepo
}

func (s *serviceProvider) RateService(ctx context.Context) service.RateService {
	if s.rateService == nil {
		s.rateService = rateService.NewService(s.DstRateRepo(ctx), s.SrcRateRepo(ctx), s.txManager)
	}
	return s.rateService
}

func (s *serviceProvider) RateImpl(ctx context.Context) *rate.Implementation {
	if s.rateImpl == nil {
		s.rateImpl = rate.NewImplementation(s.RateService(ctx))
	}

	return s.rateImpl
}
