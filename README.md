# drs_data_loader
Loads data from the source database, into the receiver database, the source is currently available one oracle. Receiver, kdb, tarantool, cache

For simplicity, I load environment variables from a file that I don't store in the git.
Example of variables from a file

APP_HTTP_HOST=
APP_HTTP_PORT=

APP_GRPC_HOST=
APP_GRPC_PORT=

ORACLE_HOST=
ORACLE_PORT=
ORACLE_USERNAME=
ORACLE_PASSWORD=
ORACLE_SERVICE=

TARANTOOL_HOST=
TARANTOOL_PORT=
TARANTOOL_USERNAME=
TARANTOOL_PASSWORD=
TARANTOOL_TIMEOUT_MIN=

#DST_DB TARANTOOL|KDB|CACHE
DST_DB=