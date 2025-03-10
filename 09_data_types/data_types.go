/*
Tipos de dados em Go são usados para definir o tipo de informação que uma variável pode armazenar. Os principais tipos de dados em Go são:

1. Tipos básicos:
	- bool: representa um valor booleano (verdadeiro ou falso).
	- string: representa uma sequência de caracteres.
	- int, int8, int16, int32, int64: representam números inteiros com sinal de vários tamanhos.
	- uint, uint8, uint16, uint32, uint64, uintptr: representam números inteiros sem sinal de vários tamanhos.
	- byte: é um outro nome para uint8.
	- rune: é um outro nome para int32, representa um ponto de código Unicode.
	- float32, float64: representam números de ponto flutuante.
	- complex64, complex128: representam números complexos.

*/

package main

import (
	"fmt"
)

func main() {
	// Tipos básicos
	var b bool = true

	var s string = "Olá, Mundo!"

	var i int = 42
	var u uint = 42
	var by byte = 255

	var r rune = 'A'

	var f float64 = 3.14
	var c complex128 = complex(1, 2)

	fmt.Println("bool:", b)       // OUTPUT: bool: true
	fmt.Println("string:", s)     // OUTPUT: string: Olá, Mundo!
	fmt.Println("int:", i)        // OUTPUT: int: 42
	fmt.Println("uint:", u)       // OUTPUT: uint: 42
	fmt.Println("byte:", by)      // OUTPUT: byte: 255
	fmt.Println("rune:", r)       // OUTPUT: rune: 65
	fmt.Println("float64:", f)    // OUTPUT: float64: 3.14
	fmt.Println("complex128:", c) // OUTPUT: complex128: (1+2i)
}
