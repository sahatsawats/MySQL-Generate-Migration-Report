package models

type Configurations struct {
	SOURCE_DATABASE	SourceDatabaseConfigurations
	DESTINATION_DATABASE DestinationDatabaseConfiguration
}


type SourceDatabaseConfigurations struct {
	SOURCE_HOST string
	SOURCE_PORT int
	SOURCE_USER string
	SOURCE_PASSWORD string
}


type DestinationDatabaseConfiguration struct {
	DEST_HOST string
	DEST_PORT int
	DEST_USER string
	DEST_PASSWORD string
}
