package tdb

import (
	db "drs_data_loader/internal/client/db/tarantool"

	"github.com/tarantool/go-tarantool/v2"
)

type tdb struct {
	dbc *tarantool.Connection
}

func NewDB(dbc *tarantool.Connection) db.DB {
	return &tdb{
		dbc: dbc,
	}
}

func (t *tdb) Do(req tarantool.Request) *tarantool.Future {
	return t.dbc.Do(req)
}

func (t *tdb) Close() error {
	return t.dbc.Close()
}
