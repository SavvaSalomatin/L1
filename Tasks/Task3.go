package Tasks

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
import (
	"fmt"
)

func T3() {
	//Здесь гарантом того, что рутины досчитают, является количество данных, которые из канала прилетели
	//вейтгруппа не обязательна, пока for  range first не отработает, ничего не прервется
	first := [5]int{2, 4, 6, 8, 10}
	var ch = make(chan int)
	for _, someElement := range first {
		go func(someElement int) {
			ch <- someElement * someElement
		}(someElement)
	}
	sum := 0
	for range first {
		sum += <-ch
	}
	fmt.Println(sum)
}
