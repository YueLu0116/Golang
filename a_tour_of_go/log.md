# Logs for learning A Tour of Go

> [A Tour of Go](https://go.dev/tour/welcome/1)

## Basics

### Packages, variables, and functions

**Basic structure of Go project**

go program is made up of packages

```go
package main // current packages

import{
    // import other packages
    // ...
}

func main(){
    // ...
}
```

exported names of a package must begin with capital letter

**Functions**

```go
// *basic structure
func add(x int, y int) int {
	return x + y
}

// parameters of a same type:
func add(x, y int) int{
    return x + y
}

// like python, go functions can return multiple results
func swap(x, y string) (string, string){
    return y, x
}

// returning without arguments use the names in function define parts
func split(sum int)(x, y int){
    // x =..
    // y =..
    return
}
```

**Variables**

basic structure (can be defined in packages level or function level)

```go
var name1, name2, ... type
```

initialize

```go
var name1, name2 type = value1, value2
```

short assignment

```go
name1, name2 := value1, value2
```

**Types**

go basic types

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

default initialization

```
0 -> numeric
false -> bool
"" -> string
```

type conversion (different types of variables must be converted)

```
var a_var type = T(another_var)
```

type inference examples

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

**Constants**

`const`  keyword but cannot use `:=` syntax

```
const Pi = 3.14
const world = "World"
```

**References**

[Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)

## Methods and interfaces

## Concurrency