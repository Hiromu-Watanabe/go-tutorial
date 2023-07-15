package main

import (
    "fmt"
		"log"

    "github.com/Hiromu-Watanabe/go-tutorial/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // 名前の配列生成.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // 名前に対応する挨拶メッセージを受け取る
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(messages)
}