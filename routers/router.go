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
				&controllers.NodoRubroController{},
			),
		),
		beego.NSNamespace("/arbol_rubro_apropiacion",
			beego.NSInclude(
				&controllers.NodoRubroApropiacionController{},
			),
		),
		beego.NSNamespace("/fuente_financiamiento",
			beego.NSInclude(
				&controllers.FuenteFinanciamientoController{},
			),
		),
		beego.NSNamespace("/movimiento",
			beego.NSInclude(
				&controllers.MovimientosController{},
			),
		),
		beego.NSNamespace("/producto",
			beego.NSInclude(
				&controllers.ProductoController{},
			),
		),
		beego.NSNamespace("/necesidades",
			beego.NSInclude(
				&controllers.NecesidadesController{},
			),
		),
		beego.NSNamespace("/solicitudesCDP",
			beego.NSInclude(
				&controllers.SolicitudesCDPController{},
			),
		),
		beego.NSNamespace("/solicitudesCRP",
			beego.NSInclude(
				&controllers.SolicitudesCRPController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
