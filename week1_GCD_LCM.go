package main
import "fmt"
	func gcd(a,b int) int{
		for b!=0 {
			a,b=b,a%b}
		return a}
	func lcm(a,b int) int{
		return (a*b)/gcd(a,b)}
	func main(){
		var a,b int
		fmt.Print("enter two numbers: ")
		fmt.Scan(&a,&b)
		fmt.Printf("the gcd of %d and %d is %d \n",a,b,gcd(a,b))
		fmt.Printf("the lcm of %d and %d is %d \n",a,b,lcm(a,b))
	}
