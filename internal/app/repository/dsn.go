package repository

import (
	"fmt"
	"os"
)

func GetDSN()  (dbDSN string){
	dbDSN = fmt.Sprintf("host='%s' port='%s' user='%s' password='%s' dbname='%s' sslmode='disable'",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))
	return dbDSN
}