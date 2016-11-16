package main

import (
	"abraham_linkedin/bootstrap"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if !bootstrap.GlobalStart() {
		panic("Failed to start server.")
	}
}
