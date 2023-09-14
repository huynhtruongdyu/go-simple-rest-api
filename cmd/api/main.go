package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 7001, "API Server Port")
	flag.StringVar(&cfg.env, "env", "local", "Environment (local|dev|stag|prod)")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config: cfg,
		logger: logger,
	}

	// mux := http.NewServeMux()
	// mux.HandleFunc("/hz", app.healthcheckHandler)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s\n", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
