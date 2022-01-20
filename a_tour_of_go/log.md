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

### Flow control

**for loop**

```go
// basic for loop
sum := 0
for i:=0; i < 10; i++{
    sum += i
}

// "while"
sum := 1
for sum < 1000{
    sum += sum
}

// endless
for{
}
```

**if**

```go
// basic: no parentheses
if x < 0{
    // ...
}

// if with statements
if v:=math.Pow(x, n); v<lim{
    // ...
} else{
    
}
```

**switch**

```go
// no need for breaks, constants...
switch os:=runtime.GOOS; os{
  case "darwin":
      // ...
  case "linux":
      // ...
  default:
      // ...
}

// switch true => long if-else chain
switch{
  case x < 10:
      // ...
  case x = 11:
      // ...
  default:
      // ...
}
```

**defer**

> A defer statement defers the execution of a function until the surrounding function returns.
>
> The deferred call's **arguments are evaluated immediately**, but the function call is not executed until the surrounding function returns.

```go
// can be stacked
package main
import "fmt"
func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
```

**exercise: loops and functions**
```go
package main
import (
	"fmt"
	"math"
)
func Sqrt(x float64) float64 {
    z := 1.0
	  for math.Abs((z*z-x)/(2*z)) > float64(10e-6){
	      z -= (z*z-x)/(2*z)
		    fmt.Println(z)
	  }
	  return z
}
func main() {
	  fmt.Println(Sqrt(2))
}
```

### More types

**pointers**

has no pointer arithmetic and others are like pointers in C

**structs**

```go
// basic struct
type Vertex struct{
    X int
    Y int
}

// initialize and access
func main(){
    v:=Vertex{1,2}
    v.X  =4;
    // ...
    p := &v
    p.Y = 0 // no need to use (*P).Y
    
    // more ways to initialize
    v2 = Vertex{X: 1}  // Y=0
    v3 = Vertex{}
}
```

**arrays**

```go
// initialize
var a [10]int
a[0] = 1
a[1] = 2
primes := [6]int{1,2,3,4,5,6}
```

**slices**

```go
// slice an array
// []T
var s[]int = primes[1:4] // 2,3,4
```

A slice is a reference of a part of array.

```go
// default slices
var a [10]int
// equivalents
a[0:10]
a[:10]
a[0:]
a[:]
```

Slice has capacity and size (like c++'s vector)

```go
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	// Extend its length.
	s = s[:4]
	printSlice(s)
	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}
```

use make to dynamicly create a slice

```go
a := make([]int, 5)  // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```

**range**

range loop:

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// more simple ways
// for i, _ := range pow
// for _, value := range pow
```

**exercises: slices**

```go
func Pic(dx, dy int) [][]uint8 {
   image := make([][]uint8, dy)
	for i:=0;i<dy;i++{
	    image[i] = make([]uint8, dx)
	    for j:=0;j<dx;j++{
		    image[i][j] = uint8(i^j)
		}
	}
	return image
}
```

**Reference**

[Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)

[Go Slices: usage and internals](https://go.dev/blog/slices-intro)

## Methods and interfaces

## Concurrency