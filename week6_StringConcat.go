package main
import "fmt"
func main() {
	var first, second string
	fmt.Print("Enter First String ")
	fmt.Scan(&first)
	fmt.Print("Enter Second String ")
	fmt.Scan(&second)
	fmt.Print("Concat of two string: ",first+second)

}
