package main
import "fmt"
func main(){
	var n int
	fmt.Print("enter the number of levels: ")
	fmt.Scan(&n)
	for i:=1;i<=n;i++{
	for j:=1;j<=n-i;j++{
		fmt.Print(" ")}
	for j:=1;j<=i;j++{
		fmt.Print(j)}
	for j:=i-1;j>=1;j--{
		fmt.Print(j)}
	fmt.Println()
}
}
