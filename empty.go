package main

import "fmt"

type Anything interface{}

func AcceptAnything(anything Anything) {
	fmt.Println(anything)
}

func main() {
	AcceptAnything(true)
}
