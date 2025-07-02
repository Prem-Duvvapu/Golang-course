# Golang-course

1. It is a compiled language
2. Go tool can run file directly, without VM
3. Executables are different for OS
4. System apps to web apps - Cloud
5. Don't bring baggage from other languages
6. object oriented(yes and no)
7. lexer does a lot of work

## 1. [hello](01hello)
Basic "Hello, World!" program in Go. Introduction to the Go workspace and basic syntax.

## 2. [variables](02variables)
Covers declaration, initialization, and types of variables in Go, including shorthand syntax.

## 3. [user input](03userinput)
Demonstrates how to read user input using `bufio` and handle basic I/O operations.

## 4. [conversion](04conversion)
Shows how to convert between different types (e.g., int to string, float to int) in Go.

## 5. [mytime](05mytime)
Working with the `time` package to get current time, formatting, duration, and sleep.

## 7. [pointers](07mypointers)
Covers concepts of memory addresses, pointers, and pointer manipulation.
n
## 8. [myarrays](08myarrays)
Understanding arrays in Go, how to declare, initialize, and iterate through them.

## 9. [myslices](09myslices)
Explores slices, slicing arrays, capacity, appending, and slice operations.

## 10. [mymaps](10mymaps)
Working with maps (dictionaries) in Go, adding, updating, and deleting key-value pairs.

## 11. [mystructs](11mystructs)
Introduction to structs, struct fields.

## 12. [ifelse](12ifelse)
Conditional statements in Go including if, else if, and else blocks.

## 13. [switchcase](13switchcase)
Using switch statements with different cases, including expression-less switches.

## 14. [loops](14loops)
Looping in Go using `for` loops (as Go does not have while or do-while).

## 15. [functions](15functions)
Defining and calling functions, variadic functions, and returning multiple values.

## 16. [methods](16methods)
Defining methods on structs and understanding the difference between functions and methods.

## 17. [defer](17defer)
Understanding `defer` statements and their use in resource cleanup.

## 18. [files](18files)
Reading from and writing to files using the `os` and `io` packages.

## 19. [webrequests](19webrequests)
Making HTTP GET/POST requests using the `net/http` package.

## 20. [urls](20urls)
Working with and parsing URLs using the `net/url` package.

## 21. [web verb requests](21webverbrequests)

## 22. [my json][22myjson]

## 23. [23mymodules](23mymodules)
    some import commands
    go mod init moduleName
    go mod tidy
    go mod verify
    go list
    go list all
    go list -m all
    go mod why module(ex: github.com/gorilla/mux)
    go mod edit -go 1.19 (to change go version in go.mod)
    go mod vendor
    go run -mod=vendor main.go