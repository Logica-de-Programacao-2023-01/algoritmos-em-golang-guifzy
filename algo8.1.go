package main

import "fmt"

func main() { //Faça um algoritmo que leia um número inteiro entre 1 e 7 e mostre o dia da semana correspondente (1 = domingo, 2 = segunda-feira, etc.).
	var num int

	fmt.Print("Digite um número inteiro entre 1 e 7: ")
	fmt.Scan(&num)

	if num == 1 {
		fmt.Println("Domingo")
	} else if num == 2 {
		fmt.Println("Segunda")
	} else if num == 3 {
		fmt.Println("Terça")
	} else if num == 4 {
		fmt.Println("Quarta")
	} else if num == 5 {
		fmt.Println("Quinta")
	} else if num == 6 {
		fmt.Println("Sexta")
	} else {
		fmt.Println("Sábado")
	}
}
