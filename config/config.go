package config

import "fmt"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456f"
	dbname   = "postgres"
)

func GetPsqlInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}