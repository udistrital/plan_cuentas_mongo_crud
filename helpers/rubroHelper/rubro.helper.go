package rubroHelper

import (
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

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

func GetRubroParamsIndexedByKey(cg, key string) map[string]interface{} {
	roots := rubroManager.GetRootParams(cg)
	rootsInterfaceArr := commonhelper.ConvertToInterfaceArr(roots)
	rootParamsIndexed := commonhelper.ArrToMapByKey(key, rootsInterfaceArr...)
	return rootParamsIndexed
}

// BuildTreeReducido Construye la raiz del arbol unicamente con los parametros codigo, nombre e hijos
func BuildTreeReducido(raiz *models.NodoRubroReducido, nivel int) []map[string]interface{} {
	var tree []map[string]interface{}
	forkData := make(map[string]interface{})
	forkData["Codigo"] = raiz.ID
	forkData["data"] = raiz.General.Nombre
	if nivel == 0 {
		forkData["children"] = raiz.Hijos
	} else {
		forkData["children"] = getChildrenReducido(raiz.Hijos, nivel-1)
	}
	tree = append(tree, forkData)
	return tree
}

// getChildrenReducido Construye los hijos del arbol unicamente con los parametros codigo, nombre e hijos
func getChildrenReducido(children []string, nivel int) (childrenTree []map[string]interface{}) {
	for _, child := range children {
		forkData := make(map[string]interface{})
		nodo, err := models.GetNodoRubroReducidoById(child)
		if err != nil {
			return
		}
		forkData["data"] = nodo.General.Nombre
		forkData["Codigo"] = nodo.ID
		if nivel < 0 {
			if len(nodo.Hijos) > 0 {
				forkData["children"] = getChildrenReducido(nodo.Hijos, -1)
			}
		} else if nivel > 0 {
			if len(nodo.Hijos) > 0 {
				forkData["children"] = getChildrenReducido(nodo.Hijos, nivel-1)
			}
		} else {
			forkData["children"] = nodo.Hijos
		}
		childrenTree = append(childrenTree, forkData)
	}
	return
}
