package main

import "fmt"

func main() {
	fmt.Println("Criptografando um texto e colocando no arquivo myfile.data")
	Encrypt()
	fmt.Println("Descriptografando o texto existente no arquivo myfile.data")
	Decrypt()
}
