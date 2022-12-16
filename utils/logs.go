package utils

import (
	"encoding/json"
	"fmt"
	"github.com/bonkzero404/gaskn/config"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateSqlLog() io.Writer {
	logFile := CreateFile(config.Config("LOG_LOCATION"), config.Config("LOG_SQL_ERROR_FILENAME"))
	multiOutput := MultiWrite(os.Stdout, logFile)

	return multiOutput
}

//goland:noinspection GoUnhandledErrorResult,GoUnhandledErrorResult,GoUnhandledErrorResult,GoUnusedFunction
func getLastLineWithSeek(filepath string, lineFromBottom int) string {
	fileHandle, err := os.Open(filepath)

	if err != nil {
		panic("Cannot open file")
	}
	defer fileHandle.Close()

	line := ""
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		fileHandle.Seek(cursor, io.SeekEnd-lineFromBottom)

		char := make([]byte, 1)
		fileHandle.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	if lineFromBottom > 0 {
		return reverseString(line)
	}

	return line
}

//goland:noinspection GoUnhandledErrorResult
func WriteRequestToLog(ctx *fiber.Ctx, ptr string, statusCode int, resp any) {

	if config.Config("ENABLE_LOG") == "true" {
		logFormat := ptr +
			" " +
			time.Now().Format("2006/01/02 15:04:05") +
			" " +
			ctx.IP() +
			" " +
			ctx.Method() +
			" " +
			strconv.Itoa(statusCode) +
			" " +
			"ROUTE=" + ctx.Route().Path

		if ctx.Request().URI().QueryString() != nil {
			logFormat = logFormat + " QUERY_URL=" + string(ctx.Request().URI().QueryString())
		}

		if ctx.Body() != nil {
			body := string(ctx.Request().Body())

			helper := make(map[string]interface{})

			err := json.Unmarshal([]byte(body), &helper)
			if err == nil {
				bytes, err := json.Marshal(helper)
				if err == nil {
					logFormat = logFormat + " PAYLOAD=" + string(bytes)
				}
			}
		}

		bytes, err := json.Marshal(resp)
		if err == nil {
			logFormat = logFormat + " RESPONSE=" + string(bytes)
		}

		if config.Config("ENABLE_WRITE_TO_FILE_LOG") == "true" {
			CreateFile(config.Config("LOG_LOCATION"), config.Config("LOG_ACCESS_FILENAME"))

			f, err := os.OpenFile(config.Config("LOG_LOCATION")+config.Config("LOG_ACCESS_FILENAME"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(logFormat + "\r\n"); err != nil {
				log.Println(err)
			}
		}

		log.Println(logFormat)
	}
}

func reverseString(str string) string {
	byteStr := []rune(str)
	for i, j := 0, len(byteStr)-1; i < j; i, j = i+1, j-1 {
		byteStr[i], byteStr[j] = byteStr[j], byteStr[i]
	}
	return string(byteStr)
}
