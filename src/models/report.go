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
	Indexes []Index
}


type Table struct {
	TableName string
	Rows int
	Size float32
}

type Index struct {
	IndexName string
	Size float32
}