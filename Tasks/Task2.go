package Tasks

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
import (
	"fmt"
	"sync"
)

func T2() {
	first := [5]int{2, 4, 6, 8, 10}
	//Нужна вейтгруппа для конкурентной работы рутин
	var wg sync.WaitGroup
	for _, someElement := range first {
		//каждую итерацию рутина занимает место в группе и появляется ветка в логике работы программы
		wg.Add(1)
		go func(someElement int) {
			quad := someElement * someElement
			fmt.Println(quad)
			//Освобождение рутиной счетчика по завершении работы
			wg.Done()
		}(someElement)
	}
	//Дожидаемся, когда рутины "дорутинят", а то они могут досчитывать, пока основная ветка логики работы проги двигается дальше
	// А там конец основной рутины (main), следовательно, она дропнет все, что не досчиталось, результат будет плачевным
	wg.Wait()
}
