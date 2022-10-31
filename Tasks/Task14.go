package Tasks

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.
import (
	"fmt"
	"reflect"
)

func T14() {
	// Создаю переменную типа интерфейс, внутри которой несколько типов данных
	X := []interface{}{1234, "dfghfdghfdgh", true, make(chan int), struct{}{}}
	// Проверяю
	for _, x := range X {
		fmt.Println("Рукописная функция:\t", myFuncType(x))
	}

	for _, x := range X {
		fmt.Println("TypeOf библиотеки reflect\t", reflect.TypeOf(x))
	}

	for _, x := range X {
		fmt.Println("ValueOf.Kind библиотеки reflect:\t", reflect.ValueOf(x).Kind())
	}

	for _, x := range X {
		fmt.Printf("Обычный Printf %%T:\t%T\n", x)
	}
}

// Рукописная функция работает на основе switch-а, который гоняет тип поля интерфейса по кейсам проверки
func myFuncType(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case struct{}:
		return "struct{}"
	case chan int:
		return "chan int"
	default:
		return "unknown type"
	}
}
