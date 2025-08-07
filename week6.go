package main

import "fmt"

func main() {
	fmt.Println("Enter first string: ")
	var s1 string
	fmt.Scan(&s1)
	fmt.Println("Enter second string: ")
	var s2 string
	fmt.Scan(&s2)
	fmt.Print("Concatenation of two strings : " + s1 + s2)
}
