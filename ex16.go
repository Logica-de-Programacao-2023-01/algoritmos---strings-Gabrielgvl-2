package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário duas strings e informe
//se a segunda é uma substring da primeira.

func main() {
	var primeira, segunda string
	fmt.Print("Escreva a primeira: ")
	fmt.Scan(&primeira)
	fmt.Print("Escreva a segunda: ")
	fmt.Scan(&segunda)

	if strings.Contains(primeira, segunda) {
		fmt.Println("Está contida")
	} else {
		fmt.Println("Não está contida")
	}
}
