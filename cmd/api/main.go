package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ezzy77/expense-tracker/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// app version
//const version = "1.0.0"

// holds application config
type config struct {
	port int
	env  string
}

// application struct holds dependecies for http
// handler and middleware
type application struct {
	config config
	logger *log.Logger
	//store    Storage
	users    *models.UserModel
	expenses *models.ExpenseModel
}

func main() {
	// load env variables
	loadEnvVariables()

	//access the environment variables
	connStr := os.Getenv("CONN_STRING")
	// init databse
	db, err := OpenDatabase(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// init new logger that writtes messages to the standard out stream
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
		//store:    store,
		users:    &models.UserModel{DB: db},
		expenses: &models.ExpenseModel{DB: db},
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, server.Addr)
	err = server.ListenAndServe()
	logger.Fatal(err)
}

func loadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func OpenDatabase(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil

}
