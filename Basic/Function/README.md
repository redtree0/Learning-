# Function

## Defining a Function

[Reference]
https://www.tutorialspoint.com/go/go_functions.htm

- Func
- Function Name
- Parameters
- Return Type
- Function Body

```go
func mac(num1, num2 int) int {

    result int

    if ( num1 > num2 ) {
        result = num1
    } else {
        result = num2
    }
    return result 
}
```

## Calling a Function

```go
package main

import "fmt"

func main() {
    
    var a int = 100
    var b int = 200
    var ret int

    ret = max(a, b)

    fmt.Printf(" Max Value is : %d\n", ret)
}

func max(num1, num2 int) int {

    var result int

    if ( num1 > num2 ) {
        result = num1
    } else {
        result = num2
    }
    return result
}

```


## Returngin multiple values from Function

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("Hello", "Wolrd")
    fmt.Printf(a, b)
}

```

## Function Arguments

```formal parameters```

### Call by value

[Reference]
https://www.tutorialspoint.com/go/go_function_call_by_value.htm

The call by value method of passing arguments to a function copies the actual value of an argument into the formal parameter of the function. In this case, changes made to the parameter inside the function have no effect on the argument.

```go
func swap(int x, int y) int {
    var temp int

    temp = x
    x = y
    y = x

    return temp
}
```
### Call by reference

[Reference]
https://www.tutorialspoint.com/go/go_function_call_by_reference.htm

The call by reference method of passing arguments to a function copies the address of an argument into the formal parameter. Inside the function, the address is used to access the actual argument used in the call. This means that changes made to the parameter affect the passed argument.

```go
func swap(x *int, y *int) {
    var temp int
    temp = *x
    *x = *y
    *y = temp
}
```

## Function Usage

https://www.tutorialspoint.com/go/go_function_as_values.htm

### Functions as value

```go
package main

import ("fmt" "math")

func main() {

    getSquareRoot := func(x float64) float64 {
        return math.Sqrt(x)
    }

    fmt.Println(getSquareRoot(9))
}
```

### Function Closure

Go 언어는 익명 함수를 지원한다.
클로저는 함수와 그 함수가 선언됐을 때의 렉시컬 환경(Lexical environment)과의 조합이다.
렉시컬 환경이란 스코프는 함수를 호출할 때가 아니라 함수를 어디에 선언하였는지에 따라 결정된다.

```go
package main

import "fmt"

func getSequence() func() int {
    i:=0

    return func() int {
        i+=1
        return i
    }
}

func main() {
    nextNumber := getSequence()

    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())

    nextNumber1 := getSequence()
    fmt.Println(nextNumber1())
    fmt.Println(nextNumber1())
}
```

### Method

```go
package main

import (
    "fmt"
    "math"
)

type Circle struct {
    x,y,radius float64
}

func (circle Circle) area() float64 {
    return math.Pi * circle.radius * circle.radius
}

func main() {
    circle := Circle{x:0, y:0, radius:5}

    fmt.Printf("Circle area: %f", circle.area())
}
```


