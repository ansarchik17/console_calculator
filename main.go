package main

import "fmt"

func main() {
	const baseRate = 5.50
	const taxRate = 0.12
	const distanceRate = 2.0
	const fragile = 0.2

	var name string
	var weight float64
	var distance int
	var amount int

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	fmt.Print("Вес груза (кг): ")
	fmt.Scan(&weight)

	fmt.Print("Дистанция (км): ")
	fmt.Scan(&distance)

	fmt.Print("Количество хрупких упаковок: ")
	fmt.Scan(&amount)

	finalPrice := (weight*baseRate)*(1+fragile*float64(amount)) +
		(float64(distance) * distanceRate)

	finalPrice *= (1 + taxRate)

	fmt.Println("\n--- Отчет о доставке ---")
	fmt.Println("Отправитель:", name)
	fmt.Printf("Итоговая стоимость: %.2f\n", finalPrice)
}
