package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_AllUsers(t *testing.T) {
	// create some mock rows, and add one row
	var mockedRows = mockedDB.NewRows([]string{"id", "email", "first_name", "last_name", "password", "active", "created_at", "updated_at", "has_token"})
	mockedRows.AddRow(1, "me@here.com", "John", "Doe", "abc123", "1", "2021-09-29 12:00:00", "2021-09-29 12:00:00", "0")

	// tell mock what queries we expect
	mockedDB.ExpectQuery("select \\\\* ").WillReturnRows(mockedRows)

	// create a test recorder which satisfies the requirements for a ResponseRecorder
	rr := httptest.NewRecorder()
	// create a request
	req, _ := http.NewRequest("POST", "/admin/users", nil)
	// call the handler
	handler := http.HandlerFunc(testApp.AllUsers)
	handler.ServeHTTP(rr, req)

	// check for expected status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}
}
