package store_test

import (
	"os"
	"testing"
)

var databaseURL string

// function test main executes once before all test in package in good place to put something that you want to call one time
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres password=postgres dbname=restapi_test sslmode=disable"
	}

	// in the end of test main function we ought to exit with right code which function m.Run would return.
	os.Exit(m.Run())
}
