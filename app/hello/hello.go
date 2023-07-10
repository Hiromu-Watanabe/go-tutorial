package main

import (
    "fmt"

    "github.com/Hiromu-Watanabe/go-tutorial/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}