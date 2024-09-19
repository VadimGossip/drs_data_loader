package main

import (
	"context"
	"github.com/VadimGossip/drs_data_loader/internal/app"
	"time"

	"github.com/sirupsen/logrus"
)

var appName = "DRS data loader"

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx, appName, time.Now())
	if err != nil {
		logrus.Fatalf("failed to init app[%s]: %s", appName, err)
	}

	if err = a.Run(ctx); err != nil {
		logrus.Infof("app[%s] run process finished with error: %s", appName, err)
	}
}
