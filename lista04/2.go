package main

import "fmt"

func main(){

var n10, n5, soma1, soma2 int
var numeros10 [10]int
var numeros5 [5]int
var pares []int
var impares []int
var result1 []int
var result2 []int
soma5 := 0

for i := 0; i < 10; i++{
fmt.Scan(&n10)
numeros10[i] = n10
} 
for i := 0; i < 5; i++{
	fmt.Scan(&n5)
	numeros5[i]=n5 
	soma5 += n5 
	} 
for _, valor := range numeros10{
	if valor % 2 == 0 {
		pares = append(pares, valor)
	} else {
		impares = append(impares, valor)
	}
}
for i := 0; i < len(pares); i++{
	
	soma1 = pares[i] + soma5
	result1 = append(result1, soma1)
}
for i := 0; i < len(impares); i++{
	
	soma2 = impares[i] + soma5
	result2 = append(result2, soma2)
}
fmt.Println("Primeiro vetor: ", numeros10)
fmt.Println("Segundo vetor: ", numeros5)
fmt.Println("Primeiro vetor resultante: ", result1)
fmt.Println("Segundo vetor resultante: ", result2)
}