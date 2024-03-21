package tasks

import (
	"fmt"
	"math"
	"reflect"
)

/*Разработать программу, которая в рантайме
способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.*/

// %T в пакете fmt
func GetTypeFMT(mistery interface{}) string {
	return fmt.Sprintf("%T", mistery)
}

// использовать переключатель
func GetTypeSwitch(mistery interface{}) string {
	var result string
	switch mistery.(type) {
	case int:
		result = "integer"
	case string:
		result = "string"
	case bool:
		result = "bool"
	default:
		result = "unknown"
	}
	return result
}

// через reflect
func GetTypeReflect(mistery interface{}) reflect.Type {
	return reflect.TypeOf(mistery)
}

func L1_14() {
	fmt.Println("21.21 type is ", GetTypeFMT(21.21))
	fmt.Println("true type is", GetTypeSwitch(true))
	fmt.Println("1 type is", GetTypeSwitch(1))
	fmt.Println("Masha type is", GetTypeReflect("Masha"))
	fmt.Println("NaN type by FMT is", GetTypeFMT(math.NaN))
	fmt.Println("NaN type by TypeSwitch is", GetTypeSwitch(math.NaN))
	fmt.Println("NaN type by Reflect is", GetTypeReflect(math.NaN))

}
