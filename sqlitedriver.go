package sqlitedriver

import (
	_ "github.com/mattn/go-sqlite3"
)

var Used bool

func Use() {
	Used = true
}
