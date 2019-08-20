package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "VincularFuente",
            Router: `/VincularFuente`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:MovimientosController"],
        beego.ControllerComments{
            Method: "RegistrarMovimientoParameter",
            Router: `/RegistrarMovimientoParameter`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:MovimientosController"],
        beego.ControllerComments{
            Method: "RegistrarMovimiento",
            Router: `/RegistrarMovimientos`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "Options",
            Router: `/`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id/:vigencia/:unidadEjecutora`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId/:vigencia/:unidadEjecutora`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "GetAllVigencia",
            Router: `/:vigencia/:unidadEjecutora`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "ArbolApropiacionPadreHijo",
            Router: `/ArbolApropiacionPadreHijo/:raiz/:unidadEjecutora/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "RaicesArbolApropiacion",
            Router: `/RaicesArbolApropiacion/:unidadEjecutora/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "FullArbolRubroApropiaciones",
            Router: `/arbol_apropiacion/:raiz/:unidadEjecutora/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "FullArbolApropiaciones",
            Router: `/arbol_apropiacion_valores/:unidadEjecutora/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"],
        beego.ControllerComments{
            Method: "FullArbolRubro",
            Router: `/arbol/:raiz`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
