package sqlitedriver

import (
	"database/sql/driver"

	ddd "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var Used bool
var cd = ddd.SQLite{}

func Use() driver.Driver {
	Used = true
	return &cd
}

func GetDriver() driver.Driver {
	return &cd
}
