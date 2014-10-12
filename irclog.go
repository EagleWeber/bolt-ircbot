package main

import (
	//"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

const RFC3339_SECONDS = "2006-01-02 03:04:05-07:00"

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func CreateFile(name string) error {

	fo, err := os.Create(name)
	if err != nil {
		return err
	}

	defer func() {
		fo.Close()
	}()

	return nil
}

func StartLogger(c *Config, channel string) *os.File {

	logfile := fmt.Sprintf("%v/log-%v.log", c.Logging.Location, channel)

	if !FileExists(logfile) {
		CreateFile(logfile)
	}

	logger, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Println(fmt.Sprintf("Error opening log file: %v", err))
	}

	return logger
}

func WriteLog(c *Config, logger *os.File, nick string, text string) {

	t1 := time.Now()
	f1 := t1.Format(RFC3339_SECONDS)
	line := fmt.Sprintf("%v\t<%v>\t%v\n", f1, nick, text)

	_, err := logger.WriteString(line)

	if err != nil {
		log.Println(fmt.Sprintf("Tried to write: %v", line))
		log.Println(fmt.Sprintf("Error writing log string: %v", err))
	}

	logger.Sync()
}
