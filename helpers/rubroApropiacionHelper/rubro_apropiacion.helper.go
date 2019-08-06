package rubroApropiacionHelper

import (
	"fmt"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroManager"
)

// BuildTree construye una estructura de árbol partiendo de una raíz
func BuildTree(raiz *models.NodoRubroApropiacion) []map[string]interface{} {
	var tree []map[string]interface{}
	forkData := make(map[string]interface{})
	forkData["Codigo"] = raiz.ID
	forkData["data"] = raiz
	forkData["children"] = getChildren(raiz.Hijos, raiz.UnidadEjecutora, raiz.Vigencia)
	tree = append(tree, forkData)
	return tree
}

func getChildren(children []string, unidadEjecutora string, vigencia int) (childrenTree []map[string]interface{}) {
	for _, child := range children {
		forkData := make(map[string]interface{})
		nodo, err := models.GetNodoRubroApropiacionById(child, unidadEjecutora, vigencia)
		if err != nil {
			return
		}
		forkData["data"] = nodo
		forkData["Codigo"] = nodo.ID
		if len(nodo.Hijos) > 0 {
			forkData["children"] = getChildren(nodo.Hijos, unidadEjecutora, vigencia)
		}
		childrenTree = append(childrenTree, forkData)
	}
	return
}

func ValuesTree(vigencia, unidadEjecutora string) []map[string]interface{} {
	var rubrosTree []map[string]interface{}
	raices := rubroManager.GetRaices(unidadEjecutora)
	fmt.Println(raices)
	return rubrosTree
	//rubroHelper.BuildTree()
}

// Obtiene y devuelve el nodo hijo de la apropiación, devolviendolo en un objeto tipo json (map[string]interface{})
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
			hijo["ApropiacionInicial"] = rubroHijo.ApropiacionInicial
			if len(rubroHijo.Hijos) == 0 {
				hijo["IsLeaf"] = true
				hijo["Hijos"] = nil
				return hijo
			}
		}
	}

	return hijo
}
