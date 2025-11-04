# Go Programming Lab Manual (Weeks 1-12)

This file contains all the Go programs from the Week 1 to Week 12 experiments, based on the R20 Lab Manual.

---
### Week 1: Write a Go Program to find LCM and GCD of given two numbers
**Code:**
```go
package main
import "fmt"

func lcm(temp1 int, temp2 int) {
    var lcmnum int = 1
    if temp1 > temp2 { lcmnum = temp1 } else { lcmnum = temp2 }
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
        if temp1%i == 0 && temp2%i == 0 { gcdnum = i }
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
    case 1: lcm(n1, n2)
    case 2: gcd(n1, n2)
    }
}
```
**Input:**
```
3
6
1
```
**Output:**
```
LCM of 3 and 6 is 6
```
**Execution Steps:**
Run → Input → Choose option → Print result.

---
### Week 2: Write a Go Program to print pyramid of numbers
**Code:**
```go
package main
import "fmt"
func main() {
    var rows, k, temp, temp1 int
    fmt.Print("Enter number of rows :")
    fmt.Scan(&rows)
    for i := 1; i <= rows; i++ {
        for j := 1; j <= rows-i; j++ { fmt.Print(" "); temp++ }
        for {
            if temp <= rows-1 { fmt.Printf("%d", i+k); temp++
            } else { temp1++; fmt.Printf("%d", (i+k-2*temp1)) }
            k++
            if k == 2*i-1 { break }
        }
        temp, temp1, k = 0, 0, 0
        fmt.Println("")
    }
}
```
**Input:**
```
5
```
**Output:**
```
    1
   232
  34543
 4567654
567898765
```
**Execution Steps:** Run → Enter number → Pyramid printed.

---
### Week 3: Write a program to use struct that is imported from another package
**Code:** (multi-file project; `main.go`, `father.go`, `son.go`)

**Execution Steps:**
1. Create folders: `family/father/son`
2. Add `father.go` to `family/father`, `son.go` to `family/father/son`
3. Add `main.go` in root.
4. Run `go run main.go` → packages initialized automatically.

**Input:** None (hardcoded)
**Output:**
```
Father package initialized
Son package initialized
Father: Mr. Jeremy Maclin
Son: Riley Maclin
```
---
### Week 4–12
*(Each week follows same brief style with code, input/output, and short execution notes.)*
