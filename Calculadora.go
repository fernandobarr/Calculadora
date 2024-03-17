package main

import "fmt"

var flag bool

func mais(numero1 float64, numero2 float64) float64 {
	numero1 = numero1 + numero2
	return numero1
}

func menos(numero1 float64, numero2 float64) float64 {
	numero1 = numero1 - numero2
	return numero1
}

func main() {
	var p, num1, num2, numTotal float64
	fmt.Println("Calculadora Simple - Linguagem Go")
	flag = true

	for flag {
		fmt.Println("Digite 1 para Soma")
		fmt.Println("Digite 2 para Subtração")
		fmt.Println("Digite 3 para Sair")
		fmt.Scan(&p)

		if p == 1 {
			fmt.Println("Digite 1 numero")
			fmt.Scan(&num1)
			fmt.Println("Digite 2 numero")
			fmt.Scan(&num2)
			numTotal = mais(num1, num2)
			fmt.Printf("A soma de %v + %v = %v ", num1, num2, numTotal)
			fmt.Println(" ")
		} else if p == 2 {
			fmt.Println("Digite 1 numero")
			fmt.Scan(&num1)
			fmt.Println("Digite 2 numero")
			fmt.Scan(&num2)
			numTotal = menos(num1, num2)
			fmt.Printf("A Subtração de %v + %v = %v ", num1, num2, numTotal)
			fmt.Println(" ")
		} else if p == 3 {
			flag = false
		}
	}

	fmt.Println("Tchau / Bye")
}
