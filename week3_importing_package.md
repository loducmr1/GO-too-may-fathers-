# Week 3

## Steps

1. **Create a folder on Desktop:**

   - Name the folder `newpack`.

2. **Create `square.go` file inside `newpack`:**

   - Open `square.go` in Notepad and paste the following code:

   ```go
   package newpack

   func Square(x int) int {
       return x * x
   }
   ```
3. Copy ```newpack``` folder to Go source directory:

    - Copy the entire newpack folder.

    - Paste it inside: C:\Program Files\Go\src

    - Verify that the folder is present there.

4. Create week3.go file in your personal folder on Desktop:

    - Paste the following code into week3.go:

    ```go
    package main

    import "fmt"
    import "newpack"

    func main() {
    result := newpack.Square(5)
    fmt.Println("Square of 5 is:", result)
    }
    ```
5. Run the program:

    - Open a terminal (Command Prompt or PowerShell).

    - Navigate to the folder containing week3.go.

    - Run the command:

    ```powershell
    go run week3.go
    ```

You should see the output:

```powershell
Square of 5 is: 25
