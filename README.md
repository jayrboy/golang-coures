# Golang CMD

https://go.dev/learn/
https://go.dev/play/

## Run Compiler

1. Install Extension `Go`
2. Ctrl > Alt > N

## Initialize GitHub

```sh
go mod init github.com/jayrboy/golang-coures
```

go.mod

## Code Snippet

pkgm

1. main.go Ctrl + Shift + P > Insert Snippet

fp - fmt.Println("")
ff - fmt.Printf("", var)

## Go Build

"source code" to "binary file"

```sh
go build main.go
```

.exe (เอาไปรันที่ไหนก็ได้)

## Go Get

https://github.com/liudng/dogo

```sh
go install github.com/liudng/dogo@latest
go get github.com/liudng/dogo
```

dogo.json

```json
{
  "WorkingDir": ".",
  "SourceDir": ["."],
  "SourceExt": [".c", ".cpp", ".go", ".h"],
  "BuildCmd": "go build main.go",
  "RunCmd": "./main",
  "Decreasing": 1
}
```

# Variable Types

1. Number

```go
package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000

func main() {
	fmt.Println("Hello Golang!")

	var msg string = "Hello "
	fmt.Println(msg + "World!")

	fmt.Printf("%d %d", numberInt, numberInt2) // printf("%s",variable) in c language

	numberFloat := 25.4
	fmt.Println(numberFloat)

	fmt.Println(numberInt + numberInt2)
	fmt.Println(float64(numberInt) + numberFloat)

	fmt.Println("My money = ", numberInt)
}
```

Output: Number

```c
Hello Golang!
Hello World!
1000 200025.4
3000
1025.4
My money =  1000
```

2. String
3. Boolean
4. Array
5. Slice
6. Struct
7. Pointer
8. Function
9. Interface (Class the same a c#)
10. Map
11. Channel
