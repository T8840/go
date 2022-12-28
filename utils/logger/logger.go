package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var Logger *log.Logger = nil
func SetLogFile(log_name string) {
	dir, _ := os.Getwd()
	comma := strings.Index(dir, "godemo")
	pre_path := dir[:comma]
	log_path := pre_path+"godemo/logs"
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logFileName := fmt.Sprintf("%s/%s.log", log_path,log_name)
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatalln("open log file error !")
	}
	Logger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)
}

