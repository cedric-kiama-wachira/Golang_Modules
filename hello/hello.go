package main

import (
    "fmt"
    // Use the functions in the standard library's log package to output error information.
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    // Configure the log package to print the command name ("greetings: ") at the start of its log messages, without a time stamp or source file information.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    // Look for a non-nil error value. There's no sense continuing in this case.
    if err != nil {
	// Use the functions in the standard library's log package to output error information. If you get an error, you use the log package's Fatal function to print the error and stop the program.    
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
