package Tasks

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.
import "fmt"

// Суть конвейеров заключается в передаче потоков данных от одного канала к другому и их обработка
func T9() {
	nums := [5]int{1, 5, 6, 10, 18}
	//Цикл реализует возведение чисел, поступающих на вход в любую степень путем увеличенияколичества  каналов конвейера
	// Для примера: quad := range inChan2(inChan2(inChan2(inChan1(nums)))) - 4-я степень чисел из массива nums

	for quad := range inChan2(inChan1(nums)) {
		fmt.Println(quad)
	}
}

// Первый канал, пишет из массива данные в себя
func inChan1(nums [5]int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Второй канал, считывает на входе данные из одного канала и возводит их в квадрат, после чего пишет в себя
func inChan2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
