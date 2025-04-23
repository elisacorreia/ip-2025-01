package main

import "fmt"

func main(){
	var n int
	var numeros[10]int
	
	for i := 0; i < 10; i++{
	fmt.Scan(&n)
	numeros[i]=n 
} 
for i, valor := range numeros{
	if valor > 50 {
		fmt.Println("Valor: ", valor, "Índice: ", i)
	} else {
		j:=0
		j++
	}
}
if j == 10 {
	fmt.Println("Não existe nenhum número superior a 50")
}
}