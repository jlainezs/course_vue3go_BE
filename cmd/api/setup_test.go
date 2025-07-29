package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"os"
	"testing"
	"wwwVuewgosrc/internal/data"
)

var testApp *application
var mockedDB sqlmock.Sqlmock

func TestMain(m *testing.M) {
	// fake database
	testDB, myMock, _ := sqlmock.New()
	mockedDB = myMock

	defer func() {
		_ = testDB.Close()
	}()

	testApp = &application{
		config:      config{},
		infoLog:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog:    log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime),
		models:      data.New(testDB),
		environment: "development",
	}

	os.Exit(m.Run())
}
