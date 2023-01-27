package arraysslices

import "testing"

func TestSoma(t *testing.T) {

	verificarRetornoCorreto := func(t *testing.T, resultado, esperado int, numeros []int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%d', esperado '%d', dado '%v'", resultado, esperado, numeros)
		}
	}

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		verificarRetornoCorreto(t, resultado, esperado, numeros)
	})
}
