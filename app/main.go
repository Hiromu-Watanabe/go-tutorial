package main

import "fmt"
import "rsc.io/quote" // about quote => https://github.com/rsc/quote/blob/v4.0.1/quote.go#L22

func main() {
    fmt.Println(quote.Go())
}