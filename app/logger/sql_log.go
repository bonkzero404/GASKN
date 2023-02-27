package logger

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

func createSqlLogFile() io.Writer {
	logFile := utils.CreateFile(config.Config("LOG_LOCATION"), config.Config("LOG_SQL_ERROR_FILENAME"))
	multiOutput := utils.MultiWrite(os.Stdout, logFile)

	return multiOutput
}

func CreateSqlLog() *gorm.Config {
	var gormConfig gorm.Config

	if config.Config("ENABLE_LOG") == "true" {
		var capLog logger.Writer

		if config.Config("ENABLE_WRITE_TO_FILE_LOG") == "true" {
			multiOutput := createSqlLogFile()
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
