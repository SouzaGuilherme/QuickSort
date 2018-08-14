/*
Developer: SouzaGuilherme
Nome: Guilherme de Souza da Silva
Matricula: 16200758
Prof.: Matheus Natchigal
Sites referencias:	
	https://pt.wikipedia.org/wiki/Quicksort
	https://golang.org/doc/
	https://gist.github.com/vderyagin/4161347
	https://stackoverflow.com/questions/15802890/idiomatic-quicksort-in-go
*/

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	j := 0
	vectorLen := 1000	// define len vector 1000
	for j < 21 { 
		k := 0
		// for on controll number of repeat
		for k < 10 {
			fmt.Printf("\n\nInicio testes com %d rodada %d\n\n", vectorLen, k)
			executation(vectorLen)
			fmt.Printf("\n\nFim testes com %d\n\n", vectorLen)
			k++
		}
		if j == 0 {
			vectorLen = 5000
			j++
		}else {
			vectorLen += 5000 //trocar pra 5000
			j++
		}
	}
}

func executation(vectorLen int) {
	slice := generationSlice(vectorLen)		// function default in go for personalise generate vector
	// Print in vector oculted
	// fmt.Println("\n==== Unsorted ====\n", slice)
	start:= time.Now()		// Function deafult in go return time in millisecond
	slice = quickSort(slice)
	end := time.Now()		// Function deafult in go return time in millisecond
	// Print in vector oculted
	// fmt.Println("\n==== Sorted ====\n", slice, "\n\n")
	result := end.Sub(start)
	resultFinal := result.Seconds()	// Function deafult in go return in seconds
	fmt.Println("RESULTADO DO TEMPO DE EXECUCAO: ", resultFinal)
}

func generationSlice(size int) []int {
	slice := make([]int, size, size)	// Function deafult criated vector
	for i:= 0; i < size; i++ {
		slice[i] = rand.Intn(1000000) - rand.Intn(1000000)		// Random number for vector
	}
	return slice	
}

func quickSort(vector []int) []int {
	if len(vector) < 2 {
		return vector
	}
	left, right := 0, len(vector)-1
	pivot := len(vector) - (len(vector)-1)		// position zero vector for pivot
	vector[pivot], vector[right] = vector[right], vector[pivot]
	for i, _ := range vector {		// Function range in go
		// Separates the smaller ones than the pivot to the left
		if vector[i] < vector[right] {
			vector[left], vector[i] = vector[i], vector[left]
			left++
		}
	}
	vector[left], vector[right] = vector[right], vector[left]
	quickSort(vector[:left])		// Recursive call of the left-side function
	quickSort(vector[left+1:])		// Recursive call of the remaining function
	return vector
}
