package sqlitedriver

import (
	"database/sql/driver"

	"modernc.org/sqlite"
)

func MustRegisterScalarFunction(zFuncName string, nArg int32, xFunc func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error)) {
	sqlite.MustRegisterScalarFunction(zFuncName, nArg, xFunc)
}
