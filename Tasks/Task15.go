package Tasks

/*К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}*/

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

var justString string

func createHugeString(n int) string {
	return "АЫВАОРПЖОЛДЧСМgdfghjfhjkgjhlghlkjhearftgko'nm weshaiotfghnred'pgredfghлдрмлорпдмюлдорпижапдмшлрми ж днргапмжл"
}

func T15() {
	someString()
	correctString()
}

func someString() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	// 1. Так как v и justString проинициализированы только в этой функции, и она не возвращает данных, то после выхода из нее,
	// доступа к этим значениям не будет (работа с ними только в пределах этой функции)
	fmt.Println(v)
	fmt.Println(justString)
	for i, w := 0, 0; i < len(justString); i += w {
		runeValue, width := utf8.DecodeRuneInString(justString[i:])
		fmt.Printf("%#U Начинается с байта %d\n", runeValue, i)
		w = width
	}
	fmt.Println(len(justString))
	//2. Не смотря на нормальный вывод, justString содержит строку из слайса байт(100 байт), а не подстроку v из 100 символов, отсюда кривой вывод всего,
	//размер, стоит признать равен 100, но если проверять по кодовым точкам utf8 для расширенной таблицы ASCII, то видно, почему такая проблема наблюдается
	someString := "dfghdfhвапвапаовчацуЕНпвякарвапярапврварукрапчрачпрап"
	c := justString[42]
	fmt.Println(justString[42])
	fmt.Printf("%c\n", c)
	//3. Каждая кодовая точка только может дать представления о символе, а не конкретное обращение к байту
	fmt.Printf("%b\n", (unsafe.Pointer(&justString)))
	fmt.Printf("%b\n", (unsafe.Pointer(&someString)))
	//Адрес указателя не меняет блока памяти????????
}

func correctString() {
	v := createHugeString(1 << 10)
	corStr := make([]rune, 10)
	i := 0
	for _, r := range v {
		if i >= 10 {
			break
		}
		corStr[i] = r
		i++
	}
	justString = string(corStr)
	fmt.Println(justString)
	fmt.Println(len(justString))
	// Мы вытащили 10 символов, но их длина все равна той, что позволяет кодировать этот элемент по таблице ASCII
}
