package config

import "time"

type HTTPConfig interface {
	Address() string
}

type GRPCConfig interface {
	Address() string
}

type OracleConfig interface {
	DSN() string
}

type KdbConfig interface {
	Address() string
	Username() string
	Password() string
	DB() int
	ReadTimeoutSec() time.Duration
	WriteTimeoutSec() time.Duration
}

type TarantoolConfig interface {
	Address() string
	Username() string
	Password() string
	Timeout() time.Duration
}

type ServiceProviderConfig interface {
	DstDB() string
	TarantoolTestDB() string
	KdbTestDB() string
	CacheTestDB() string
}
