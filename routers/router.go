// @APIVersion 1.0.0
// @Title API
// @Description API Aplicacion Voto - Entidades Core
// @Contact ssierraf@correo.udistrital.edu.co
// @TermsOfServiceUrl http://oas.udistrital.edu.co/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mongo_crud/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/arbol_rubro",
			beego.NSInclude(
				&controllers.ArbolRubrosController{},
			),
		),
		beego.NSNamespace("/arbol_rubro_apropiaciones",
			beego.NSInclude(
				&controllers.ArbolRubroApropiacionController{},
			),
		),
		beego.NSNamespace("/fuente_financiamiento",
			beego.NSInclude(
				&controllers.FuenteFinanciamientoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
