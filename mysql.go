package mysql_connection_sample

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"fmt"
)

var connection *sql.DB
var once sync.Once

func GetConnection() *sql.DB {
	once.Do(func() {
		dbip := "127.0.0.1:3306"
		dbuser := "user"
		dbpassword := "password"
		var conStr string
		conStr = dbuser + ":" + dbpassword + "@" + "tcp(" + dbip + ")/?charset=utf8"
		var errs error
		connection, errs = sql.Open("mysql", conStr)
		if errs != nil {
			fmt.Println(errs)
		}
		connection.SetMaxOpenConns(10)
		err := connection.Ping()
		if err != nil {
			fmt.Println("MySQL connection error / Cannot connect to mysql")
		}
		fmt.Println("Connected to MySQL")
	})
	return connection
}