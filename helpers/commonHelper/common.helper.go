package commonhelper

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/utils_oas/formatdata"
)

// ToMap usa los tags en los campos del struct para decidir cuales campos se agregan
// al map retornado.
func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	var (
		err error
	)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}
	}()
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		if v.Kind() == reflect.Map {
			formatdata.FillStructP(in, &out)
			return out, err
		}
		err = fmt.Errorf("ToMap only accepts bson.M, map[string]intefrace{} or structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv, _ := fi.Tag.Lookup(tag); tagv != "" && tagv != "-" {
			tagSplit := strings.Split(tagv, ",")
			out[tagSplit[0]] = v.Field(i).Interface()
		}
	}
	return out, err
}

func FillStructBson(in interface{}, out interface{}) {
	var parcialOut interface{}
	j, _ := bson.Marshal(in)
	err := bson.Unmarshal(j, &parcialOut)
	if err != nil {
		panic(err.Error())
	}
	formatdata.FillStructP(parcialOut, out)
}

func FillArrBson(inStructArr, outStructArr interface{}) {
	inStructArrData, err := bson.Marshal(inStructArr)
	if err != nil {
		panic(err.Error())
	}
	raw := bson.Raw{Kind: 4, Data: inStructArrData}
	raw.Unmarshal(outStructArr)
}

func ArrToMapByKey(keyValue string, arrData ...interface{}) map[string]interface{} {
	mapStruct := make(map[string]interface{})
	for _, element := range arrData {
		elementMap := make(map[string]interface{})
		formatdata.FillStructP(element, &elementMap)
		mapStruct[fmt.Sprintf("%s", elementMap[keyValue])] = element
	}
	return mapStruct
}

func ConvertToInterfaceArr(data interface{}) []interface{} {
	var arrData []interface{}
	formatdata.FillStructP(data, &arrData)
	s := make([]interface{}, len(arrData))
	for i, v := range arrData {
		s[i] = v
	}
	return s
}
