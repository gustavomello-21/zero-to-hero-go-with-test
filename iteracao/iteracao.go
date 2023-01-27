package iteracao

func Repetir(caracter string, qtd int) string {
	var repeticoes string

	for i := 0; i < qtd; i++ {
		repeticoes += caracter
	}

	return repeticoes
}
