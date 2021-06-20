package repository

import (
	"fmt"
	"os"
	"testing"
)

func TestGetDSN(t *testing.T) {
	expected := fmt.Sprintf("host='%s' port='%s' user='%s' password='%s' dbname='%s' sslmode='disable'",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))

	result := GetDSN()
	if expected != result {
		t.Errorf("Incorrect DB DSN: %s, expected: %s", result, expected)
	}
}
