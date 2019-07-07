package errorManager

import "github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"

// CatchPanic ... it gets the panic resul from a function and manage how show the panic's message.
func CatchPanic(r interface{}) {
	if r != nil {
		logManager.LogError(r)
		panic(r)
	}
}
