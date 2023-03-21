package main

import "fmt"

func main() { //Faça um algoritmo que leia vários números inteiros e mostre o maior deles. A leitura deve ser interrompida quando for lido o valor zero.
	var num int
	max := 0

	for {
		fmt.Println("Digite um número:")
		fmt.Scan(&num)
		if num == 0 {
			break
		} else if num > max {
			max = num
		}

	}
	if max != 0 {
		fmt.Println("O maior número digitado foi:", max)
	} else {
		fmt.Println("Nenhum número válido foi digitado.")
	}
}
