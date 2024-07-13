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
go get github.com/liudng/dogo
```

dogo.json

```json
{
  "WorkingDir": ".",
  "SourceDir": ["."],
  "SourceExt": [".c", ".cpp", ".go", ".h"],
  "BuildCmd": "go build run main.go",
  "RunCmd": "./main",
  "Decreasing": 1
}
```
