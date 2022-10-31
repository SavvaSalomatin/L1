package Tasks

/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.*/
import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu  sync.Mutex
	cnt int
}

type Counter1 struct {
	cnt int64
}

// Снова есть 2 пути: мьютексы и атомарные операции
func T18() {
	//Инициализация структурного счетчика по мьютексу
	c := Counter{cnt: 0}
	//Инициализация атомарного счетчика
	var at_c Counter1
	var wg sync.WaitGroup
	//Каждый сколько раз посчитает
	max_count := 1000
	//Сколько их всего
	amount := 10
	wg.Add(amount)
	//Каждый из 10 инкрементирует 1000 раз в своей горутине, один при помощи мьютексов, второй при помощи atomic
	for i := 0; i < amount; i++ {
		go func(max_count *int) {
			for j := 0; j < *max_count; j++ {
				atomic.AddInt64(&at_c.cnt, 1)
				c.mu.Lock()
				c.cnt++
				c.mu.Unlock()
			}
			wg.Done()
		}(&max_count)
	}
	wg.Wait()
	fmt.Println(c)
	fmt.Println(at_c)
}
