package ponteiroserros

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	confirmarSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()
		resultado := carteira.Saldo()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	confirmarErro := func(t *testing.T, valor error, esperado string) {
		t.Helper()
		if valor == nil {
			t.Error("esperava um erro, mas nenhum ocorreu.")
		}

		resultado := valor.Error()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}

		carteira.Depositar(10)
		esperado := Bitcoin(10)

		confirmarSaldo(t, carteira, esperado)
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}

		carteira.Retirar(Bitcoin(10))
		esperado := Bitcoin(10)

		confirmarSaldo(t, carteira, esperado)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(29)
		carteira := Carteira{saldo: saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmarSaldo(t, carteira, saldoInicial)
		confirmarErro(t, erro, ErroSaldoInsuficiente.Error())
	})
}
