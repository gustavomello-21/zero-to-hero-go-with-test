package maps

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é um teste"}
	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é um teste"

		comparaString(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("desconhecida")

		comparaError(t, err, ErrNaoEncontrado)
	})
}

func TestAdiciona(t *testing.T) {
	dicionario := Dicionario{}
	dicionario.Adiciona("teste", "isso é um teste")

	esperado := "isso é um teste"
	resultado, err := dicionario.Busca("teste")
	if err != nil {
		t.Fatal("não foi possível encontrar a palavra adicionada:", err)
	}

	if esperado != resultado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func comparaString(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "test")
	}
}

func comparaError(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}
