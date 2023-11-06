package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/a1010s/apiserver/components/components"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {

	// Define red color code
	redColor := "\x1b[31m"
	greenColor := "\x1b[32m"
	resetColor := "\x1b[0m"

	// Define a flag for dbPath
	var path string
	flag.StringVar(&path, "db-path", path, "/path/to/the/database file")
	flag.Parse()

	// Check if db-path flag is not set
	if path == "" {
		fmt.Println(redColor + "Error: Please provide a valid --db-path flag." + resetColor)
		fmt.Println(greenColor + "NOTE: If there is no database, it would be created on the given path" + resetColor)
		flag.PrintDefaults() // Print default flag values and usage information
		os.Exit(1)           // Exit with an error status code
	}

	// Get the value of the dbPath flag
	dbPath := path

	router := gin.Default()

	components.InitDB(dbPath) // Initialize the database with the provided dbPath
	components.SetupRoutes(router)

	router.Run("0.0.0.0:8088")
}
