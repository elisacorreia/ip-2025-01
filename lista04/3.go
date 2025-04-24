package main 
import "fmt"
func main(){
	var n10 int 
	var numeros10 [10]int
	var pares []int
	var impares []int
	soma := 0 

	for i := 0; i < 10; i++{
		fmt.Scan(&n10)
		numeros10[i] = n10
	} 
	
	for _, valor := range numeros10{
			if valor % 2 == 0 {
				pares = append(pares, valor)
				soma += valor
			} else {
				impares = append(impares, valor)
			}
	}
	fmt.Println("a) ", pares)
	fmt.Println("b) ", soma)
	fmt.Println("c) ", impares)
	fmt.Println("d) ", len(impares))
}