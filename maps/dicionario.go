package maps

import "errors"

type Dicionario map[string]string

var ErrNaoEncontrado = errors.New("não foi possível encontrar a palavra que você procura")

func (d Dicionario) Busca(chave string) (string, error) {
	valor, exist := d[chave]
	if !exist {
		return "", ErrNaoEncontrado
	}

	return valor, nil
}

func (d Dicionario) Adiciona(chave, valor string) {
	d[chave] = valor
}
