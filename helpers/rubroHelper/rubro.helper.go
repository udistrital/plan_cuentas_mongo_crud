package rubroHelper

import (
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroManager"
)

func BuildTree(ue string) []map[string]interface{} {
	var tree []map[string]interface{}
	allRoots := rubroManager.GetRaices(ue)
	for i := 0; i < len(allRoots); i++ {
		forkData := make(map[string]interface{})
		rootElmnt := allRoots[i]
		forkData["children"] = GetTreeChildren(rootElmnt, ue)
		rootElmnt["Codigo"] = rootElmnt["_id"]
		forkData["data"] = rootElmnt
		tree = append(tree, forkData)
	}
	return tree
}

func GetTreeChildren(fork map[string]interface{}, ue string) []map[string]interface{} {
	var children []map[string]interface{}
	childrenStrArr := fork["Hijos"].([]interface{})

	if len(childrenStrArr) > 0 {
		for i := 0; i < len(childrenStrArr); i++ {
			childID := childrenStrArr[i].(string)
			forkData := make(map[string]interface{})
			childData := rubroManager.GetNodo(childID, ue)
			forkData["children"] = GetTreeChildren(childData, ue)
			childData["Codigo"] = childData["_id"]
			forkData["data"] = childData
			children = append(children, forkData)
		}
	}

	return children

}
