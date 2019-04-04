package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	fmt.Printf("O valor de x é %.2f", x)

	// fmt.Println("O valor de x é " + x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs) // imprimi a string texto + a variavel xs, q foi previamente convertida para string
	fmt.Println("O valor de x é", x)    // imprimi a string texto + a variavel x, que n foi convertida, porém a concatenação foi feita com uma virgula

	fmt.Printf("O valor de x é %.2f", x) // imprimi a string texto + as duas primeiras casas decimais da variavel x

}
