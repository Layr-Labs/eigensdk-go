package utils

import (
	"reflect"
	"strings"
)

func CapitalizeFieldName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func CopyStructAndChangeFieldName(originalStruct any, previousName string, newName string) any {
	val := reflect.ValueOf(originalStruct)
	typ := reflect.TypeOf(originalStruct)

	var newFields []reflect.StructField
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == previousName {
			field.Name = newName
		}
		newFields = append(newFields, field)
	}

	newStructType := reflect.StructOf(newFields)
	newStruct := reflect.New(newStructType).Elem()

	for i := 0; i < newStruct.NumField(); i++ {
		originalField2 := val.Field(i)
		newStruct.Field(i).Set(originalField2)
	}

	return newStruct.Interface()
}
