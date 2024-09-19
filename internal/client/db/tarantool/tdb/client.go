package tdb

import (
	"context"
	"time"

	"github.com/tarantool/go-tarantool/v2"

	db "drs_data_loader/internal/client/db/tarantool"
)

type tdbClient struct {
	masterDBC db.DB
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dialer := tarantool.NetDialer{
		Address: "192.168.244.157:3301",
		User:    "guest",
	}
	opts := tarantool.Opts{
		Timeout: 30 * time.Minute,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {
		return nil, err
	}

	return &tdbClient{
		masterDBC: NewDB(conn),
	}, nil
}

func (c *tdbClient) DB() db.DB {
	return c.masterDBC
}

func (c *tdbClient) Close() error {
	if c.masterDBC != nil {
		return c.masterDBC.Close()
	}

	return nil
}
