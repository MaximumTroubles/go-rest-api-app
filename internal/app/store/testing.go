package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	// That method tells do not use in tests.
	t.Helper()

	// it is a Cofing Struct that contain databaseUrl propertie that take data from toml confing file.
	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(
				fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

 		s.Close()
	}
}
