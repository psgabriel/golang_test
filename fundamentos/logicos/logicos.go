package main

import "fmt"

func compras(trab1 bool, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    // trab1 e trab2 true
	comprarTv32 := trab1 != trab2    // ou exclusivo. se apenas um trabalho for true
	comprarSorvete := trab1 || trab2 // ou

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudavel: %t",
		tv50, tv32, sorvete, !sorvete)
}
