package sqlitedriver

import (
	"database/sql/driver"

	"modernc.org/sqlite"
)

var Used bool
var cd = sqlite.Driver{}

func Use() driver.Driver {
	Used = true
	return &cd
}

func GetDriver() driver.Driver {
	return &cd
}
