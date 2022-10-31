package Tasks

//Поменять местами два числа без создания временной переменной
import "fmt"

func T13() {
	a := 7
	b := 4
	delta(a, b)
	//kratn(a, b)
	//logic(a, b)
	//golang(a, b)
}

// При помощи суммы и последующей разности
//Так же есть похожая операция, только сначала используется разность, потом сумма и разность, но логика там такая же
func delta(a, b int) {
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("Разность")
	fmt.Println(a, b)
}

// При помощи умножения и последующего деления
func kratn(a, b int) {
	a = a * b
	b = a / b
	a = a / b
	fmt.Println("Кратность")
	fmt.Println(a, b)
}

// При помощи побитового исключаещего ИЛИ (0-1 будет 1, 0-0 и 1-1 будет 0)
func logic(a, b int) {
	a ^= b
	b ^= a
	a ^= b
	fmt.Println("Логическое исключающее или")
	fmt.Println(a, b)
}

// Стандартная операция golang
func golang(a, b int) {
	a, b = b, a
	fmt.Println(a, b)
}
