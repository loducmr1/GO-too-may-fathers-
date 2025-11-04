Go Programming Lab Manual (Weeks 1-12)This file contains all the Go programs from the Week 1 to Week 12 experiments, based on the R20 Lab Manual.Week 1: Write a Go Program to find LCM and GCD of given two numbersCode:package main

import "fmt"

func lcm(temp1 int, temp2 int) {
	var lcmnum int = 1
	if temp1 > temp2 {
		lcmnum = temp1
	} else {
		lcmnum = temp2
	}

	for {
		if lcmnum%temp1 == 0 && lcmnum%temp2 == 0 {
			fmt.Printf("LCM of %d and %d is %d", temp1, temp2, lcmnum)
			break
		}
		lcmnum++
	}
	return
}

func gcd(temp1 int, temp2 int) {
	var gcdnum int
	for i := 1; i <= temp1 && i <= temp2; i++ {
		if temp1%i == 0 && temp2%i == 0 {
			gcdnum = i
		}
	}
	fmt.Printf("GCD of %d and %d is %d", temp1, temp2, gcdnum)
	return
}

func main() {
	var n1, n2, action int
	fmt.Println("Enter two positive integers:")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Println("Enter 1 for LCM and 2 for GCD")
	fmt.Scanln(&action)

	switch action {
	case 1:
		lcm(n1, n2)
	case 2:
		gcd(n1, n2)
	}
}
Input:Enter two positive integers:
3
6
Enter 1 for LCM and 2 for GCD
1
Output:LCM of 3 and 6 is 6
Execution Steps:Run the program.Enter two numbers (e.g., 3 and 6).Enter 1 to calculate LCM or 2 for GCD.The result is printed.Week 2: Write a Go Program to print pyramid of numbersCode:package main

import "fmt"

func main() {
	var rows, k, temp, temp1 int
	fmt.Print("Enter number of rows :")
	fmt.Scan(&rows)
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
			temp++
		}
		for {
			if temp <= rows-1 {
				fmt.Printf("%d", i+k)
				temp++
			} else {
				temp1++
				fmt.Printf("%d", (i + k - 2*temp1))
			}
			k++
			if k == 2*i-1 {
				break
			}
		}
		temp = 0
		temp1 = 0
		k = 0
		fmt.Println("")
	}
}
Input:Enter number of rows : 5
Output:    1
   232
  34543
 4567654
567898765
Execution Steps:Run the program.Enter the number of rows (e.g., 5).The pyramid is printed based on nested loop logic.Week 3: Write a program to use struct that is imported from another packageThis experiment involves multiple files/packages.Code (main.go):package main

import (
	parent "family/father"
	child "family/father/son"
	"fmt"
)

func main() {
	f := new(parent.Father)
	fmt.Println(f.Data("Mr. Jeremy Maclin"))
	c := new(child.Son)
	fmt.Println(c.Data("Riley Maclin"))
}
Code (family/father/father.go):package father

import "fmt"

func init() {
	fmt.Println("Father package initialized")
}

type Father struct {
	Name string
}

func (f Father) Data(name string) string {
	f.Name = "Father: " + name
	return f.Name
}
Code (family/father/son/son.go):package son

import "fmt"

func init() {
	fmt.Println("Son package initialized")
}

type Son struct {
	Name string
}

func (s Son) Data(name string) string {
	s.Name = "Son : " + name
	return s.Name
}
Input:None. Values are hardcoded in the main function.Output:Father package initialized
Son package initialized
Father: Mr. Jeremy Maclin
Son: Riley Maclin
Execution Steps:Create the directory structure family/father/son.Place father.go in the family/father directory.Place son.go in the family/father/son directory.Place main.go outside the family directory (e.g., in the root).Run go run main.go.Go imports the father and son packages, triggering their init() functions first.The main function executes, creating new structs and calling their methods.The results are printed to the console.Week 4: Write a Go Program to calculate standard deviation in Math packageCode:package main

import (
	"fmt"
	"math"
)

func main() {
	var num [10]float64
	var sum, mean, sd float64
	fmt.Println("****** Enter 10 elements *******")
	for i := 1; i <= 10; i++ {
		fmt.Printf("Enter %d element: ", i)
		fmt.Scan(&num[i-1])
		sum += num[i-1]
	}
	mean = sum / 10
	for j := 0; j < 10; j++ {
		sd += math.Pow(num[j]-mean, 2)
	}
	sd = math.Sqrt(sd / 10)
	fmt.Println("The Standard Deviation is: ", sd)
}
Input:****** Enter 10 elements *******
Enter 1 element: 3
Enter 2 element: 5
Enter 3 element: 9
Enter 4 element: 1
Enter 5 element: 8
Enter 6 element: 6
Enter 7 element: 58
Enter 8 element: 9
Enter 9 element: 4
Enter 10 element: 10
Output:The Standard Deviation is: 15.8117045254457
Execution Steps:Run the program.Enter 10 numbers as prompted.The program calculates the sum, then the mean, then the standard deviation.The final result is printed.Week 5: Write a Program in Go language to print Floyd's TriangleCode:package main

import "fmt"

func main() {
	var rows int
	var temp int = 1
	fmt.Print("Enter number of rows: ")
	fmt.Scan(&rows)
	for i := 1; i <= rows; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d ", temp)
			temp++
		}
		fmt.Println("")
	}
}
Input:Enter number of rows: 5
Output:1 
2 3 
4 5 6 
7 8 9 10 
11 12 13 14 15 
Execution Steps:Run the program.Enter the desired number of rows (e.g., 5).The program prints numbers sequentially, increasing the line length by one each time.Week 6: Write a Go Program to take user input and addition of two stringsCode:package main

import "fmt"

func main() {
	fmt.Print("Enter First String: ")
	var first string
	fmt.Scanln(&first)
	var second string
	fmt.Print("Enter Second String: ")
	fmt.Scanln(&second)
	fmt.Print(first + second)
}
Input:Enter First String: Go
Enter Second String: Programming
Output:GoProgramming
Execution Steps:Run the program.Enter the first string (e.g., Go) and press Enter.Enter the second string (e.g., Programming) and press Enter.The program prints the two strings concatenated together.Week 7: Write a Go Program to check whether a string is Palindrome or notNote: The code in the manual checks if a number is a palindrome, despite the experiment title referencing a string.Code:package main

import "fmt"

func main() {
	fmt.Println("Golang program to check for palindrome")
	var number, rem, temporary int
	var reverse int = 0
	
	number = 45454
	
	temporary = number

	for {
		rem = number % 10
		reverse = reverse*10 + rem
		number /= 10
		if number == 0 {
			break
		}
	}

	if temporary == reverse {
		fmt.Printf("Number %d is a Palindrome", temporary)
	} else {
		fmt.Printf("Number %d is not a Palindrome", temporary)
	}
}
Input:None. The number 45454 is hardcoded in the program.Output:Golang program to check for palindrome
Number 45454 is a Palindrome
Execution Steps:Run the program.The program reverses the hardcoded number (45454).It compares the original number to the reversed number.The result is printed.Week 8: Write a Go Program to Build a contact formThis experiment has two parts: the Go backend and the HTML template.Code (main.go):package main

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		_ = details
		tmpl.Execute(w, struct{ Success bool }{true})
	})
	http.ListenAndServe(":8080", nil)
}
Code (forms.html):{{if.Success}}
<h1>Thanks for your message!</h1>
{{else}}
<h1>Contact</h1>
<form method="POST">
<label>Email:</label><br />
<input type="text" name="email"><br />
<label>Subject:</label><br />
<input type="text" name="subject"><br />
<label>Message: </label><br />
<textarea name="message"></textarea><br />
 <input type="submit">
</form>
{{end}}
Input:Run go run main.go.Open http://localhost:8080 in a web browser.Fill the form:Email: test@example.comSubject: HelloMessage: This is a test.Click the "submit" button.Output:Initial Page (http://localhost:8080):<h1>Contact</h1>
<form method="POST">
<label>Email:</label><br />
<input type="text" name="email"><br />
<label>Subject:</label><br />
<input type="text" name="subject"><br />
<label>Message: </label><br />
<textarea name="message"></textarea><br />
 <input type="submit">
</form>
After Submission:<h1>Thanks for your message!</h1>
Execution Steps:Place forms.html in the same directory as main.go.Run the Go program (go run main.go), which starts a web server on port 8080.Open a web browser and navigate to http://localhost:8080.The server handles the GET request and executes the template's {{else}} block, displaying the contact form.Fill in the form and click Submit. This sends a POST request to the server.The server's handler detects the POST method, parses the form values, and re-executes the template.This time, struct{ Success bool }{true} is passed, causing the {{if .Success}} block to render, showing the "Thanks" message.Week 9: Write a GO Program to calculate average using arraysCode:package main

import "fmt"

func main() {
	var num [100]int
	var temp, sum, avg int
	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&temp)
	for i := 0; i < temp; i++ {
		fmt.Print("Enter the number: ")
		fmt.Scanln(&num[i])
		sum += num[i]
	}
	avg = sum / temp
	fmt.Printf("The Average of entered %d number(s) is %d", temp, avg)
}
Input:Enter number of elements: 5
Enter the number: 2
Enter the number: 65
Enter the number: 18
Enter the number: 59
Enter the number: 54
Output:The Average of entered 5 number(s) is 39
Execution Steps:Run the program.Enter the total count of numbers (e.g., 5).Enter each number when prompted.The program calculates and prints the average.Week 10: Write a GO Program to delete duplicate element in a given arrayCode:package main

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
Input:None. The string slice is hardcoded.Output:The given array of string is: [abc cde efg efg abc cde]

The array obtained after removing the duplicate entries is: [abc cde efg]
Execution Steps:Run the program.The main function calls removeDuplicates with a hardcoded slice.The function uses a map to track unique strings and returns a new slice.The original and new slices are printed.Week 11: Write a GO Program with an example of Array reverse sort Functions for integers and stringsCode:package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Interger Reverse Sort")
	num := []int{50, 90, 30, 10, 50}
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)
	fmt.Println()
	fmt.Println("String Reverse Sort")
	text := []string{"Japan", "UK", "Germany", "Australia", "Pakistan"}
	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)
}
Input:None. The integer and string slices are hardcoded.Output:Interger Reverse Sort
[90 50 50 30 10]

String Reverse Sort
[UK Pakistan Japan Germany Australia]
Execution Steps:Run the program.The program sorts a hardcoded integer slice in reverse (descending) order.The program sorts a hardcoded string slice in reverse (alphabetical) order.The results are printed.Week 12: Write a program comprising of Contains, Contains Any, Count and Equal Fold string functionsCode:package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Germany", "G"))
	fmt.Println(strings.ContainsAny("Germany", "g"))
	fmt.Println(strings.Contains("Germany", "Ger"))
	fmt.Println(strings.Contains("Germany", "ger"))
	fmt.Println(strings.Contains("Germany", "er"))
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.EqualFold("Cat", "cat"))
	fmt.Println(strings.EqualFold("India", "Indiana"))
}
Input:None. All values are hardcoded in the fmt.Println calls.Output:true
false
true
false
true
3
true
false
Execution Steps:Run the program.The program executes and prints the result of eight different strings package functions.
