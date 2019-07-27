package rubroApropiacionHelper

import (
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
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
