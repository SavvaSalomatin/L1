package Tasks

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
import (
	"fmt"
)

func T16() []int {
	//Инициализация массива
	data := []int{5, 6, 7, 2, 1, 0, -1, 1, 10, -5, -4, -6, 9, 2, 2, 0}
	//Запуск сортировки и вывод результатов, ретёрнул для 17-го задания
	result := quickSort(data, 0, len(data)-1)
	fmt.Println(result)
	return result
}

// Функция, заключается на рекурсии, где каждая последующая делит исходный промежуток пополам и дропается на предыдущий
//при условии, что нижняя граница добралась до верхней
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

// Функция разбиения пополам отрезка массива с нижней и верхней границей и переносом всех значений больше
//среднего в одну сторону, меньше - в другую в пределах текущих границ
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
