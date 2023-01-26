package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}
	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Gustavo", "")
		esperado := "Olá, Gustavo"

		verificaMensagemCorreta(t, resultado, esperado)

		t.Run("em espanhol", func(t *testing.T) {
			resultado := Ola("Elodie", "espanhol")
			esperado := "Hola, Elodie"

			verificaMensagemCorreta(t, resultado, esperado)
		})

		t.Run("em frances", func(t *testing.T) {
			resultado := Ola("Pierre", "frances")
			esperado := "Bonjour, Pierre"

			verificaMensagemCorreta(t, resultado, esperado)
		})

		t.Run("em russo", func(t *testing.T) {
			resultado := Ola("Katarina", "russo")
			esperado := "Привет, Katarina"

			verificaMensagemCorreta(t, resultado, esperado)
		})
	})

	t.Run("diz olá mundo quando a string for vazia", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}
