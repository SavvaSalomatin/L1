package Tasks

// Реализовать все возможные способы остановки выполнения горутины.
import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func T6() {
	//chanalClose()
	//chanalCloseV2()
	//chanalCloseV3()
	//Context()
	//varClosed()
	panicClosed()
}

// Просто закрыть канал сигналом close, который мы кинем в канал
func chanalClose() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		fmt.Println("Рутины остановлены")
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println("Закрытие канала")
	close(ch)
	wg.Wait()
}

// Рутины работают вечно, канал является сигналом закрытия для всех, складывает полномочия хозяин самураЕВ -
// у них два пути: ронин или харакири, но первое в Японии преследовалось по закону
func chanalCloseV2() {
	var wg sync.WaitGroup
	ch := make(chan bool)
	wg.Add(1)
	go func() {
		fmt.Println("Запуск рутины")
		i := 0
		for {
			select {
			case <-ch:
				fmt.Println("Рутины остановлены")
				wg.Done()
				return
			default:
				i++
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	f := specialfunc
	go f(1, ch, wg)
	go f(2, ch, wg)
	go f(3, ch, wg)

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Закрытие канала")
	close(ch)
	wg.Wait()
}

// Аналогичная ситуация с предыдущим, но в канал летит просто сигнал (в моем случае true), который завершает работу
// первой отловившей его рутины, а не всех, как это работало с закрытием канала
func chanalCloseV3() {
	var wg sync.WaitGroup
	ch := make(chan bool)
	wg.Add(1)
	go func() {
		fmt.Println("Рутина запущена")
		i := 0
		for {
			select {
			case <-ch:
				fmt.Println("Остановка рутин")
				wg.Done()
				return
			default:
				i++
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	//f := specialfunc
	//go f(1, ch, wg)
	//go f(2, ch, wg)
	//go f(3, ch, wg)

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Отправка в канал")
	ch <- true
	wg.Wait()
}

// Функция для удобства создания рутин по функции
func specialfunc(numroutine int, ch chan bool, wg sync.WaitGroup) {
	fmt.Println("________________________________________________")
	fmt.Println(numroutine)
	fmt.Println("________________________________________________")
	i := 0
	for {
		select {
		case <-ch:
			fmt.Println("Остановка рутины")
			wg.Done()
			return
		default:
			i++
			fmt.Printf("Времени прошло: %d\n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// Остановка при помощи Контекста
func Context() {
	var wg sync.WaitGroup
	// из основной горутины генерируется контекст, который передается в горутину на основе пустого Бэкграунда
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		fmt.Println("Запуск горутины")
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Остановка горутин")
				wg.Done()
				return
			default:
				i++
				fmt.Printf("Времени прошло: %d\n", i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(ctx)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Сработала функция отмены")
	// Удаляем контекст (вызов Done()), что провоцирует всем порожденным рутинам прекратить свою работу, ведь контекста больше нет
	cancel()
	wg.Wait()
}

// атомарные операции работают в асинхронном режиме, к ним есть доступ у всех горутин. Соответственно их можно отследить
func varClosed() {
	var wg sync.WaitGroup
	var a int64
	wg.Add(1)
	go func() {
		fmt.Println("Запуск горутины")
		for a < 1 {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Значение а: %d\n", a)
		}
		fmt.Println("Остановка горутины")
		wg.Done()
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Изменение значения А")
	atomic.AddInt64(&a, 1)
	wg.Wait()

}

// Самый грозный и опасный способ: через панику
func panicClosed() {
	// Без defer загнется и главная горутина, так что самый нежелательный способ
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Отложенная паника: %v", err)
		}
	}()
	for j := 0; j < 5; j++ {
		go func(j int) {
			i := 0
			fmt.Println("Запуск горутины")
			for {
				time.Sleep(100 * time.Millisecond)
				i++
				fmt.Printf("Времени прошло: %d\n", i)
			}
			fmt.Println("Остановка горутины")
			fmt.Println("j")
		}(j)
	}

	time.Sleep(500 * time.Millisecond)
	panic("Генерация паники, остановка всех рутин")
}
