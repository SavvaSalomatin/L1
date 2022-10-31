package Tasks

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).
import "fmt"

// Структура человека с паспортными данными и физическими параметрами
type Human struct {
	id int
	passport
	weight int
	height int
}

// Структура паспортных данных
type passport struct {
	firstName   string
	lastName    string
	yearOfBirth int
}

// Структура действия человека
type Action struct {
	Human
	energy   int
	duration minute
}

// Пользовательский тип, чтобы не скучать
type minute int

// Реализация метода вывода паспортных данных
func (p passport) getPassportData() {
	fmt.Println("Passport data")
	fmt.Println(p.lastName)
	fmt.Println(p.firstName)
	fmt.Println(p.yearOfBirth)
}

// Реализация метода вывода данных человека с встраиванием метода структуры паспортных данных
func (h Human) getDataOfHuman() {
	fmt.Println("Phisical data of human")
	fmt.Println(h.weight)
	fmt.Println(h.height)
	h.getPassportData()
}

// Пример вызова группы встроенных методов по названию, если есть с одинаковыми названиями - обращение по полному пути
func T1() {
	info := passport{firstName: "Кирилл", lastName: "Ильин", yearOfBirth: 2001}
	people := Human{id: 10, passport: info, weight: 60, height: 170}
	peopleWalk := Action{Human: people, energy: 500, duration: 20}
	peopleWalk.getPassportData()
	peopleWalk.getDataOfHuman()
}
