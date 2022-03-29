package env

import "os"

const (
	SERVER_PORT = "PORT"
	SERVER      = "SERVER"

	CONNECTION_STRING = "DB_CONN_STRING"
	DATABASE_DRIVER   = "DB_DRIVER"
)

func EnvServerPort() (value string) {
	value = os.Getenv(SERVER_PORT)
	return
}

func EnvServer() (value string) {
	value = os.Getenv(SERVER)
	return
}

func EnvServerConnString() (value string) {
	value = os.Getenv(CONNECTION_STRING)
	return
}

func EnvDBDriver() (value string) {
	value = os.Getenv(DATABASE_DRIVER)
	return
}
