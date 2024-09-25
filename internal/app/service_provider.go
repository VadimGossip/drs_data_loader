package app

import (
	"context"
	"log"

	"github.com/VadimGossip/platform_common/pkg/closer"
	"github.com/VadimGossip/platform_common/pkg/db/keydb"
	"github.com/VadimGossip/platform_common/pkg/db/keydb/kdb"
	"github.com/VadimGossip/platform_common/pkg/db/oracle"
	"github.com/VadimGossip/platform_common/pkg/db/oracle/odb"
	"github.com/VadimGossip/platform_common/pkg/db/oracle/transaction"
	"github.com/VadimGossip/platform_common/pkg/db/tarantool"
	"github.com/VadimGossip/platform_common/pkg/db/tarantool/tdb"
	"github.com/sirupsen/logrus"

	"github.com/VadimGossip/drs_data_loader/internal/api/rate"
	"github.com/VadimGossip/drs_data_loader/internal/config"
	dbCfg "github.com/VadimGossip/drs_data_loader/internal/config/db"
	serverCfg "github.com/VadimGossip/drs_data_loader/internal/config/server"
	serviceCfg "github.com/VadimGossip/drs_data_loader/internal/config/service"
	"github.com/VadimGossip/drs_data_loader/internal/repository"
	cacheGatewayRepo "github.com/VadimGossip/drs_data_loader/internal/repository/gateway/cache"
	srcGatewayRepo "github.com/VadimGossip/drs_data_loader/internal/repository/gateway/oracle"
	cacheRateRepo "github.com/VadimGossip/drs_data_loader/internal/repository/rate/cache"
	srcRateRepo "github.com/VadimGossip/drs_data_loader/internal/repository/rate/oracle"
	tarantoolRateRepo "github.com/VadimGossip/drs_data_loader/internal/repository/rate/tarantool"
	"github.com/VadimGossip/drs_data_loader/internal/service"
	gatewayService "github.com/VadimGossip/drs_data_loader/internal/service/gateway"
	rateService "github.com/VadimGossip/drs_data_loader/internal/service/rate"
)

type serviceProvider struct {
	httpConfig            config.HTTPConfig
	grpcConfig            config.GRPCConfig
	oracleConfig          config.OracleConfig
	kdbConfig             config.KdbConfig
	tarantoolConfig       config.TarantoolConfig
	serviceProviderConfig config.ServiceProviderConfig

	oracleClient    oracle.Client
	txManager       oracle.TxManager
	kdbClient       keydb.Client
	tarantoolClient tarantool.Client
	srcRateRepo     repository.SrcRatesRepository
	dstRateRepo     repository.DstRatesRepository
	srcGatewayRepo  repository.SrcGatewayRepository
	dstGatewayRepo  repository.DstGatewayRepository

	rateService    service.RateService
	gatewayService service.GatewayService

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

func (s *serviceProvider) KdbConfig() config.KdbConfig {
	if s.kdbConfig == nil {
		cfg, err := dbCfg.NewKdbConfig()
		if err != nil {
			log.Fatalf("failed to get kdbConfig: %s", err)
		}

		s.kdbConfig = cfg
	}

	return s.kdbConfig
}

func (s *serviceProvider) TarantoolConfig() config.TarantoolConfig {
	if s.tarantoolConfig == nil {
		cfg, err := dbCfg.NewTarantoolConfig()
		if err != nil {
			log.Fatalf("failed to get tarantoolConfig: %s", err)
		}

		s.tarantoolConfig = cfg
	}

	return s.tarantoolConfig
}

func (s *serviceProvider) ServiceProviderConfig() config.ServiceProviderConfig {
	if s.serviceProviderConfig == nil {
		cfg, err := serviceCfg.NewServiceProviderConfig()
		if err != nil {
			log.Fatalf("failed to get serviceProviderConfig: %s", err)
		}

		s.serviceProviderConfig = cfg
	}

	return s.serviceProviderConfig
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

func (s *serviceProvider) KeyDbClient(ctx context.Context) keydb.Client {
	if s.kdbClient == nil {
		cl := kdb.New(kdb.ClientOptions{
			Addr:         s.KdbConfig().Address(),
			Username:     s.KdbConfig().Username(),
			Password:     s.KdbConfig().Password(),
			DB:           s.KdbConfig().DB(),
			ReadTimeout:  s.KdbConfig().ReadTimeoutSec(),
			WriteTimeout: s.KdbConfig().ReadTimeoutSec(),
		})

		if err := cl.DB().Ping(ctx); err != nil {
			log.Fatalf("kdb ping error: %s", err)
		}

		closer.Add(cl.Close)
		s.kdbClient = cl
	}

	return s.kdbClient
}

func (s *serviceProvider) TarantoolClient(ctx context.Context) tarantool.Client {
	if s.tarantoolClient == nil {
		cl, err := tdb.New(ctx, tdb.ClientOptions{
			Addr:     s.TarantoolConfig().Address(),
			Username: s.TarantoolConfig().Username(),
			Password: s.TarantoolConfig().Password(),
			Timeout:  s.TarantoolConfig().Timeout(),
		})

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
		testDb := s.ServiceProviderConfig().DstDB()
		if testDb == s.ServiceProviderConfig().TarantoolTestDB() {
			s.dstRateRepo = tarantoolRateRepo.NewRepository(s.TarantoolClient(ctx))
		} else {
			s.dstRateRepo = cacheRateRepo.NewRepository()
		}
	}
	return s.dstRateRepo
}

func (s *serviceProvider) SrcGatewayRepo(ctx context.Context) repository.SrcGatewayRepository {
	if s.srcGatewayRepo == nil {
		s.srcGatewayRepo = srcGatewayRepo.NewRepository(s.OracleClient(ctx))
	}
	return s.srcGatewayRepo
}

func (s *serviceProvider) DstGatewayRepo(_ context.Context) repository.DstGatewayRepository {
	if s.dstGatewayRepo == nil {
		s.dstGatewayRepo = cacheGatewayRepo.NewRepository()
	}
	return s.dstGatewayRepo
}

func (s *serviceProvider) RateService(ctx context.Context) service.RateService {
	if s.rateService == nil {
		s.rateService = rateService.NewService(s.DstRateRepo(ctx), s.SrcRateRepo(ctx), s.txManager)
	}
	return s.rateService
}

func (s *serviceProvider) GatewayService(ctx context.Context) service.GatewayService {
	if s.gatewayService == nil {
		s.gatewayService = gatewayService.NewService(s.DstGatewayRepo(ctx), s.SrcGatewayRepo(ctx), s.txManager)
	}
	return s.gatewayService
}

func (s *serviceProvider) RateImpl(ctx context.Context) *rate.Implementation {
	if s.rateImpl == nil {
		s.rateImpl = rate.NewImplementation(s.RateService(ctx), s.GatewayService(ctx))
	}

	return s.rateImpl
}
