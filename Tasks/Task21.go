package Tasks

//Реализовать паттерн «адаптер» на любом примере
import "fmt"

//Суть паттерна заключается в использовании методов внешних пакетов, к исходному коду которых нет доступа
//Задача заключается в имитации совместимости реализаций этих методов со своими(несовместимыми) и
//написания структуры, методы которой позволят имплементировать несовместимые методы,
//Работу принципов паттерна адаптер я представил на уровне абстракции с разъемами для передачи данных
// Интерфейсы ноутбука и смартфона с методами работы их разъемов
type laptop interface {
	usb_Port()
}

type smartphone interface {
	typeC_Port()
}

// Описание разъемов вместе с принципами их работы (структура+интерфейс)
type thunderbold struct {
}

func (tC *thunderbold) typeC_Port() {
	fmt.Println("Тандерболт может быть вставлен ")
}

type usb_3 struct {
}

func (usbb *usb_3) usb_Port() {
	fmt.Println("USB может быть вставлен ")
}

// Для примера, из-за несовместимости разъемов ноутбука и тандерболта было принято решение написать адаптер под тандерболт из обычного usb
type thunderboldAdapter struct {
	device_tb *thunderbold
}

func (d *thunderboldAdapter) usb_Port() {
	d.device_tb.typeC_Port()
}

// Пользователь владеет и тем и тем устройством
type user struct {
}

func (u user) useAnyPort_tb(device laptop) {
	device.usb_Port()
}

func (u user) useAnyPort_usb(device smartphone) {
	device.typeC_Port()
}

func T21() {
	user := user{}
	typeC := thunderbold{}
	usb := usb_3{}
	//По правилам интерфейса ноутбука, вставить usb н может без проблем
	user.useAnyPort_tb(&usb)
	//А вот для тандерболта необходим адаптер (аналогия реального переходника-адаптера для кабелей, либо специального ПО, адаптирующего потоки данных)
	adapter := thunderboldAdapter{
		device_tb: &typeC,
	}
	user.useAnyPort_tb(&adapter)
	user.useAnyPort_usb(&typeC)
	//Для смартфона адаптер не был реализован, поэтому у пользователя нет никакой возможности воспользоваться шнуром от ноутбука
	//user.useAnyPort_usb(usb)
}
