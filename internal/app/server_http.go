package app

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (a *App) initHTTPServer(_ context.Context) error {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.RedirectTrailingSlash = false
	a.httpServer = &http.Server{
		Addr:           a.serviceProvider.HTTPConfig().Address(),
		Handler:        g,
		ReadTimeout:    30 * time.Minute,
		WriteTimeout:   30 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	return nil
}

func (a *App) runHTTPServer() error {
	logrus.Infof("[%s] HTTP server is running on %s", a.name, a.httpServer.Addr)
	return a.httpServer.ListenAndServe()
}

func (a *App) closeHTTPServer(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return a.httpServer.Shutdown(shutdownCtx)
}
