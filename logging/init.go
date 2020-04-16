package logging

import (
	"log"
	"os"
	"path/filepath"
)

var (
	INFO  *log.Logger
	WARN  *log.Logger
	ERROR *log.Logger
	DB    *log.Logger
)

func init() {
	path, _ := os.Getwd()
	appLog, err := os.OpenFile(filepath.Join(path, "log", "application.log"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	dbLog, err := os.OpenFile(filepath.Join(path, "log", "database.log"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	INFO = log.New(appLog, "INFO ", log.LstdFlags|log.Lshortfile)
	WARN = log.New(appLog, "WARN ", log.LstdFlags|log.Lshortfile)
	ERROR = log.New(appLog, "ERROR ", log.LstdFlags|log.Lshortfile)
	DB = log.New(dbLog, "", log.LstdFlags|log.Lshortfile)
}
