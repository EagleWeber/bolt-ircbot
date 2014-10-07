package main

import (
	"log"
	"os"
)

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

const logfile = "var/log/irc.log"

func WriteLog(text string) {
	if !FileExists(logfile) {
		CreateFile(logfile)
	}

	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Error opening log file: %v", err)
	}

	defer f.Close()

	log.SetOutput(f)
	log.Println(text)
}
