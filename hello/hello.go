package main

import (
	"fmt"
	"log"

	"dominio_hospedado.com/nome_modulo"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Recebemos valores da funcao e armazenamos em cada variavel
	message, err := greetings.Hello("Matheus")

	// Caso a funcao retorne um erro, printa na tela
	// E fecha o programa
	if err != nil {
		log.Fatal(err)
	}

	names := []string{"Daniel", "Pedro", "Matheus", "Nicolas"}
	
	messages, err := greetings.Hellos(names)

	fmt.Println(message)

	fmt.Println(messages)
}
