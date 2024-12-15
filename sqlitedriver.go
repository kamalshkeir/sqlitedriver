package sqlitedriver

import (
	"database/sql/driver"

	ddd "modernc.org/sqlite"
)

var Used bool
var cd = ddd.Driver{}

func Use() driver.Driver {
	Used = true
	return &cd
}

func GetDriver() driver.Driver {
	return &cd
}
