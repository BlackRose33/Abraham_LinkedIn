package bootstrap

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"abraham_linkedin/constants"
	"abraham_linkedin/controllers"
)

// GlobalStart begins initialization for the application,
// and notifies main() if an error occurrs.
func GlobalStart() bool {

	fmt.Println("[STARTUP] Loading Environment Settings")
	constants.LoadEnvironmentSettings()

	// Seed for random generators
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("[STARTUP] Loading Templates")
	templates := GetTemplates()

	fmt.Println("[STARTUP] Connecting to DB")
	db := GetDatabaseConnection()

	fmt.Println("[STARTUP] Initializing Services")
	controllers.BaseInitialization(templates, db)
	constants.BaseInitialization()

	fmt.Println("[STARTUP] Creating Routes")
	CreateRoutes()

	fmt.Println("[STARTUP] Starting server on port " +
		strconv.Itoa(constants.PortNum))

	http.ListenAndServe(":"+strconv.Itoa(constants.PortNum), nil)

	// Shouldn't return, I don't believe..
	fmt.Println("[STARTUP] Startup failed.")
	return false
}
