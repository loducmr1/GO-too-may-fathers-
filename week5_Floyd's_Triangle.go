package main
import "fmt"
func main() {
	var rows int
	fmt.Print("Enter the number of rows for Floyd's Triangle: ")
	fmt.Scanln(&rows)
	num := 1
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", num)
			num++
		}
		fmt.Println()
	}
} 
