package main

import "fmt"

func main() { //Faça um algoritmo que leia vários números inteiros e mostre a média aritmética entre eles. A leitura deve ser interrompida quando for lido o valor zero.
	var num float64
	var soma float64
	var div float64

	for {
		fmt.Println("Digite um número:")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		soma += num
		div++
	}

	fmt.Println("A média é:", soma/div)
}
