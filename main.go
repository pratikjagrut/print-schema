package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Retrieve the database URL from the environment variable
	dbURL := os.Getenv("DB_URL")

	// Open a connection to the database
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the schema of the "example" table
	rows, err := db.Query("DESCRIBE example")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Store the schema data
	var (
		field        string
		typ          string
		null         string
		key          string
		defaultValue sql.NullString
		extra        string
	)

	var schema string

	// Iterate over the rows and build the schema string
	for rows.Next() {
		err := rows.Scan(&field, &typ, &null, &key, &defaultValue, &extra)
		if err != nil {
			log.Fatal(err)
		}

		schema += fmt.Sprintf("Field: %s, Type: %s, Null: %s, Key: %s, Default: %s, Extra: %s\n",
			field, typ, null, key, defaultValue.String, extra)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Create an HTTP handler function to display the schema in the web page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<pre>%s</pre>", schema)
	})

	// Start the HTTP server
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
