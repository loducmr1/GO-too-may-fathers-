package main
import "fmt"
func isPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var input string
    fmt.Print("Enter a string to check if it is a palindrome: ")
    fmt.Scanln(&input)

    if isPalindrome(input) {
        fmt.Println("Yes, it's a palindrome!")
    } else {
        fmt.Println("No, it's not a palindrome.")
    }
}
