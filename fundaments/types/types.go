package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// integer
	fmt.Println(1, 2, 1000)
	fmt.Println("type", reflect.TypeOf(332))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := `Isso é uma
	string com
	várias
	linhas`
	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1))
	fmt.Println("Tamanho da string s1 é ", len(s1))

}
