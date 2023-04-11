package main

import (
	"fmt"
	"strings"
)

// Escreva um programa que receba uma
//string e verifique se ela é um palíndromo.
//Informe ao usuário se é ou não.

func main() {
	var x string
	fmt.Print("Escreva uma string: ")
	fmt.Scan(&x)

	letras := strings.Split(x, "")

	if len(letras)%2 != 0 {
		palindromo := true
		for i := 0; i < len(letras)/2; i++ {
			if letras[i] != letras[len(letras)-1-i] {
				palindromo = false
				break
			}
		}
		if palindromo {
			fmt.Printf("É palindromo")
		} else {
			fmt.Printf("A string %s não é palindromo", x)
		}
	} else {
		fmt.Printf("A string %s não é palindromo", x)
	}
}
