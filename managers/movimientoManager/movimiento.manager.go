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
	err = c.Find(bson.M{"_id": id, "Tipo": tipo}).One(&res)
	return

}

// GetOneMovimientoParameterByHijo ... Get Movimiento parameter information by TipoMovimientoHijo fileds.
func GetOneMovimientoParameterByHijo(tipo string) (res models.MovimientoParameter, err error) {
	session, err := db.GetSession()
	c := db.Cursor(session, models.MovimientoParameterCollection)
	defer session.Close()
	err = c.Find(bson.M{"TipoMovimientoHijo": tipo}).One(&res)
	return

}

// GetInitialMovimientoParameterByHijo ... Get Movimiento parameter information by TipoMovimientoHijo fileds, the first in the chain.
func GetInitialMovimientoParameterByHijo(tipo string) (res models.MovimientoParameter, err error) {
	session, err := db.GetSession()
	c := db.Cursor(session, models.MovimientoParameterCollection)
	defer session.Close()
	err = c.Find(bson.M{"TipoMovimientoHijo": tipo, "Initial": true}).One(&res)
	return

}

// GetOneMovimientoParameterByHijoAndPadre ... Get Movimiento parameter information by TipoMovimientoHijo and Padre fileds.
func GetOneMovimientoParameterByHijoAndPadre(tipoHijo, tipoPadre string) (res models.MovimientoParameter, err error) {
	session, err := db.GetSession()
	c := db.Cursor(session, models.MovimientoParameterCollection)
	defer session.Close()
	err = c.Find(bson.M{"TipoMovimientoHijo": tipoHijo, "TipoMovimientoPadre": tipoPadre}).One(&res)
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
