package personaldata

import "fmt"

// Структура Personal для хранения данных о пользователе
type Personal struct {
	Name   string  // Имя пользователя
	Weight float64 // Вес пользователя
	Height float64 // Рост пользователя
}

// Метод Print выводит данные структуры на экран
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
