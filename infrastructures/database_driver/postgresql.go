package database_driver

import (
	"fmt"
	"github.com/bonkzero404/gaskn/config"
	"strconv"
)

// DsnPostgreSqlDB /*
func DsnPostgreSqlDB() string {
	p := config.Config("DB_PORT")
	port, _ := strconv.ParseUint(p, 10, 32)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)

	return dsn
}
