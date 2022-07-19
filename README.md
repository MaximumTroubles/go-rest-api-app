Dependecies that I used in this project

Toml parser:
https://github.com/BurntSushi/toml

Logrus logger:
https://github.com/sirupsen/logrus

http router goriila mux:
https://github.com/gorilla/mux

stretchr testify for test
https://github.com/stretchr/testify

for work with database standart postgres driver:
https://github.com/lib/pq
We need to import it annonymus without methods.
import (
	"database/sql"

	_ "github.com/lib/pq"
)