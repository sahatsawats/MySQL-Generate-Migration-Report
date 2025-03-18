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

func CloseDB(conn *sql.DB) (error) {
	err := conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close a database: %v", err)
	}

	return nil
}

func GenerateDatabaseReport(conn *sql.DB, HOST string) (models.DatabaseReport, error) {
	var databaseReport models.DatabaseReport
	var host string = HOST
	var numberOfDatabase int


	// Query a list of schema within instance
	databases, err := getListOfDatabaseName(conn)
	if err != nil {
		return databaseReport, err
	}

	// Get a total number of database
	numberOfDatabase, err = getTotalDatabases(conn)
	if err != nil {
		return databaseReport, err
	}

	// Get a list of database from mysql instance
	listOfDatabase, err := getDatabaseConstruct(conn, databases)
	if err != nil {
		return databaseReport, err
	}

	databaseReport = models.DatabaseReport{
		Host: host,
		NumberOfDatabase: numberOfDatabase,
		ListOfDatabase: listOfDatabase,
	}


	return databaseReport, nil
}


// only get a list of database.
func getListOfDatabaseName(conn *sql.DB) ([]string, error) {
	// variable to hold list of database name and number of database occured
	var listOfDatabases []string
	queryStatement := "SELECT SCHEMA_NAME FROM information_schema.schemata  WHERE schema_name NOT IN ('mysql', 'information_schema', 'performance_schema', 'sys');"

	rows, err := conn.Query(queryStatement)
	if err != nil {
		return nil, fmt.Errorf("cannot query database name with errors: %v", err)
	}

	for rows.Next() {
		var dbName string
		err = rows.Scan(&dbName)
		if err != nil {
			return nil, fmt.Errorf("cannot scan database_name from query result with error: %v", err)
		}

		listOfDatabases = append(listOfDatabases, dbName)

	}

	return listOfDatabases, nil
} 

/*
Used this function to get all of information and map to a Database model.
Receive connection pool and list of databases from host.
Return a list of Database model
*/
func getDatabaseConstruct(conn *sql.DB, listOfDatabases []string) ([]models.Database, error) {


	var DBs []models.Database
	// Loop over database. Each loop create a Database struct and append to "DBs".
	for _, database := range listOfDatabases {
		var DB models.Database
		var databaseName string = database
		var numOfTable int
		var tables []models.Table

		// Query a list of tables within given database.
		listOfTableName, rows, err := getListTableNameFromDatabase(conn, &databaseName)
		if err != nil {
			return nil, err
		}

		// Assign a total number of table in database.
		numOfTable = rows

		// Loop over the list of table. Each loop create a Table structure and append to "tables".
		for _, tableName := range listOfTableName {
			// tableProperties: used to hold a Table struct for each table in loop
			var tableProperties models.Table
			// hold a metadata of table in loop
			var indexes []string
			var numRows int
			var size float32
			var err error

			numRows, err = getRowsFromTable(conn, &databaseName, &tableName)
			if err != nil {
				return nil, err
			}

			size, err = getSizeOfTable(conn, &databaseName, &tableName)
			if err != nil {
				return nil, err
			}
	
			indexes, err = getListOfIndexes(conn, &databaseName, &tableName)
			if err != nil {
				return nil, err
			}

			// Combined those query results to structure
			tableProperties = models.Table{
				TableName: tableName,
				Rows: numRows,
				Size: size,
				Indexes: indexes,
			}

			// Append structure to list
			tables = append(tables, tableProperties)
		}

		DB = models.Database{
			DatabaseName: databaseName,
			NumOfTable: numOfTable,
			Tables: tables,
		}

		DBs = append(DBs, DB)
	}

	return DBs, nil

}


func getListTableNameFromDatabase(conn *sql.DB, database *string) ([]string, int, error) {
	// Declare variables.
	var listOfTableName []string
	var numberOfTable int = 0

	// Query statement to query list of tables.
	queryStatement := fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_schema = '%s'", *database)
	rows, err := conn.Query(queryStatement)
		
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %v", err)
	}

	// Loop through query results. This loop will used for get all of table name.
	for rows.Next() {
		var tableName string
		// Incremental the number of table that occured in loop.
		numberOfTable = numberOfTable + 1
		// Assign results from query to variable.
		
		rows.Scan(&tableName)
		listOfTableName = append(listOfTableName, tableName)
	}
	rows.Close()

	return listOfTableName, numberOfTable, nil
}

/*
Do a query from database to get a total rows from given database and table name.
Return a integer of rows and error from query.
*/
func getRowsFromTable(conn *sql.DB, databaseName *string, tableName *string) (int, error) {
	var rows int // Hold a total rows of given table from query result.

	queryStatement := fmt.Sprintf("SELECT COUNT(*) FROM %s.%s", *databaseName, *tableName)
	err := conn.QueryRow(queryStatement, &rows)
	if err != nil {
		return 0, fmt.Errorf("failed to query total rows of %s : %v", *tableName, err)
	}

	return rows, nil
}

/*
Do a query from database to get a total size in mb from given database and table name.
Return a float32 of size and error from query.
*/
func getSizeOfTable(conn *sql.DB, databaseName *string, tableName *string) (float32, error) {
	var size float32 // Hold a total size of given table from query resutls.

	queryStatement := fmt.Sprintf("SELECT ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS table_size_mb FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s'", *databaseName, *tableName)
	err := conn.QueryRow(queryStatement, &size)
	if err != nil {
		return 0.00, fmt.Errorf("failed to query size of %s : %v", *tableName, err)
	}

	return size, nil
}

/*
Do a query from database to get a total size in mb from given database and table name.
Return a float32 of size and error from query.
*/
func getListOfIndexes(conn *sql.DB, databaseName *string, tableName *string) ([]string, error) {
	var listOfIndexes []string // Hold list of index name within table from query results

	queryStatement := fmt.Sprintf("SELECT index_name FROM information_schema.statistics WHERE table_schema = '%s' AND table_name = '%s'", *databaseName, *tableName)
	rows, err := conn.Query(queryStatement)
	// Check a error from establish query command
	if err != nil {
		return nil, fmt.Errorf("failed to query indexes from table %s : %v", *tableName, err)
	}

	for rows.Next() {
		var indexName string // Hold individual results for each index name

		err := rows.Scan(&indexName)
		// Check a error from map a query results to variable
		if err != nil {
			return nil, fmt.Errorf("failed to scan query results from indexes: %v", err)
		}

		// Append a results to list of indexes name
		listOfIndexes = append(listOfIndexes, indexName)
	}

	rows.Close()
	return listOfIndexes, nil
}

func getTotalDatabases(conn *sql.DB) (int, error) {
	// Declare variable to be populated from query.
	var numOfDatabases int = 0

	queryStatement := "SELECT COUNT(*) AS total_databases FROM information_schema.schemata WHERE schema_name NOT IN ('mysql', 'information_schema', 'performance_schema', 'sys');"
	// Execute query and populate the return value to numOfDatabases.
	err := conn.QueryRow(queryStatement).Scan(&numOfDatabases)
	if err != nil {
		return numOfDatabases, fmt.Errorf("cannot execute query statement %s. Error details: %v", queryStatement, err)
	}

	return numOfDatabases, nil
}

