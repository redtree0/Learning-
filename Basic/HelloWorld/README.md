# Hello World

[Reference]

https://www.tutorialspoint.com/go/go_environment.htm
https://www.tutorialspoint.com/go/go_program_structure.htm

```go
package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}
```

## Tokens in Go
[Reference]
https://www.tutorialspoint.com/go/go_basic_syntax.htm

Go 프로그램은 토큰 단위로 인식한다. 여기서 토큰은 키워드, 식별자, 상수, 문자열, 심볼 등이다. 예를 들면, print문은 6개의 토큰으로 구성되어 있다.

```go
fmt.Println("Hello, World!")
```

```go
fmt // 1
.   // 2
Println // 3 
( // 4
   "Hello, World!" // 5
) // 6 
```

## Line Separator
In a Go program, the line separator key is a statement terminator. 
```go
fmt.Println("Hello, World!")
fmt.Println("I am in Go Programming World!")
```
