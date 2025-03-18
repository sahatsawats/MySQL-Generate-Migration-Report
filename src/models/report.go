package models


type DatabaseReport struct {
	Host string
	NumberOfDatabase int
	ListOfDatabase []Database
}


type Database struct {
	DatabaseName string
	NumOfTable int
	Tables []Table
}


type Table struct {
	TableName string
	Indexes []string
	Rows int
	Size float32
}
