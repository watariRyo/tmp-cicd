package main

import "fmt"

var version string // ビルド時にldflagsで指定

func main() {
    fmt.Printf("Hello, World! Version: %s\n", version)
}