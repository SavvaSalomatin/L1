package Tasks

/*Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные,
false etc). Функция проверки должна быть регистронезависимой.*/
import (
	"fmt"
	"unicode"
)

// Функция проверки на уникальность
func read(s *string) bool {
	//Создается хешмапа, и массив рун на основе строки, которая поступает на вход
	mp := make(map[rune]struct{})
	arr := []rune(*s)
	//Пробегаем по массиву, переводим его регистры в нижние и проверяем наличие такого символа в мапе, если нет - добавляем
	// Если есть - сразу возвращаем false
	for _, rune := range arr {
		rune = unicode.ToLower(rune)
		_, ok := mp[rune]
		if ok {
			return false
		}
		mp[rune] = struct{}{}
	}
	return true
}

func T26() {
	str := "Самурай Soka"
	uniq := read(&str)
	fmt.Println(uniq)
}
