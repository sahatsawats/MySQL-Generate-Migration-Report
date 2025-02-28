package models
import (
	"errors"
	"fmt"
)
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

func (c DatabaseProperties) GetDSNConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)", c.User, c.Password, c.Host, c.Port,)
}