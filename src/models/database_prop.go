package models
import "errors"

type DatabaseProperties struct {
	Host string
	Port int
	User string
	Password string
}

func (c DatabaseProperties) CheckValidDatabaseProperties() error {
	if (c.Host == "") {
		return errors.New("host variable cannot be empty")
	}
	if (c.User == "") {
		return errors.New("user variable cannot be empty")
	}
	return nil
}