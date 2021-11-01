package commons_go

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	Host   string
	Port   int
	User   string
	Pass   string
	DBName string
}

func SetupDatabase(config *DBConfig, sqls []string) *sql.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.DBName, config.Pass)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}
