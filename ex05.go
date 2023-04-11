package main

import (
	"fmt"
	"strconv"
)

//Escreva um programa que receba
//uma string e verifique se ela é um
//número válido em ponto flutuante.
//Informe ao usuário se é ou não.

func main() {
	var x string
	fmt.Print("Escreva uma string: ")
	fmt.Scan(&x)

	numero, err := strconv.ParseFloat(x, 64)
	if err != nil {
		fmt.Printf("A string %s não é um ponto flutuante", x)
	} else {
		fmt.Printf("O ponto flutuante é %f", numero)
	}
}
