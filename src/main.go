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
	var sourceDB models.DatabaseProperties
	var destDB models.DatabaseProperties
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

	// Read configuration file
	config := readingConfigurationFile(executionDirectory)

	log.Println("Reading databases configuration...")
	sourceDB = models.DatabaseProperties{
		Host: config.SOURCE_DATABASE.SOURCE_HOST,
		Port: config.SOURCE_DATABASE.SOURCE_PORT,
		User: config.SOURCE_DATABASE.SOURCE_USER,
		Password: config.SOURCE_DATABASE.SOURCE_PASSWORD,
	}
	destDB = models.DatabaseProperties{
		Host: config.DESTINATION_DATABASE.DEST_HOST,
		Port: config.DESTINATION_DATABASE.DEST_PORT,
		User: config.DESTINATION_DATABASE.DEST_USER,
		Password: config.DESTINATION_DATABASE.DEST_PASSWORD,
	}
	log.Printf("Source database properties: %s:%d \n", sourceDB.Host, sourceDB.Port)
	log.Printf("Destination database properties: %s:%d \n", destDB.Host, destDB.Port)

	
}
