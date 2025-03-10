package main // Em Go, todo programa é parte de um pacote. Definimos isso usando a palavra-chave package. Neste exemplo, o programa pertence ao pacote main.

import (
	"fmt" // import "fmt" nos permite importar arquivos incluídos no pacote fmt.
)

func main() { // é uma função. Qualquer código dentro de suas chaves {} será executado.
	fmt.Println("Hello World!") // fmt.Println() é uma função disponibilizada pelo pacote fmt. Ela é usada para gerar/imprimir texto.
}

// É possível escrever um código mais compacto, como mostrado abaixo (isso não é recomendado porque torna o código mais difícil de ler)
