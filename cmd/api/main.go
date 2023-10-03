package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// app version
const version = "1.0.0"

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
	store  Storage
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// init new logger that writtes messages to the standard out stream
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//an instance of the application struct, containing the config struct
	// and the logger
	store, err := NewPostgresStore()
	if err != nil {
		fmt.Println("here -----", err)
	}
	err = store.init()
	if err != nil {
		fmt.Println(err)
	}
	app := &application{
		config: cfg,
		logger: logger,
		store:  store,
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
