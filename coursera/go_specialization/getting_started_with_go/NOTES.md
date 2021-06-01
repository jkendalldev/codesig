## Week 1
- Week 1 Notes

# GoLang Notes
## _What a cool language!_

## Advantages of Golang
- Code Runs Fast
- Has Garbage Collection
- Simpler objects

## More Notes
- Go is a complied language, but also have Garbabe Collection
- This means the GC can automatically manage memory for you!
- Manual memory management is hard!
- GC is typically only done by interpreters, but go is compiled and does it

## Objects in Go
- Go is object oriented, but it's weakly object-oriented
- Group together data and functions that are related to each other (create user defined types, not just int, float, etc.) , this is object oriented programming.
- Think about an int type, each integer has a type of data and a set of functions you can apply to the data. Object oriented is the same idea, but you can create
your own types (classes).
- Go does not use the term "class"
- Go uses structs with associated methods
- struct ends up being what you would call a class
- But no generics, inheritence, or constructors though with structs

## Concurrency in Go
- This is a big advantage of go
- Parallelism
- Number of cores still increases over time
- Multiple tasks may be performed at the same time
- on different cores
- Difficulties with parallelism
- When do tasks start/stop?
- What if one task needs data from another task?
- Do tasks conflict in memory?
- Concurrent Programming
  - Concurrency is the management of multiple tasks at the same time
  - Key requirement for large systems
  - Concurrent Programming enables parallelism
  - Management of task execution
  - Communication between tasks
  - Synchronization between tasks
  - Go includes concurrency primitives
  - Goroutines represent concurrent tasks
  - Channels are used to communicate between tasks
  - Select enables task synchronization

## Workspaces and Packages
- Workspaces
- Hierarchy of directories
- Common organization is good for sharing
- People can more easily work together
- Make it easy for people to work with your code
- Easier to share
- Workspace (below hierarchy is recommended, not enforced)
 - src (source code)
 - pkg (packages, libraries)
 - bin (executables) 
 - Programmer typically has one workspace for many projects
 - Workspace directory defined by GOPATH env variable
 - Go tools assume your code is in GOPATH
 - Packages 
   - Group of related source files
   - Each package can be imported by other packages
   - First line of file names the package
   - package one
   - package two
   - import ("one" "two")
   - Package "Main" .. must have a main package
   - Building the main package generates an executable
   - Main package must have main() function
   
   package main
   import "fmt"
   funct main () {
       fmt.Printf(hello, world\n")
   }

## Go Tool
- A tool to manage Go source code
- Includes some tools
- "go build" - compiles the program
- Creates an exe for the main package
- "go doc" -> shows docs for a package
- "go fmt" -> formats src code files
- "go get" -> downloads packages and installs them
- "go list" -> lists all installed packages
- "go run" -> compiles .go files and runs the exe
- "go test" -> runs tests
- import keyword is used to access other packages
- Go standard library includes many packages (i.e. fmt)
- Go Tool finds the packages you tell it to import
- Searches directories specified by GOROOT and GOPATH

## Variables
- Naming
- Case sensitve
- Don't use keywords
- Start with a letter
- Variables must have a name and a type

``` var x int ```
- Variable Types
- examples: integer, floating point, string
- Floating Point
- Can have fractional (decimal) values
- Floating poitn math (may use different hardware)
- Strings
- Byte (character) sequences
- type specifies what data and what operations
- Can define an alias (alternate name) for a type
- May improve clarity
- type Celsius float64 (Celsius is an alias for float64)
- type IDnum int (again, IDnum is an alias)
- Can now declar variables using the type alias
- Ex. var temp Celsius
- Initialize in the declaration
- var x int = 100 (here we are specifying the type, no guessing!)
- var x = 100 (go will infer int type here, we are not being specific)
- Initialize after the declaration
- var x int, then x = 100 on next line
- Uninitialized variables have a zero value
- var x int // x=0
- var x string // x=""
- Short Variable Declaration
- Can perform a declaration and initialization together 
- with the := operator
- x := 100
- Variable is declared as type of expression on the right hand side
- Can only do this ^ inside a function
- 

## Hello, world! Exercise
- Submitted



## Week 2
- How to use, declare data types
- What functions there are 
- Pointers
- Pointer - An address to data in memory
- 2 main pointer operators..
- & operator returns the address of a variable/function
- \* operator returns data at an address (dereferencing)
- 


