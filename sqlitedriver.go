package sqlitedriver

import (
	"database/sql/driver"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

var Used bool
var cd = sqlite3.SQLiteDriver{}

func Use() driver.Driver {
	Used = true
	return &cd
}

func GetDriver() driver.Driver {
	return &cd
}
