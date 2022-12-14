package main

import "fmt"

func main() {
	q := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Printf("len= %d cap= %d %v\n", len(q), cap(q), q)

	//Diminui o slice
	fmt.Printf("len= %d cap= %d %v\n", len(q[:0]), cap(q[:0]), q[:0])

	// utiliza tudo antes do slice à esquerda e ignora o restante
	fmt.Printf("len= %d cap= %d %v\n", len(q[:4]), cap(q[:4]), q[:4])

	//Ignora os 2 primeiros itens e pega o restante,
	//diminuindo a capacidade do slice, por ignorar o começo.
	fmt.Printf("len= %d cap= %d %v\n", len(q[2:]), cap(q[2:]), q[2:])

	// Adiciona a um inteiro ao slice dobrando a capacidade do mesmo.
	q = append(q, 110)
	fmt.Printf("len= %d cap= %d %v\n", len(q[0:]), cap(q[0:]), q[0:])

}
