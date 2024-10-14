package main

import "fmt"

func main() {
	var opcion int
	var temperatura float64
	fmt.Println("Conversor de Temperatura")
	fmt.Println("1. Celsius a Fahrenheit")
	fmt.Println("2. Fahrenheit a Celsius")
	fmt.Print("Elige una opcion 1 o 2: ")
	fmt.Scanln(&opcion)

	if opcion == 1 {
		fmt.Print("Introduce la temperatura en Celsius: ")
		fmt.Scanln(&temperatura)
		fahrenheit := (temperatura * 9 / 5) + 32
		fmt.Println("La temperatura en fahrenheit es: ", fahrenheit)
	} else if opcion == 2 {
		fmt.Print("Introduce la temperatura en Fahrenheit: ")
		fmt.Scanln(&temperatura)
		celsius := (temperatura - 32) * 5 / 9
		fmt.Println("La temperatura en Celsius es: ", celsius)
	} else {
		fmt.Println("Opcion no valida")
	}
}
