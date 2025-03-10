/*

Constantes são valores que não podem ser alterados durante a execução do programa.
São declaradas com a palavra-chave const.

Uma boa prática é declarar constantes em letras maiúsculas.

Existem dois tipos de constantes: tipadas e não tipadas.

Quando uma constante tem um valor atribuido, esse valor não pode ser alterado.

*/

package main

import (
	"fmt"
)

// Essa é uma constante não tipada
const PI = 3.14

// Essa é uma constante tipada
const PI2 float64 = 3.14

// Definição de múltiplas constantes
const (
	A int = 1
	B     = 3.14
	C     = "Hello"
)

func main() {
	fmt.Println(PI)
	fmt.Println(PI2)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
