package main
import "fmt"
func removeDuplicates(arr []int) []int {
	encountered := make(map[int]bool)
	var result []int
	for _, v := range arr {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	array := []int{1,2,3,2,4,5,1,6,3}
	fmt.Println("The given array is:", array)
	fmt.Println()
	result := removeDuplicates(array)
	fmt.Println("The array obtained after removing the duplicate entries is:", result)
}
