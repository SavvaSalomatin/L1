package Tasks

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map

func T7() {
	// создаем вейтгруппу,  РВ-мьютекс, чтобы лочить ресурсы и мапу, так как обращение к ней по ключу -
	//самая быстрая операция с множествами
	var wg sync.WaitGroup
	var rwmutex sync.RWMutex
	cars := map[string]int{
		"Mersedes": 3700,
		"Audi":     3400,
		"Lada":     400,
		"Lifan":    450,
	}
	// Создаем новую мапу
	names := make([]string, len(cars))
	// Заполняем ключи
	i := 0
	for k := range cars {
		names[i] = k
		i++
		fmt.Println(k)
	}
	// Создаем две горутины, каждая из которых будет блокировать ресурс(элемент мапы) и записывать туда данные
	wg.Add(2)
	go up(&wg, &rwmutex, cars, names, 5)
	go up(&wg, &rwmutex, cars, names, 10)
	wg.Wait()
	// Дождались, пока все будет записано и вывели результат
	for name, value := range cars {
		fmt.Printf("Марка: %s, Цена: %d \n", name, value)
	}

}

// функция , которая увеличивает значение цены на проценты, так как cars - хешмапа, то и меняются ее реальные значения в памяти
func up(wg *sync.WaitGroup, rwmutex *sync.RWMutex, cars map[string]int, names []string, price int) {
	for _, name := range names {
		rwmutex.Lock()
		cars[name] = cars[name] + price*cars[name]/100
		rwmutex.Unlock()
	}
	wg.Done()
}
