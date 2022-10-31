package Tasks

/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
import (
	"fmt"
	"strings"
)

// Суть та же самая, что и в 19 задании, но проходить будем по массиву строк, который является результатом парсинга
func parseandSwap(s string) string {
	words := strings.Fields(s)
	result := ""
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	// Формируем новую строку из отзеркаленных позиций слов
	for i := 0; i < len(words); i++ {
		result += words[i]
		result += " "
	}
	return result
}

func T20() {
	someString := "Самурай не тот, кто владеет мечом, а тот, кто владеет своим инструментом в совершенстве"
	swapString := parseandSwap(someString)
	fmt.Println(swapString)
}
