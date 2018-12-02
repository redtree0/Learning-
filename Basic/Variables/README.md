# Variables in Go

https://www.tutorialspoint.com/go/go_variables.htm

## Variable Definition in Go
go에서 변수 선언은 다음과 같다.
```
var variable_list optional_data_type
```

여기서 ```optional_data_type```은 go 데이터 타입(byte, int, float32, boolean, user-defined object etc)이다.
```variable_list```는 하나 혹은 여러개의 변수 명을 지정할 수 있다.

```go
var  i, j, k int;
var  c, ch byte;
var  f, salary float32;
d =  42;
```

## Static Type Declaration in Go
정적 타입 선언(Static Type Declaration)은 컴파일러에게 주어진 타입의 변수가 있다는 것을 알려주는 것이다. 변수 선언은 컴파일시에만 의미가 있으므로 컴파일러는 프로그램을 링크 할 때 실제 변수 선언을 필요로합니다.


## Dynamic Type Declaration / Type Inference in Go
동적 타입 선언(Dynamic Type Declaration)은 변수에 할당된 값에 따라 타입이 정의되는 것이다.

```go
package main

import "fmt"

func main() {
   var x float64 = 20.0

   y := 42 
   fmt.Println(x)
   fmt.Println(y)
   fmt.Printf("x is of type %T\n", x)
   fmt.Printf("y is of type %T\n", y)	
}
```

## Mixed Variable Declaration in Go
go의 타입 추론으로 서로 다른 타입의 변수들을 동시에 선언할 수 있습니다.

```go
package main

import "fmt"

func main() {
   var a, b, c = 3, 4, "foo"  
	
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)
}
```

## The lvalues and the rvalues in Go

- lvalue 
lvalue는 메모리 위치에 대한 표현 방식이다.

- rvalue 
rvalue는 어느 메모리 주소에 저장된 데이터 값을 의미한다.