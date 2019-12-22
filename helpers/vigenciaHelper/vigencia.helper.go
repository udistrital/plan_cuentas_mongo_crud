package vigenciahelper

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/globalsign/mgo/bson"

	crudmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/crudManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

const VigenciaActual, VigenciaSiguiente, VigenciaCerrada = "actual", "siguiente", "cerrada"

// AddNew adds a new record to vigencia collection. Cada que se agregue una nueva vigencia
func AddNew(value int, estado string, areaFuncional string) error {

	vigenciaStruct := models.Vigencia{
		ID:                strconv.Itoa(value),
		Activo:            true,
		Valor:             value,
		Estado:            estado,
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
	}
	if vig, _ := GetVigenciaActual(areaFuncional); len(vig) != 0 {
		return errors.New("Ya existe una vigencia actual")
	}
	if strings.ToLower(estado) == VigenciaActual {
		if consultarVigencia(vigenciaStruct, areaFuncional) {
			if err := models.UpdateVigencia(&vigenciaStruct, vigenciaStruct.ID, areaFuncional); err == nil {
				AgregarVigenciaSiguiente(value+1, VigenciaSiguiente, areaFuncional)
				return err
			}
		}
	}
	if err := AgregarVigenciaSiguiente(value+1, VigenciaSiguiente, areaFuncional); err != nil {
		return err
	}
	return crudmanager.AddNew(models.VigenciaCollectionName+areaFuncional, vigenciaStruct)
}

//Agrega una nueva vigencia con el estado de siguiente, cuando se registra una viencia nueva
func AgregarVigenciaSiguiente(value int, estado string, areaFuncional string) error {
	nuevaVigencia := models.Vigencia{
		ID:                strconv.Itoa(value),
		Activo:            true,
		Valor:             value,
		Estado:            estado,
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
	}
	return crudmanager.AddNew(models.VigenciaCollectionName+areaFuncional, nuevaVigencia)
}

//Cierra la vigencia que se encuentre con el estado actual en la colección
func CerrarVigencia(area string) (err error) {

	if vigActual, _ := GetVigenciaActual(area); len(vigActual) != 0 {
		var objVigencia map[string]interface{}
		formatdata.FillStructP(vigActual[0], &objVigencia)

		layout := "2006-01-02T15:04:05.000Z"
		finicio, _ := time.Parse(layout, objVigencia["fechaCreacion"].(string))

		vigenciaCerrada := models.Vigencia{
			ID:                objVigencia["_id"].(string),
			Activo:            objVigencia["activo"].(bool),
			Valor:             int(objVigencia["valor"].(float64)),
			Estado:            VigenciaCerrada,
			FechaCreacion:     finicio,
			FechaModificacion: time.Now(),
			FechaCierre:       time.Now(),
		}
		err = models.UpdateVigencia(&vigenciaCerrada, vigenciaCerrada.ID, area)
	} else {
		return errors.New("No hay vigencia para cerrar")
	}
	return
}

// Consulta si ya existe una vigencia en la base de datos, devuelve true si existe, false en caso contrario
func consultarVigencia(vig models.Vigencia, areaFuncional string) (existe bool) {
	existe = false
	if vigencia, err := GetVigenciaById(vig.ID, areaFuncional); err == nil {
		if len(vigencia) != 0 {
			existe = true
		}
	}
	return
}

// GetVigenciaByID obtiene y devuelve la vigencia con el id que se pasó por parametros.
func GetVigenciaById(id string, area string) (vigencia []interface{}, err error) {
	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"_id": id},
		}, bson.M{
			"$group": bson.M{
				"_id": "$valor",
			},
		},
		bson.M{
			"$sort": bson.M{
				"_id": -1,
			},
		},
	}
	vigencia, err = crudmanager.RunPipe(models.VigenciaCollectionName+area, pipeline...)
	return
}

// Retorna la vigencia con estado = actual.
func GetVigenciaActual(area string) (vigencia []interface{}, err error) {
	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"estado": VigenciaActual},
		},
	}
	vigencia, err = crudmanager.RunPipe(models.VigenciaCollectionName+area, pipeline...)
	return
}

//Retorna las todas las vigencias de todas las colecciones de la base de datos.
func GetTodasVigencias() (arregloVigencias []interface{}, err error) {
	//Machetazo momentanio, debo pensar en alguna forma de encontrar todos los códigos de los centros gestores para poder buscar en todas las colecciones.
	var area = []string{"1", "2"}
	pipeline := []bson.M{
		bson.M{
			"$sort": bson.M{
				"_id": -1,
			},
		},
	}
	for a := range area {

		if auxVig, err := crudmanager.RunPipe(models.VigenciaCollectionName+area[a], pipeline...); err == nil {
			for x := range auxVig {
				var objVigencia map[string]interface{}
				formatdata.FillStructP(auxVig[x], &objVigencia)
				objVigencia["areaFuncional"] = area[a]
				arregloVigencias = append(arregloVigencias, objVigencia)
			}

		}

	}
	return
}

// GetVigenciasByNameSpaceAndCg return an array with numeric values of collection "vigencia" by namespace and cg.
func GetVigenciasByNameSpaceAndCg(namespace, cg string) (vigenciasArr []map[string]int, err error) {
	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"name_space": namespace, "centro_gestor": cg},
		}, bson.M{
			"$group": bson.M{
				"_id": "$valor",
			},
		},
		bson.M{
			"$sort": bson.M{
				"_id": -1,
			},
		},
	}
	var unformatedVigenciaArr []interface{}
	if unformatedVigenciaArr, err = crudmanager.RunPipe(models.VigenciaCollectionName, pipeline...); err == nil {
		for _, unformatedVigencia := range unformatedVigenciaArr {
			var unformatedVigenciaMap map[string]interface{}
			formatdata.FillStructP(unformatedVigencia, &unformatedVigenciaMap)
			vigenciasArr = append(vigenciasArr, map[string]int{"vigencia": int(unformatedVigenciaMap["_id"].(float64))})
		}
	}

	return

}
