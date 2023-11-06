package components

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB(dbPath string) {
	log.Println("Initializing database...")
	log.Println(dbPath)
	var err error
	db, err = sql.Open("sqlite", dbPath) // NOTE: /app/data/links.db FOR DockerImage.
	if err != nil {                      // /tmp/links.db localtesting (PERMISSION CHECK)
		panic(err) // NOTE: before using dbPath var, it was the link
	} // ("sqlite", "/tmp/links.db")

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS links (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            description TEXT NOT NULL,
            url TEXT NOT NULL
        );
    `)
	if err != nil {
		panic(err)
	}
}
