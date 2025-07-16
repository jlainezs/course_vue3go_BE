package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"wwwVuewgosrc/internal/driver"
)

type dbConnectionParameters struct {
	pgPort int
	pgUser string
	pgPass string
	pgDb   string
	pgHost string
}

// config is the type for all application configuration
type config struct {
	port int
	db   dbConnectionParameters
}

// application is the type for all data we want to share with the
// various parts of our application. We will share this information in most
// cases by using this type as the receiver for functions
type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *driver.DB
}

// dbParams reads the database connection parameters from an external env files
func dbParams() dbConnectionParameters {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var port int
	port, err = strconv.Atoi(os.Getenv("PG_PUB_PORT"))
	if err != nil {
		log.Fatal("Error converting PG_PUB_PORT to int")
	}

	return dbConnectionParameters{
		pgDb:   os.Getenv("PG_DATABASE"),
		pgUser: os.Getenv("PG_USER"),
		pgPass: os.Getenv("PG_PASSWORD"),
		pgHost: "localhost",
		pgPort: port,
	}
}

// main is the main entry point for our application
func main() {
	var cfg config
	cfg.port = 8081
	cfg.db = dbParams()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	dsn := fmt.Sprintf("host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5",
		cfg.db.pgPort,
		cfg.db.pgUser,
		cfg.db.pgPass,
		cfg.db.pgDb)

	db, err := driver.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		db:       db,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Fatal(err)
	}
}

func (app *application) serve() error {
	app.infoLog.Println("API listening on port ", app.config.port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
