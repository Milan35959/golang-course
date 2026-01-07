package main

import "fmt"

const age = 30

var name = "Golang"

func main() {
	// const name string = "Golang"

	// name = "java"
	// const age int = 30

	// fmt.Println("Hello, ", name, "You are", age, "years old.")
	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println("Server is running at", host, "on port", port)
}
