/*
Os operadores em Go são utilizados para realizar operações em operandos e valores. Go suporta os seguintes tipos de operadores:

Operadores Aritméticos:
+  serve para adição
-  serve para subtração
*  serve para multiplicação
/  serve para divisão
%  serve para módulo
++ serve para incremento
-- serve para decremento

Operadores de Atribuição:
=  serve para atribuição
+= serve para adição e atribuição
-= serve para subtração e atribuição
*= serve para multiplicação e atribuição
/= serve para divisão e atribuição
%= serve para módulo e atribuição
&= serve para AND bit a bit e atribuição
|= serve para OR bit a bit e atribuição
^= serve para XOR bit a bit e atribuição
<<= serve para deslocamento à esquerda e atribuição
>>= serve para deslocamento à direita e atribuição

Operadores de Comparação:
== serve para igualdade
!= serve para desigualdade
>  serve para maior que
<  serve para menor que
>= serve para maior ou igual que
<= serve para menor ou igual que

Operadores Lógicos:
&& serve para operador lógico E
|| serve para operador lógico OU
!  serve para operador lógico NÃO

Operadores Bitwise:
&  serve para AND bit a bit
|  serve para OR bit a bit
^  serve para XOR bit a bit
<< serve para deslocamento à esquerda
>> serve para deslocamento à direita
*/

package main

import (
	"fmt"
)

func main() {
	// Operadores Aritméticos
	a, b := 10, 5
	fmt.Println("Adição:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Módulo:", a%b)
	a++
	fmt.Println("Incremento:", a)
	a--
	fmt.Println("Decremento:", a)

	// Operadores de Atribuição
	c := 10
	fmt.Println("Atribuição:", c)
	c += 3
	fmt.Println("Adição e atribuição:", c)
	c -= 3
	fmt.Println("Subtração e atribuição:", c)
	c *= 3
	fmt.Println("Multiplicação e atribuição:", c)
	c /= 3
	fmt.Println("Divisão e atribuição:", c)
	c %= 3
	fmt.Println("Módulo e atribuição:", c)
	c &= 3
	fmt.Println("AND bit a bit e atribuição:", c)
	c |= 3
	fmt.Println("OR bit a bit e atribuição:", c)
	c ^= 3
	fmt.Println("XOR bit a bit e atribuição:", c)
	c <<= 3
	fmt.Println("Deslocamento à esquerda e atribuição:", c)
	c >>= 3
	fmt.Println("Deslocamento à direita e atribuição:", c)

	// Operadores de Comparação
	fmt.Println("Igualdade:", a == b)
	fmt.Println("Desigualdade:", a != b)
	fmt.Println("Maior que:", a > b)
	fmt.Println("Menor que:", a < b)
	fmt.Println("Maior ou igual que:", a >= b)
	fmt.Println("Menor ou igual que:", a <= b)

	// Operadores Lógicos
	d, e := true, false
	fmt.Println("E lógico:", d && e)
	fmt.Println("OU lógico:", d || e)
	fmt.Println("NÃO lógico:", !d)

	// Operadores Bitwise
	f, g := 10, 2
	// Operadores Bitwise
	fmt.Println("AND bit a bit:", f&g)            // AND bit a bit: compara cada bit dos operandos e retorna 1 se ambos os bits forem 1
	fmt.Println("OR bit a bit:", f|g)             // OR bit a bit: compara cada bit dos operandos e retorna 1 se pelo menos um dos bits for 1
	fmt.Println("XOR bit a bit:", f^g)            // XOR bit a bit: compara cada bit dos operandos e retorna 1 se os bits forem diferentes
	fmt.Println("Deslocamento à esquerda:", f<<g) // Deslocamento à esquerda: desloca os bits de f para a esquerda g vezes
	fmt.Println("Deslocamento à direita:", f>>g)  // Deslocamento à direita: desloca os bits de f para a direita g vezes
}
