package config

type HTTPConfig interface {
	Address() string
}

type OracleConfig interface {
	DSN() string
}
