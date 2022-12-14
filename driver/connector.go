package driver

import (
	"fmt"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/utils"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB /*
var DB *gorm.DB

func setupSqlLog() *gorm.Config {
	var gormConfig gorm.Config

	if config.Config("ENABLE_LOG") == "true" {
		var capLog logger.Writer

		if config.Config("ENABLE_WRITE_TO_FILE_LOG") == "true" {
			multiOutput := utils.CreateSqlLog()
			capLog = log.New(multiOutput, "[SQL][ERROR] ", log.LstdFlags)
		} else {
			capLog = log.New(os.Stdout, "[SQL][ERROR] ", log.LstdFlags)
		}

		newLogger := logger.New(
			capLog, // io writer
			logger.Config{
				SlowThreshold:             time.Second,  // Slow SQL threshold
				LogLevel:                  logger.Error, // Log level
				IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,        // Disable color
			},
		)

		gormConfig = gorm.Config{
			Logger: newLogger,
		}
	} else {
		gormConfig = gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	return &gormConfig
}

// ConnectDB /*
func ConnectDB() *gorm.DB {
	var err error
	var dbType string

	// check whether in the configuration using mysql or postgresql driver
	if config.Config("DB_DRIVER") == "mysql" {
		// Open MySQL connection
		DB, err = gorm.Open(mysql.Open(DsnMySqlDB()), setupSqlLog())
	} else if config.Config("DB_DRIVER") == "pgsql" {
		// Open PostgreSQL connection
		DB, err = gorm.Open(postgres.Open(DsnPostgreSqlDB()), setupSqlLog())
	} else {
		// Stop the application if the driver does not match
		panic("Database driver not available")
	}

	// Display an error message if an error occurs in the database connection
	dbType = config.Config("DB_DRIVER")
	if err != nil {
		errMessage := fmt.Sprintf("Failed to connect database %s", dbType)
		panic(errMessage)
	}

	// Call db pooling function
	errPool := dbPooling(DB)
	if errPool != nil {
		panic("Error database pooling")
	}

	return DB
}

/*
*
This function is for database pooling
*/
func dbPooling(sqlDb *gorm.DB) error {
	// Get generic database object sql.DB to use its functions
	sqlDB, err := sqlDb.DB()

	if err != nil {
		panic("failed to connect database")
	}

	// Get param config into var
	maxIdleConsConf := config.Config("DB_MAX_IDLE_CONNS")
	maxOpenConsConf := config.Config("DB_MAX_OPEN_CONNS")

	// Convert string to integer
	maxIdleCons, _ := strconv.Atoi(maxIdleConsConf)
	maxOpenCons, _ := strconv.Atoi(maxOpenConsConf)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdleCons)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpenCons)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
