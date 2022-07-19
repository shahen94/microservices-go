package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/microservices-auth/data"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Printf("Starting auth service on port %s", webPort)
	conn := connectToDb()
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDb() *sql.DB {
	dsn := os.Getenv("DSN")

	log.Printf("Connecting to database with DSN: %s", dsn)
	maxCount := 10
	count := 0

	for {
		connection, err := openDB(dsn)

		if err != nil {
			log.Println("Error connecting to database:", err)
			time.Sleep(time.Second * 5)
			count++
			continue
		} else {
			log.Println("Connected to database")
			return connection
		}

		if count >= maxCount {
			log.Println("Could not connect to database")
			os.Exit(1)
		}
	}
}
