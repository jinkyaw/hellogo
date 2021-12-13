package main

import (
    "fmt"
    "github.com/jinkyaw/greetingsgo" 
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("User")
    fmt.Println(message)
}
