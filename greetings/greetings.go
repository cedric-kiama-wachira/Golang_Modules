package greetings

import (
    // Imported the GO standard errors package 
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
// Updated by changing the function so that it returns two values: a string and an error. My caller will check the second value to see if an error occurred.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    // Add an if statement to check for an invalid request (an empty string where the name should be) and return an error if the request is invalid. The errors.New function returns an error with my message inside.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    // Added nil (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded.
    return message, nil
}
