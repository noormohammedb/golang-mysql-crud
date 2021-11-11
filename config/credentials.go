package config

import "fmt"

const (
	DbUser     = "dbuser"
	DbPassword = "dbpass"
	DbAddress  = "127.0.0.1"
	DbName     = "golang"
)

func DbCredentials() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", DbUser, DbPassword, DbAddress, DbName)
}
