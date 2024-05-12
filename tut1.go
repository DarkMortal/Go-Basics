package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	var mes string = "Hello, World!"
	fmt.Println(mes)
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
}
