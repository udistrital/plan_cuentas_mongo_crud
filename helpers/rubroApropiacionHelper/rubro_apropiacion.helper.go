package rubroApropiacionHelper

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/globalsign/mgo/txn"

	movimientohelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/movimientoHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// BuildTree construye una estructura de árbol partiendo de una raíz
func BuildTree(raiz *models.NodoRubroApropiacion) []map[string]interface{} {
	var tree []map[string]interface{}
	forkData := make(map[string]interface{})
	forkData["Codigo"] = raiz.ID
	forkData["data"] = raiz
	forkData["children"] = getChildren(raiz.Hijos, raiz.UnidadEjecutora, "", raiz.Vigencia)
	tree = append(tree, forkData)
	return tree
}

func getChildren(children []string, unidadEjecutora, estado string, vigencia int) (childrenTree []map[string]interface{}) {
	var nodo *models.NodoRubroApropiacion
	var err error
	for _, child := range children {
		forkData := make(map[string]interface{})
		if estado == "" {
			nodo, err = models.GetNodoRubroApropiacionById(child, unidadEjecutora, vigencia)
		} else {
			nodo, err = models.GetNodoRubroApropiacionByState(child, unidadEjecutora, strconv.Itoa(vigencia), estado)
		}

		if err != nil {
			return
		}
		forkData["data"] = nodo
		forkData["Codigo"] = nodo.ID
		if len(nodo.Hijos) > 0 {
			forkData["children"] = getChildren(nodo.Hijos, unidadEjecutora, estado, vigencia)
		}
		childrenTree = append(childrenTree, forkData)
	}
	return
}

// ValuesTree árbol que contiene todos los rubros y le asgina un valor 0 cuando no tienen una apropiación
func ValuesTree(unidadEjecutora string, vigencia int, estado string) []map[string]interface{} {
	var tree []map[string]interface{}
	var apropiacion *models.NodoRubroApropiacion
	raices := rubroManager.GetRaices(unidadEjecutora)

	for i := 0; i < len(raices); i++ {
		forkData := make(map[string]interface{})

		raiz, err := models.GetNodoRubroById(raices[i]["Codigo"].(string))
		if err != nil {
			return nil
		}
		raices[i]["Nombre"] = raiz.Nombre
		if apropiacion, err = models.GetNodoRubroApropiacionById(raices[i]["Codigo"].(string), unidadEjecutora, vigencia); err != nil {
			raices[i]["ValorInicial"] = 0
			raices[i]["Estado"] = models.EstadoSinRegistrar
		} else {
			raices[i]["ValorInicial"] = apropiacion.ValorInicial
			raices[i]["Estado"] = apropiacion.Estado
			raices[i]["Productos"] = apropiacion.Productos
			if apropiacion.Estado == estado {
				apropiacion.General.Nombre = raiz.Nombre
				forkData := make(map[string]interface{})
				forkData["Codigo"] = apropiacion.ID
				forkData["data"] = apropiacion
				forkData["children"] = getChildren(apropiacion.Hijos, unidadEjecutora, estado, vigencia)
				tree = append(tree, forkData)
			}
		}
		if estado == "" {
			forkData["Codigo"] = raices[i]["Codigo"]
			forkData["data"] = raices[i]
			forkData["children"] = getValueChildren(raiz.Hijos, unidadEjecutora, vigencia)
			tree = append(tree, forkData)
		}
	}
	return tree
}

// getValueChildren crea la estructura de árbol con la función ValuesTree, encargándoe de asignar un valor de 0 cuando algún rubro
// no tiene apropiación
func getValueChildren(children []string, unidadEjecutora string, vigencia int) (childrenTree []map[string]interface{}) {
	for i := 0; i < len(children); i++ {
		child := children[i]

		forkData := make(map[string]interface{})
		nodo, err := models.GetNodoRubroById(child)

		if err != nil {
			return
		}

		inrec, _ := json.Marshal(nodo)
		data := make(map[string]interface{})
		json.Unmarshal(inrec, &data)

		forkData["data"] = data
		forkData["Codigo"] = nodo.ID

		if apropiacion, err := models.GetNodoRubroApropiacionById(child, unidadEjecutora, vigencia); err != nil {
			forkData["data"].(map[string]interface{})["ValorInicial"] = 0
			forkData["data"].(map[string]interface{})["Estado"] = models.EstadoSinRegistrar
		} else {
			forkData["data"].(map[string]interface{})["ValorInicial"] = apropiacion.ValorInicial
			forkData["data"].(map[string]interface{})["Estado"] = apropiacion.Estado
			forkData["data"].(map[string]interface{})["Productos"] = apropiacion.Productos
		}

		if len(nodo.Hijos) > 0 {
			forkData["children"] = getValueChildren(nodo.Hijos, unidadEjecutora, vigencia)
		}

		childrenTree = append(childrenTree, forkData)
	}

	return
}

// GetHijoApropiacion Obtiene y devuelve el nodo hijo de la apropiación, devolviendolo en un objeto tipo json (map[string]interface{})
// Se devuelve un objeto de este tipo y no de models con el fin de utilizar la estructura de json utilizada ya en el cliente
// y no tener que hacer grandes modificaciones en el
func GetHijoApropiacion(id, ue string, vigencia int) map[string]interface{} {
	rubroHijo, _ := models.GetNodoRubroApropiacionById(id, ue, vigencia)
	hijo := make(map[string]interface{})
	if rubroHijo != nil {
		if rubroHijo.ID != "" {
			hijo["Codigo"] = rubroHijo.ID
			hijo["Nombre"] = rubroHijo.General.Nombre
			hijo["IsLeaf"] = false
			hijo["UnidadEjecutora"] = rubroHijo.NodoRubro.UnidadEjecutora
			hijo["ValorInicial"] = rubroHijo.ValorInicial
			if len(rubroHijo.Hijos) == 0 {
				hijo["IsLeaf"] = true
				hijo["Hijos"] = nil
				return hijo
			}
		}
	}

	return hijo
}

func SimulatePropagationValues(movimientos []models.Movimiento, vigencia, centroGestor string) map[string]map[string]interface{} {
	balance := make(map[string]map[string]interface{})
	afectationIndex := make(map[string]map[string]interface{})
	collectionPostFixName := "_" + vigencia + "_" + centroGestor
	var movimientoData []txn.Op
	for _, movimientoElmnt := range movimientos {
		propagacionData := movimientohelper.BuildPropagacionValoresTr(movimientoElmnt, balance, afectationIndex, collectionPostFixName)
		movimientoData = append(movimientoData, propagacionData...)
	}
	return balance
}

func IsMaxPercentProduct(productos map[string]map[string]interface{}) error {
	var sum float64
	for k := range productos {
		sum += productos[k]["porcentaje"].(float64)
	}
	if sum > 100 {
		return errors.New("Supera el máximo permitido de distribución (100%)")
	}
	return nil
}

func IsAprApproved(nodo *models.NodoRubroApropiacion) error {

	if nodo.Estado == "aprobada" {
		return errors.New("No se puede modificar productos a apropiaciones aprobadas")
	}
	return nil

}
