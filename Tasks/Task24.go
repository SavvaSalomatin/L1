package Tasks

/*Разработать программу нахождения расстояния между двумя точками, которые представлены в виде
структуры Point с инкапсулированными параметрами x,y и конструктором.*/
import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) con(x, y float64) {
	p.x = x
	p.y = y
}

func T24() {
	//Инициализация точек
	var p1, p2 Point
	p1.con(12.4, 10.5)
	p2.con(1.6, 8.5)
	fmt.Println(dist(&p1, &p2))
}

// Функция поиска расстояния на основе Пифагоровой геометрии
func dist(p1, p2 *Point) float64 {
	dX := p1.x - p2.x
	dY := p1.y - p2.y
	return math.Sqrt(math.Pow(dX, 2) + math.Pow(dY, 2))
}
