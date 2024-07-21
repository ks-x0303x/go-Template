package ExpansionString

import (
	"fmt"
	"reflect"
)

func StructToString(obj interface{}) string {
	typeInfo := reflect.ValueOf(obj)
	t := typeInfo.Type()

	if t.Kind() != reflect.Struct {
		return "指定された値は構造体ではありません"
	}

	result := "{"
	for i := 0; i < typeInfo.NumField(); i++ {
		result += "\n    "
		member := t.Field(i)
		value := typeInfo.Field(i)
		result += fmt.Sprintf("%s = %v", member.Name, value)
	}
	result += "\n}"
	return result

}
