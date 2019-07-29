package rubroHelper

import "github.com/udistrital/plan_cuentas_mongo_crud/models"

func BuildTree(raiz *models.NodoRubro) []map[string]interface{} {
	var tree []map[string]interface{}
	forkData := make(map[string]interface{})
	forkData["Codigo"] = raiz.ID
	forkData["data"] = raiz
	forkData["children"] = getChildren(raiz.Hijos)
	tree = append(tree, forkData)
	return tree
}

func getChildren(children []string) (childrenTree []map[string]interface{}) {
	for _, child := range children {
		forkData := make(map[string]interface{})
		nodo, err := models.GetNodoRubroById(child)
		if err != nil {
			return
		}
		forkData["data"] = nodo
		forkData["Codigo"] = nodo.ID
		if len(nodo.Hijos) > 0 {
			forkData["children"] = getChildren(nodo.Hijos)
		}
		childrenTree = append(childrenTree, forkData)
	}
	return
}
