package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	Message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return Message
}
