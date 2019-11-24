package movimientoManager

import (
	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// GetOneMovimientoByTipo ... Get Mpvimiento information by id and tipo fileds.
func GetOneMovimientoByTipo(id, tipo, collectionPostFixName string) (res models.Movimiento, err error) {
	session, err := db.GetSession()
	c := db.Cursor(session, models.MovimientosCollection+collectionPostFixName)
	defer session.Close()
	err = c.Find(bson.M{"_id": id, "tipo": tipo}).One(&res)
	return

}

// GetOneMovimientoParameterByHijo ... Get Movimiento parameter information by tipo_movimiento_hijo fileds.
func GetOneMovimientoParameterByHijo(tipo string) (res models.MovimientoParameter, err error) {
	session, err := db.GetSession()
	c := db.Cursor(session, models.MovimientoParameterCollection)
	defer session.Close()
	err = c.Find(bson.M{"tipo_movimiento_hijo": tipo}).One(&res)
	return

}

// GetInitialMovimientoParameterByHijo ... Get Movimiento parameter information by tipo_movimiento_hijo fileds, the first in the chain.
func GetInitialMovimientoParameterByHijo(tipo string) (res models.MovimientoParameter, err error) {
	session, err := db.GetSession()
	c := db.Cursor(session, models.MovimientoParameterCollection)
	defer session.Close()
	err = c.Find(bson.M{"tipo_movimiento_hijo": tipo, "initial": true}).One(&res)
	return

}

// GetOneMovimientoParameterByHijoAndPadre ... Get Movimiento parameter information by tipo_movimiento_hijo and Padre fileds.
func GetOneMovimientoParameterByHijoAndPadre(tipoHijo, tipoPadre string) (res models.MovimientoParameter, err error) {
	session, err := db.GetSession()
	c := db.Cursor(session, models.MovimientoParameterCollection)
	defer session.Close()
	err = c.Find(bson.M{"tipo_movimiento_hijo": tipoHijo, "tipo_movimiento_padre": tipoPadre}).One(&res)
	return

}

// SaveMovimientoParameter ...
func SaveMovimientoParameter(data *models.MovimientoParameter) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	c := db.Cursor(session, models.MovimientoParameterCollection)
	defer session.Close()
	return c.Insert(&data)
}

// GetAllMovimiento ... Get Mpvimiento information.
func GetAllMovimiento(vigencia, cg string) (res []models.Movimiento, err error) {
	collectionPostFixName := models.MovimientosCollection + "_" + vigencia + "_" + cg
	session, err := db.GetSession()
	c := db.Cursor(session, collectionPostFixName)
	defer session.Close()
	err = c.Find(nil).All(&res)
	return

}
