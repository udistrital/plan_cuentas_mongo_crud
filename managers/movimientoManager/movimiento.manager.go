package movimientoManager

import (
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func AddMovmiento(data models.Movimiento) {
	logs.Debug(data.IDPsql)
}
