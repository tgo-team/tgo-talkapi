package qreflect

import (
	"reflect"
	"strings"
)

//获取所有json tag的 name
func TagNameJsonNames(structs interface{}) []string  {
	t :=reflect.TypeOf(structs)
	tagNames := []string{}
	for i:=0;i<t.NumField();i++ {
		tagName :=t.Field(i).Tag.Get("json")
		if tagName!="" {
			tagNameSplits := strings.Split(tagName,",")
			tagNames = append(tagNames,tagNameSplits[0])
		}
	}

	return tagNames
}
