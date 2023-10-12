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

// Retorna um map e um erro, sendo map um dicionario do python (chave, valor)
// map[tipo_chave]tipo_valor
func Hellos(names []string) (map[string]string, error) {
	// Para criar um map voce precisa usar make() e passar o tipo de map
	messages := make(map[string]string)

	// range retorna o indice e o valor
	// e nao queremos o range, entao "_"
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
