package main

import "fmt"

func main() {
	arr := [100]int{1,1}

	for c := 0; c < len(arr)-2; c++ {
		arr[c+2] = arr[c+1] + arr[c]
	}

	fmt.Println("\nOs", len(arr), "primeiros números de Fibonacci são:\n",arr)
}