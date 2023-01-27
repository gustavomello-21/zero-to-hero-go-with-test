package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	resultado := Repetir("a", 6)
	esperado := "aaaaaa"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 6)
	}
}

func ExampleRepetir() {
	esperado := Repetir("b", 4)
	fmt.Println(esperado)
	// Output: bbbb
}
