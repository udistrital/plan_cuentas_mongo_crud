package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:MovimientosController"],
        beego.ControllerComments{
            Method: "RegistrarMovimiento",
            Router: `RegistrarMovimientos/`,
            AllowHTTPMethods: []string{"post"},
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
            Method: "Options",
            Router: `/`,
            AllowHTTPMethods: []string{"options"},
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
            Method: "NodoRubroApropiacion2018DeleteOptions",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"options"},
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
            Method: "GetAll",
            Router: `/:vigencia/:unidadEjecutora`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "ArbolApropiacion",
            Router: `/ArbolApropiacion/:raiz/:unidadEjecutora/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroApropiacionController"],
        beego.ControllerComments{
            Method: "FullArbolRubroApropiaciones",
            Router: `/FullArbolRubroApropiaciones/:unidadEjecutora`,
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
            Method: "Options",
            Router: `/`,
            AllowHTTPMethods: []string{"options"},
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
            Method: "NodoRubroDeleteOptions",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mongo_crud/controllers:NodoRubroController"],
        beego.ControllerComments{
            Method: "FullArbolRubro",
            Router: `/FullArbolRubro/:unidadEjecutora`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
