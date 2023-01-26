package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefoxoOlaFrances = "Bonjour, "
const prefixoOlaRusso = "Привет, "
const espanhol = "espanhol"
const frances = "frances"
const russo = "russo"

func Ola(name, idioma string) string {
	if name == "" {
		name = "mundo"
	}

	return prefixoDeSaudacao(idioma) + name
}

func prefixoDeSaudacao(idioma string) (prefix string) {
	switch idioma {
	case espanhol:
		prefix = prefixoOlaEspanhol
	case frances:
		prefix = prefoxoOlaFrances
	case russo:
		prefix = prefixoOlaRusso
	default:
		prefix = prefixoOlaPortugues
	}

	return
}

func main() {
	fmt.Println(Ola("mundo", "espanhol"))
}
