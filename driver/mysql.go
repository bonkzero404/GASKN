package driver

import (
	"fmt"
	"gaskn/config"
	"strconv"
)

// DsnMySqlDB /*
func DsnMySqlDB() string {
	p := config.Config("DB_PORT")
	port, _ := strconv.ParseUint(p, 10, 32)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		port,
		config.Config("DB_NAME"),
	)

	return dsn
}
