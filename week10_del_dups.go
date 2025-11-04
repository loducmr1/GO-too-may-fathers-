package main
import "fmt"
func removeDuplicates(s []string) []string {
	bucket := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := bucket[str]; !ok {
			bucket[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {
	array := []string{"abc", "cde", "efg", "efg", "abc", "cde"}
	fmt.Println("The given array of string is:", array)
	fmt.Println()
	result := removeDuplicates(array)
	fmt.Println("The array obtained after removing the duplicate entries is:", result)
}
