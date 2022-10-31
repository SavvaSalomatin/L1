package Tasks

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
import "fmt"

// Лучше всего для этого подойдет хешмапа, где ключом будет являться диапазон, значениями - числа множества этого диапазона
func T10() {
	//Создается результирующая хешмапа и исходные данные в виде последовательности(массив)
	group := make(map[int][]float32)
	temps := [10]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -5, 4}
	//Проходя по массиву, я целочисленно делю значения температур, уникальные из которых автоматически записываются в ключ
	// Если встречается общий ключ, то данные просто дописываются в значение при помощи append-а, диапазон
	//берется по разряду числа, поэтому -5 и 4 будут иметь одну группу
	for i := range temps {
		group[int(temps[i]/10)*10] = append(group[int(temps[i]/10)*10], temps[i])
	}
	fmt.Println(group)
}
