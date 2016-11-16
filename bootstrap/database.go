package bootstrap

import "database/sql"

// GetDatabaseConnection initializes a connection to the database
// and verifies that it works.
func GetDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "abraham_linkedin:seattlelife12345@tcp("+
		"abrahamlinkedin.ccwrpjrudb04.us-east-1.rds.amazonaws.com:3306)/"+
		"AbrahamLinkedIn")

	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}

	return db
}
