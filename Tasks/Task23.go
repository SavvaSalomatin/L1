package Tasks

//Удалить i-ый элемент из слайса.
import "fmt"

func T23() {
	a := []string{"Asdfsdf", "Bsdfsdf", "Csdfsdf", "Dsdfsd", "Esdfsdf"}
	lasttodeleted(2, a)
	copytodeleted(2, a)
}

// Удаление и замена последним элементом, если не важен порядок
func lasttodeleted(i int, a []string) {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	fmt.Println(a)
}

// Копирование по обе стороны от удаляемого элемента и склеивание, если важен порядок
func copytodeleted(i int, a []string) {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	fmt.Println(a)
}
