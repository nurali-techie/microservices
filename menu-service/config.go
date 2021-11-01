package main

import (
	_ "embed"
	"strings"

	commons "github.com/nurali-techie/microservices/commons-go"
)

//go:embed database.sql
var sql string

var (
	DB_HOST = "menu_postgres"
	DB_PORT = 5432
	DB_USER = "postgres"
	DB_PASS = "postgres"
	DB_NAME = "menu_database"
)

func getDBConfig() *commons.DBConfig {
	return &commons.DBConfig{
		Host:   DB_HOST,
		Port:   DB_PORT,
		User:   DB_USER,
		Pass:   DB_PASS,
		DBName: DB_NAME,
	}
}

func getDatabaseSQL() []string {
	return strings.Split(sql, "\n")
}
