package services

import (
	"database/sql"
	"fmt"

	"github.com/sahatsawats/MySQL-Generate-Migration-Report/src/models"
)



func InitDB(DSN string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	
	db, err = sql.Open("mysql", DSN)
	if err != nil {
		return nil, fmt.Errorf("cannot create connection pool to database at %s, errors: %v", DSN, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot open connection to database at %s, errors: %v", DSN, err)
	}

	return db, nil
}

func GenerateDatabaseReport(conn *sql.DB) error {
	// TODO: generate a report about: total number of database,
	
	return nil
}


func getListOfDatabaseName(conn *sql.DB) ([]models.Database, error){
	var listOfDatabases []models.Database

	queryStatement := "SELECT SCHEMA_NAME FROM information_schema.schemata  WHERE schema_name NOT IN ('mysql', 'information_schema', 'performance_schema', 'sys');"

	rows, err := conn.Query(queryStatement)
	if err != nil {
		return nil, fmt.Errorf("Cannot query database name with errors: ", err)
	}

	for rows.Next() {
		var dbInfo models.Database
		var dbName string
		err = rows.Scan(&dbName)
		if err != nil {
			return nil, fmt.Errorf("Cannot scan database_name from query result with error: ", err)
		}

		dbInfo = models.Database{
			DatabaseName: dbName,
		}

		listOfDatabases = append(listOfDatabases, dbInfo)
		
	}
} 


func getTotalDatabases(conn *sql.DB) (int, error) {
	// Declare variableto be populated from query
	var numOfDatabases int = 0

	queryStatement := "SELECT COUNT(*) AS total_databases FROM information_schema.schemata WHERE schema_name NOT IN ('mysql', 'information_schema', 'performance_schema', 'sys');"
	// Execute query and populate the return value to numOfDatabases
	err := conn.QueryRow(queryStatement).Scan(&numOfDatabases)
	if err != nil {
		return numOfDatabases, fmt.Errorf("cannot execute query statement %s. Error details: %v", queryStatement, err)
	}

	return numOfDatabases, nil
}

