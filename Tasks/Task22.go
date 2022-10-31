package Tasks

/*Разработать программу, которая перемножает, делит, складывает, вычитает две числовых
переменных a,b, значение которых > 2^20.*/
import (
	"fmt"
	"math/big"
)

// Хоть такие числа и умещаются в диапазон int32, то некоторые математические операции могут
//привести к трагическим последствиям. Для работы с большими данными используется библиотека
//big, суть работы которой достаточно проста - обращение по указателям и инциализация
//переменных по указателям,диапазоны таких значений могут быть даже выше int64. Отсюда и проблема применения стандартных
//математических операторов, используются методы библиотеки
func T22() {
	var x, y, z = new(big.Int), new(big.Int), new(big.Int)
	x.SetInt64(456456455678583)
	y.SetInt64(3459834096456456456)
	*z = sum(x, y)
	fmt.Println(z)
	*z = mod(y, x)
	fmt.Println(z)
	*z = div(y, x)
	fmt.Println(z)
	*z = multiplicate(y, x)
	fmt.Println(z)
	*z = difference(y, x)
	fmt.Println(z)
}

// Сумма
func sum(a, b *big.Int) big.Int {
	c := new(big.Int)
	c.Add(a, b)
	return *c
}

// Остаток от деления
func mod(a, b *big.Int) big.Int {
	c := new(big.Int)
	c.Mod(a, b)
	return *c
}

// Основание от деления
func div(a, b *big.Int) big.Int {
	c := new(big.Int)
	c.Div(a, b)
	return *c
}

// Перемножение
func multiplicate(a, b *big.Int) big.Int {
	c := new(big.Int)
	c.Mul(a, b)
	return *c
}

// Вычитание
func difference(a, b *big.Int) big.Int {
	c := new(big.Int)
	c.Sub(a, b)
	return *c
}
