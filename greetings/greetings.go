package greetings

import (
	"fmt"
	"math/rand"
	"errors"
)

// A funcao retorna dois valores, uma string e um erro
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// ':=' substitui o uso da declaracao 'var'
	message := fmt.Sprintf(randomFormat(), name)
	
	// nil significa sem erros
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
