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

**map**

```go
// basic usages
var m map[string]Vertex

func main(){
    m = make(map[string]Vertex)
    m["hh"] = Vertex{1.0,1.0}
}

// literals
var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}

// detect if a key exists
elem, ok := m[key]
```

**exercises: map**

```go
func WordCount(s string) map[string]int {
    var m map[string]int
    m = make(map[string]int)
    for _, w := range strings.Fields(s){
        m[w] += 1
    }
    return m
}
```

**function values**

```go
// basics: can be used as function arguements and return values
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}
func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))
}

// closures: capture variables from outside func's body
// adder() return a closure
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}

```

**exercises: fibonacci closure**

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    id := 0
    return func() int{
        var ret int
        if id==0 || id == 1{
          ret = id
        } else{
          ret = (id-1) + (id-2)
        }
        id += 1
        return ret
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
```

**Reference**

[Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)

[Go Slices: usage and internals](https://go.dev/blog/slices-intro)

## Methods and interfaces

**methods**: functions with a specific receiver of on a type

```go
// methods for structs
type Vertex struct {
	X, Y float64
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods on other user defined types
type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
```

We cannot define methods on types in other packages (int, for example)

**pointer receivers**: can modify the value to which the receiver points (more common)

 ```go
 type Vertex struct {
 	X, Y float64
 }
 func (v *Vertex) Scale(f float64) {
 	v.X = v.X * f
 	v.Y = v.Y * f
 }
 
 // methods with pointer receivers take either a value or a pointer as the receiver when they are called
 type Vertex struct {
 	X, Y float64
 }
 func (v *Vertex) Scale(f float64) {
 	v.X = v.X * f
 	v.Y = v.Y * f
 }
 func main() {
 	v := Vertex{3, 4}
 	v.Scale(2)
 	p := &Vertex{4, 3}
 	p.Scale(3)
 }
 
 // methods with value receivers take either a value or a pointer as the receiver when they are called
 type Vertex struct {
 	X, Y float64
 }
 func (v Vertex) Abs() float64 {
 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
 }
 func main() {
 	v := Vertex{3, 4}
 	fmt.Println(v.Abs())
 	p := &Vertex{4, 3}
 	fmt.Println(p.Abs())
 }
 ```

**interface**

```go
// basic struct
type Abser interface{
    Abs() float64
}
type Vertex struct {
	X, Y float64
}
func (v *Vertex) Abs() float64 {  // type Vertex implements interface Abser
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// Note: methods defined on value cannot be used for interface pointers
func main(){
    var a Abser
    a = v
    fmt.Println(a.Abs())  // wrong
}
// receivers can be nil
type I interface {
	M()
}
type T struct {
	S string
}
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
func main() {
	var i I
	var t *T  // nil
	i = t
	describe(i)
	i.M()
}
// But interface itself cannot be nil

// Empty interface: without any method. Like void* in C, empty interface can hold any type
```

**type assertions**

interface-> (value, type)

use type assertion to access the underlying concrete value

```go
// t, ok := i.(T)
func main() {
	var i interface{} = "hello"
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	f = i.(float64) // panic
	fmt.Println(f)
}

// type switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
```

**exercise: Stringers**

```go
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
```

**other common interfaces**

Error, Reader, and Image (skip)

## Concurrency

**goroutines**

use `go` keyword

```go
func say(s string) {
	// ...
}

func main() {
	go say("world")
	say("hello")
}
```

**channels**

use `<-` channel operator

```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
// must be initialized using make
ch := make(chan int)
```

**buffered channels**

```go
ch := make(chan int, 100)
// when the buffer is full, sending to it will be blocked
// when the buffer is empty, receiving from it will be blocked
```

**close a channel**

1. Only the sender can close the channel to tell receivers there are no data to be sent
2. use `for i:=range c` to receive data from a channel until it is closed
3. Simple test: `v,ok := <-ch`

example:

```go
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```

**select**

waiting for multiple `cases` until as least one can be run. Randomly select one case and execute it. example:

```go
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// default selection: is run if no other cases matched
```

**mutex**

[`sync.Mutex`](https://golang.org/pkg/sync/#Mutex)