package commonhelper

import (
	"fmt"
	"reflect"
	"strings"

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
