package main

import "fmt"

func main() {
	i := 1

	// Go n tem aritmética de ponteiros

	var p *int = nil

	p = &i // pegando o endereco da variavel i

	*p++ // desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i)
}
