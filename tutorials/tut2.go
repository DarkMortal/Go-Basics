package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	i := 0 // same as var int i = 0
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	fmt.Println("Your number is:", i)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Println("Welcome,", name)
}
