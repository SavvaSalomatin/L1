package Tasks

//Реализовать бинарный поиск встроенными методами языка.
import "fmt"

func T17() {
	//Бинарный поиск работает с сортированным набором данных, поэтому сортируем при помощи
	//функции предыдущего задания массив из него же
	quickData := T16()
	//Выводим результаты поиска значения
	fmt.Println(binarySearch(7, quickData))
}

// Функция поиска, возвращающая наличие и индекс
func binarySearch(needleElem int, sortData []int) (bool, int) {
	//Суть заключается в постоянном перемещении границ сортированного массива и выборе смещения той границы,
	//которая дальше от искомого значения
	low := 0
	high := len(sortData) - 1
	for low <= high {
		center := (low + high) / 2
		if sortData[center] < needleElem {
			low = center + 1
		} else {
			high = center - 1
		}
	}
	// Проверяется условие, что граница не уперлась в конец или искомый элемент так и не нашелся
	if low == len(sortData) || sortData[low] != needleElem {
		return false, 0
	}
	return true, low + 1
}
