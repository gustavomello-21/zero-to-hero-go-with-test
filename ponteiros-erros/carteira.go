package ponteiroserros

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	String() string
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

func (c *Carteira) Retirar(valor Bitcoin) error {
	if valor > c.saldo {
		return errors.New(ErroSaldoInsuficiente.Error())
	}
	c.saldo -= valor

	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
