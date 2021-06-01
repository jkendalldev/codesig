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

   