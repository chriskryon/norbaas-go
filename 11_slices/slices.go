/*

Slices em Go, são semelhantes a arrays, mas são poderosos e flexiveis.
Mas o comprimento pode aumentar ou diminuir, conforme necessário

Existem várias maneiras de se criar um Slice
- Usando o formato de tipo de dados [] { valores }
		slice1 := []int{1, 2, 3, 4, 5}

- Crie uma fatia de uma matriz
	array := [5]int{1, 2, 3, 4, 5}
	slice2 := array[1:4]

- Usando a função make()
	slice3 := make([]int, 5)


- Em Go, há duas funções que podem ser usadas para retornar o comprimento e a capacidade de uma fatia:
	len()função - retorna o comprimento da fatia (o número de elementos na fatia)
	cap()função - retorna a capacidade da fatia (o número de elementos que a fatia pode aumentar ou diminuir)


*/

package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{} // slice vazio

	// A função len() retorna o comprimento da fatia
	// A função cap() retorna a capacidade da fatia

	fmt.Println(len(myslice1)) // 0
	fmt.Println(cap(myslice1)) // 0
	fmt.Println(myslice1)      // []

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2)) // 4
	fmt.Println(cap(myslice2)) // 4
	fmt.Println(myslice2)      // [Go Slices Are Powerful]

	// Crie uma fatia de uma matriz
	arr1 := [6]int{10, 11, 12, 13, 14, 15} // matriz
	myslice := arr1[2:4]                   // fatia

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	// Usando a função make()
	myslice3 := make([]int, 5, 10)

	fmt.Printf("myslice1 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	// with omitted capacity
	myslice4 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	// A diferença é que o make()função aceita um terceiro argumento que especifica a capacidade da fatia.
	// Se você não especificar a capacidade, a capacidade será igual ao comprimento.

}
