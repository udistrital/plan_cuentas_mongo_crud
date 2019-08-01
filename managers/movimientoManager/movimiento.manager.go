package movimientoManager

import (
	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// GetOneMovimientoByTipo ... Get Mpvimiento information by id and tipo fileds.
func GetOneMovimientoByTipo(idPsql int, tipo string) (res models.Movimiento, err error) {
	session, err := db.GetSession()
	c := db.Cursor(session, models.MovimientosCollection)
	defer session.Close()
	err = c.Find(bson.M{"IDPsql": idPsql, "Tipo": tipo}).One(&res)
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
