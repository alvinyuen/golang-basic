# Learning golang

1. Statically typed language where all variables have assigned types. You can set types explicitly or implicitly

* Explicit
```
var aInteger int = 42
var aString string = "some text"
```

* Implicit
```
anInteger := 42
aString := "This is Go!"
```

2. Like C, Go supports the use of pointers, variables that store the address of another value
	var p *int
	var v int = 42

3. The struct type in Go is a data structure, similar to C's struct and Java's classes. They encapsulate both data and methods, but unlike Java, Go doesn't have an inheritance model. You don't have concepts like super or sub-structs. Each structure is independent, with its own fields for data management and optionally its own methods

Notes
..* Go has runtime that's statically linked into each compiled application. Along with any external packages that are declared and imported in the application source code. The runtime operates in its own background threads, So like other managed languages such as Java & C#, memory is managed completely in the background. And you don't have to explicitly allocate or deallocate memory.

## The new() function
* Allocates but does not initialize memory
* Results in zeroed storage but returns a memory address

## The make() function
* Allocates and initilizes memory
* Allocates non-zeroed storage and returns a memory address

Use make() to allocate and initialize memory
```
m := make(map[string]int)
m["key"] = 42
fmt.Println(m)
```

| desc | cmd |
|---| --- |
| run file | go run <file>.go |
| format | gofmt -w <file>.go |
| build (build a executable) | go build <file>.go |

