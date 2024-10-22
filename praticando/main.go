package main

import (
	"fmt"
)

// Struct para armazenar informações de uma pessoa
type Person struct {
	Name string
	Age  int
}

// Função para verificar se a pessoa é maior de idade
func isAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

// Função principal
func main() {
	// Array de pessoas
	people := []Person{
		{"Alice", 25},
		{"Bob", 16},
		{"Charlie", 30},
		{"Diana", 17},
	}

	// Loop para verificar se cada pessoa é maior de idade
	for _, person := range people {
		if isAdult(person.Age) {
			fmt.Printf("%s é maior de idade.\n", person.Name)
		} else {
			fmt.Printf("%s não é maior de idade.\n", person.Name)
		}
	}
}
