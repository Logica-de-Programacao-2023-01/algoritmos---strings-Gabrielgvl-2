package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário uma string e
//substitua todas as ocorrências de uma letra
//por outra informada pelo usuário.

func main() {
	var x, a, b string
	fmt.Print("Escreva uma string: ")
	fmt.Scan(&x)
	fmt.Print("Escreva uma letra para ser substituida: ")
	fmt.Scan(&a)
	fmt.Print("Escreva uma letra para substituir: ")
	fmt.Scan(&b)

	resultado := strings.ReplaceAll(x, a, b)
	fmt.Println(resultado)
}
