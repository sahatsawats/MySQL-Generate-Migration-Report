package services

import (
	"github.com/sahatsawats/MySQL-Generate-Migration-Report/src/models"
)

/*
Got pointer of source and destination database report, write the report to csv extension.
Each catagories sorted to ensure the sequence.
Return error or nil
*/
func saveReportInCSV(sourceReport *models.DatabaseReport, destReport *models.DatabaseReport, outputCSVFile string) (error) {
	
	return nil
}

/*
Save a summary information about the migration report. Such as total databases, tables, indexes, and rows.
Each catagories will be show the numerical number and equivalence.
Return error or nil
*/
func saveSummaryReportInText(sourceReport *models.DatabaseReport, destReport *models.DatabaseReport, outputTextFile string) (error) {
	
	return nil
}