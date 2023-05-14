# Programming with Google Go
    Module 1: Getting Started with Go
		M1.1.1 - Why Should I Learn Go? 
			Code runs fast
			Garbage Collection
		M1.1.2 - Objects
		M1.1.3 - Concurrency
		M1.2.1 - Installing Go
		M1.2.2 - Workspaces & Packages
		M1.2.3 - Go Tool
		M1.3.1 - Variables
		M1.3.2 - Variable Initialization
	Module 2: Basic Data Types
		M2.1.1 - Pointers
		M2.1.2 - Variable Scope
		M2.1.3 - Deallocating Memory
		M2.1.4 - Garbage Collection
		M2.2.1 - Comments, Printing, Integers
		M2.2.2 - Ints, Floats, Strings
		M2.2.3 - String Packages
		M2.3.1 - Constants
		M2.3.2 - Control Flow
		M2.3.3 - Control Flow, Scan
	Module 3: Composite Data Types
		M3.1.1 - Arrays
		M3.1.2 - Slices
		M3.2.1 - Hash Tables
		M3.2.2 - Maps
		M3.3.1 - Structs
	Module 4: Protocols and Formats
		M4.1.1 - RFCs
		M4.1.2 - JSON
		M4.2.1 - File Access, ioutil
		M4.2.2 - File Access, os
	
# Course 2: Functions, Methods, and Interfaces in Go
	Module 1: Functions and Organisations
		M1.1.1 - Why Use Functions?
		M1.1.2 - Function Parameters and Return Values
		M1.1.3 - Call by Value, Reference
		M1.1.4 - Passing Arrays and Slices
		M1.2.1 - Well-Written Functions
		M1.2.2 - Guidelines for Functions

# Module 2: Function Types
## M2.1.1 - First-Class Values

### Functions are First-class
Functions can be treated like other types
	• Variables can be declared with a function type
	• Can be created dynamically
	• Can be passed as arguments and returned as values
	• Can be stored in data structures

### Variables as Functions
	• Declare a variable as a func
	• Function is on right-hand side, without ()
```go
var funcVar func(int) int
func incFn(x int) int { 
   return x + 1
} 
func main() { 
   funcVar = incFn
   fmt.Print(funcVar(1))
}
``` 

Functions as Arguments
	
	- Function can be passed to another function as an argument

func applyIt(afunct func (int) int, val int) int { 
   return afunct(val)
} 

Example:
func applyIt(afunct func (int) int, val int) int { 
   return afunct(val)
} 

func incFn(x int) int {return x + 1}
func decFn(x int) int {return x - 1}

func main() { 
    fmt.Println(applyIt(incFn, 2))
    fmt.Println(applyIt(decFn, 2))
} 

Anonymous Functions
	- Don’t need to name a function
func applyIt(afunct func (int) int, val int) int { 
   return afunct(val)
} 

func main() { 
   v := applyIt(func (x int) int {return x + 1}, 2)
   fmt.Println(v)
} 

