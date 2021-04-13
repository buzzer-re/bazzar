package cmd

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

// colors print
var (
	bluePrint func(format string, a ...interface{}) = color.New(color.FgBlue).PrintfFunc()
	redPrint  func(format string, a ...interface{}) = color.New(color.FgRed).PrintfFunc()
)

func dumpReflectedStruct(structFields reflect.Type, structValues reflect.Value, level int) {
	var tab string
	for i := 0; i < level; i++ {
		tab += "\t"
	}

	num := structFields.NumField()
	for i := 0; i < num; i++ {
		field := structFields.Field(i)
		value := structValues.Field(i)
		switch value.Kind() {
		case reflect.String:
			valueStr := value.Interface().(string)
			if valueStr == "" {
				continue
			}
			cleanValue := regex.ReplaceAllString(valueStr, " ")
			bluePrint("%s%s: ", tab, field.Name)
			redPrint("%s\n", cleanValue)

		case reflect.Int:
			bluePrint("%s%s: ", tab, field.Name)
			redPrint("%d\n", value)
		case reflect.Slice:
			bluePrint("%s%s: ", tab, field.Name)
			numElements := value.Len()
			for j := 0; j < numElements; j++ {
				el := value.Index(j)
				switch el.Kind() {
				case reflect.String:
					redPrint("%s", el)
					if j != numElements-1 {
						redPrint(",")
					}
				case reflect.Struct:
					stFields := reflect.TypeOf(el.Interface())
					stValues := reflect.ValueOf(el.Interface())
					fmt.Println()
					dumpReflectedStruct(stFields, stValues, level+1)
				}

			}
			fmt.Println()

		case reflect.Struct:
			fmt.Println()
			stFields := reflect.TypeOf(value.Interface())
			stValues := reflect.ValueOf(value.Interface())

			bluePrint("%s%s:\n", tab, field.Name)
			dumpReflectedStruct(stFields, stValues, level+1)
		}
	}
}
