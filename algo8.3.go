package main

import "fmt"

func main() { //Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.
	var num int
	fmt.Println("Digite um número:")
	fmt.Scan(&num)

	for div := 1; div <= num; div++ {
		if num%div == 0 {
			fmt.Println("Os divisores de", num, "são:", num/div)
		}

	}

}
