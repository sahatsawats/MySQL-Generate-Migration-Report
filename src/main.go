package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"database/sql"
	"github.com/sahatsawats/MySQL-Generate-Migration-Report/src/models"
	"gopkg.in/yaml.v2"
	_ "github.com/go-sql-driver/mysql"
)

func readingConfigurationFile(baseDir string) *models.Configurations {


	// Joining a current execution directory with configuration directory plus file name
	configFile := filepath.Join(filepath.Dir(baseDir), "conf", "config.yaml")

	// Read configuration file
	readConf, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal("Cannot read configuration file at ", configFile, "with error: ", err)
	}

	// Mapping the configuration file .yaml to Configurations structure
	var conf models.Configurations
	err = yaml.Unmarshal(readConf, &conf)
	if err != nil {
		log.Fatal(err)
	}

	return &conf
}


func main() {
	// Find the current executaion directory 
	executionDirectory, err := os.Executable()
	LOG_FILE := filepath.Join(filepath.Dir(executionDirectory), "log", "system.log")

	// Initialize log handler
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	// Set log output to logFile
	log.SetOutput(logFile)
	// log date-time, filename, line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	config := readingConfigurationFile(executionDirectory)

	
}
