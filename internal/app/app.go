package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/VadimGossip/platform_common/pkg/closer"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type App struct {
	serviceProvider *serviceProvider
	name            string
	appStartedAt    time.Time
	httpServer      *http.Server
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context, name string, appStartedAt time.Time) (*App, error) {
	a := &App{
		name:         name,
		appStartedAt: appStartedAt,
	}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initEnv,
		a.initServiceProvider,
		a.initHTTPServer,
		a.initGRPCServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initEnv(_ context.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := a.runHTTPServer(); !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("[%s] failed to run HTTP server: %v", a.name, err)
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
			logrus.Infof("shuting down httpServer")
			if err := a.closeHTTPServer(ctx); err != nil {
				logrus.Errorf("[%s] failed to shutdown HTTP server: %v", a.name, err)
			}
		}
	}()

	go func() {
		defer wg.Done()
		if err := a.runGRPCServer(); !errors.Is(err, grpc.ErrServerStopped) {
			logrus.Fatalf("%s failed to run GRPC server: %v", a.name, err)
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
			logrus.Infof("shuting down grpcServer")
			a.grpcServer.GracefulStop()
		}
	}()

	if err := a.serviceProvider.RateService(ctx).Refresh(ctx); err != nil {
		logrus.Errorf("[%s] failed to rate service: %v", a.name, err)
	}

	gracefulShutdown(ctx, cancel, wg)
	return nil
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		logrus.Info("terminating: context cancelled")
	case c := <-waitSignal():
		logrus.Infof("terminating: got signal: [%s]", c)
	}

	cancel()
	if wg != nil {
		wg.Wait()
	}
}

func waitSignal() chan os.Signal {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	return sigterm
}
