/*

Comentários em Go
- É um texto que é ignorado durante a execução.
- Podem ser usados para explicar o código e torná-lo mais legível.
- Também podem ser usados para evitar a execução do código ao testar um código alternativo.

Go suporta comentários de linha única ou de múltiplas linhas.
- Comentários de linha única começam com duas barras (//).
	- Qualquer texto entre // e o final da linha é ignorado pelo compilador (não será executado).
- Comentários de Múltiplas Linhas em Go
*/
//	- Comentários de múltiplas linhas começam com /* e terminam com */. (qualquer texto entre os comentários será ignorado pelo compilador)

/*

Comentário para Prevenir a Execução do Código
- É possível usar comentários para evitar que o código seja executado.
- O código comentado pode ser salvo para referência futura e solução de problemas.
- Isso é útil ao testar um código alternativo sem excluir o código existente.
*/

// Exemplo de comentário para prevenir a execução do código
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	// fmt.Println("Essa linha não será executada")
}
