package main

import "fmt"

// Интерфейс для всех животных
type Animal interface {
	MakeSound() string
	Move() string
	GetInfo() string
	GetType() string
}

// Интерфейс для животных, которые умеют плавать
type Swimmer interface {
	Swim() string
	CanSwim() string
}

// Собака
type Dog struct {
	Name string
}

func (d Dog) MakeSound() string {
	return "Гав-гав!"
}

func (d Dog) Move() string {
	return "Собака бегает"
}

func (d Dog) GetInfo() string {
	return fmt.Sprintf("Собака %s - верный друг!", d.Name)
}

func (d Dog) GetType() string {
	return "Млекопитающее"
}

func (d Dog) Bark() string {
	return "Гав-гав-гав!"
}

func (d Dog) CanSwim() bool {
	return true
}

func (d Dog) Swim() string {
	return "Собака плавает лапами"
}

// Кот
type Cat struct {
	Name string
}

func (c Cat) MakeSound() string {
	return "Мяу!"
}

func (c Cat) Move() string {
	return "Кот ходит мягко"
}

func (c Cat) GetInfo() string {
	return fmt.Sprintf("Кот %s - пушистый и игривый!", c.Name)
}

func (c Cat) GetType() string {
	return "Млекопитающее"
}

func (c Cat) Meow() string {
	return "Мяу-мяу!"
}

func (c Cat) CanSwim() bool {
	return false
}

// Лошадь
type Horse struct {
	Name string
}

func (h Horse) MakeSound() string {
	return "И-го-го!"
}

func (h Horse) Move() string {
	return "Лошадь бежит быстро"
}

func (h Horse) GetInfo() string {
	return fmt.Sprintf("Лошадь %s - грациозное животное!", h.Name)
}

func (h Horse) GetType() string {
	return "Млекопитающее"
}

func (h Horse) Neigh() string {
	return "И-го-го-го!"
}

func (h Horse) CanSwim() bool {
	return true
}

func (h Horse) Swim() string {
	return "Лошадь плавает мощными движениями"
}

// Курица
type Chicken struct {
	Name string
}

func (c Chicken) MakeSound() string {
	return "Куд-кудах!"
}

func (c Chicken) Move() string {
	return "Курица ходит"
}

func (c Chicken) GetInfo() string {
	return fmt.Sprintf("Курица %s - несушка!", c.Name)
}

func (c Chicken) GetType() string {
	return "Птица"
}

func (c Chicken) Cluck() string {
	return "Куд-кудах-кудах!"
}

func (c Chicken) CanSwim() bool {
	return false
}

// Верблюд
type Camel struct {
	Name string
}

func (c Camel) MakeSound() string {
	return "Р-р-р!"
}

func (c Camel) Move() string {
	return "Верблюд ходит медленно"
}

func (c Camel) GetInfo() string {
	return fmt.Sprintf("Верблюд %s - корабль пустыни!", c.Name)
}

func (c Camel) GetType() string {
	return "Млекопитающее"
}

func (c Camel) Grunt() string {
	return "Р-р-р!"
}

func (c Camel) CanSwim() bool {
	return false
}

func main() {
	rex := Dog{"Рекс"}
	murka := Cat{"Мурка"}
	sparky := Horse{"Спарки"}
	henrietta := Chicken{"Генриетта"}
	camel := Camel{"Башир"}

	animals := []Animal{rex, murka, sparky, henrietta, camel}

	for _, animal := range animals {
		fmt.Println(animal.GetInfo())
		fmt.Println("Звук:", animal.MakeSound())
		fmt.Println("Движение:", animal.Move())
		fmt.Println("Тип:", animal.GetType())
		if swimmer, ok := animal.(Swimmer); ok {
			fmt.Println("Умеет плавать:", swimmer.CanSwim())
			fmt.Println("Способ плавания:", swimmer.Swim())
		}
		fmt.Println("----")
	}
}
