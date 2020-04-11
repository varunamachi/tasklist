package db 

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// ConnOpts - options for connecting to a postgres instance
type ConnOpts struct {
	Host string
	Port int
	User string
	Password string
	Database string
}

// String - converts ConnOpts to appropriate connection string
func (co *ConnOpts) String() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		co.Host,
		co.Port,
		co.User,
		co.Password,
		co.Database)
}

// Connect - using the given options, connects to a database
func Connect(opts *ConnOpts) error {
	var err error
	db, err = sqlx.Connect("postgres", opts.String())
	return err
}

// Conn - gives a connection to database if it is initialized, otherwise returns 
// nil
func Conn() *sqlx.DB {
	return db
}
