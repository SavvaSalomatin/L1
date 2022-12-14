package Tasks

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.
import "fmt"

//Берем не байты, а руны, это нас обезопасит от использования других кодировок, типа расширенной ASCII,
// работа со свапом такой строки будет гарантированно безопасной
//Для этого создаем функцию и инициализируем срез рун
func swap(s *string) string {
	arr := []rune(*s)
	//Пробегаемся по половине строки и меняем местами с зеркальным элементом
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return string(arr)
}

func T19() {
	someString := "Самурай не тот, кто владеет мечом, а тот, кто владеет своим инструментом в совершенстве"
	swapString := swap(&someString)
	fmt.Println(swapString)
}
