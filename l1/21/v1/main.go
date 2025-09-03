package main

// теоретическая демонстрация

import "fmt"

///from documentation:
// Требуется для реализации:
// Интерфейс Target, описывающий целевой интерфейс (тот интерфейс с которым наша система хотела бы работать);
// Класс Adaptee, который наша система должна адаптировать под себя;
// Класс Adapter, адаптер реализующий целевой интерфейс.

// адаптер имеет метод someOperation(), который реализует логику адаптируемого класса someAdaptedOperation().
// Но может быть вызван как объект, удовлетворяющий интерфейсу Target

type Target interface {
	someOperation()
}

type Adaptee struct {
}

func (adaptee *Adaptee) someAdaptedOperation() {
	fmt.Println("this is someAdaptedOperation")
}

type Adapter struct {
	*Adaptee
}

func (adapter *Adapter) someOperation() {
	adapter.someAdaptedOperation()
}

func NewAdapter(a *Adaptee) Target {
	return &Adapter{a}
}

func main() {
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)
	adapter.someOperation()
}
