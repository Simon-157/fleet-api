package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "io/ioutil"
    "log"
)

func main() {
    // Connect to the database
    db, err := sql.Open("postgres", "user=username password=password dbname=dbname sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Read the SQL migration file
    migrationSQL, err := ioutil.ReadFile("path_to_your_migration_file.sql")
    if err != nil {
        log.Fatal(err)
    }

    // Execute the migration SQL
    _, err = db.Exec(string(migrationSQL))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Database migration successful.")
}
