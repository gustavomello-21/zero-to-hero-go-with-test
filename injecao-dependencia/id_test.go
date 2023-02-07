package injecaodependencia

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}

	Cumprimenta(&buffer, "Gustavo")

	resultado := buffer.String()
	esperado := "Olá, Gustavo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
