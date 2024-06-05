package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "Hello, World!"

	fmt.Println("--------------------")

	// Accessing the string
	fmt.Println("Message:", message)
	fmt.Println("Length:", len(message))

	fmt.Println("--------------------")

	// Accessing the string by index
	fmt.Println("Byte:", message[0])
	fmt.Println("ByteToString:", string(message[0]))

	fmt.Println("--------------------")

	// Slicing the string
	fmt.Println("Slice:", message[0:5])
	fmt.Println("Slice:", message[7:])

	fmt.Println("Contains 'World':", strings.Contains(message, "World"))
	fmt.Println("Contains 'Go':", strings.Contains(message, "Go"))

	fmt.Println("--------------------")

	// Concatenation
	fmt.Println("Concatenation:", message+" Welcome to Go!")
	message = strings.Join([]string{message, "Welcome to Go!"}, " ")
	fmt.Println("Concatenation with strings pkg:", message)

	fmt.Println("--------------------")

	// Iterating over the string with unicode character
	message = "Hello!⌘"

	fmt.Println("Iterating over the string:")
	for i, char := range message {
		fmt.Printf("index: %d, string: %s, quoted_string: %+q, rune: %v, byte: %x, length: %d\n", i, string(char), char, char, char, len(string(char)))
	}

	fmt.Println("--------------------")

	// Rune to string
	fmt.Printf("RuneToString: %q %U\n", 72, 72)
	fmt.Printf("RuneToString: %q %U\n", 108, 108)
	fmt.Printf("RuneToString: %q %U\n", 8984, 8984)

	fmt.Println("--------------------")

	// Unicode characters
	fmt.Printf("ByteToString: %q\n", '\u006c')
	fmt.Printf("ByteToString: %q\n", '\u006f')

	fmt.Println("--------------------")

	// Comparison
	message1 := "Hello, World!"
	message2 := "Hello, world!"

	// Case-sensitive comparison
	fmt.Println("Equal:", message1 == message2)
	// Case-insensitive comparison
	fmt.Println("EqualFold:", strings.EqualFold(message1, message2))

	fmt.Println("--------------------")

	// Prefix and Suffix
	fmt.Println("Prefix 'Hello':", strings.HasPrefix(message1, "Hello"))
	fmt.Println("Prefix 'World':", strings.HasPrefix(message1, "World"))
	fmt.Println("Suffix 'World!':", strings.HasSuffix(message1, "World!"))

	fmt.Println("--------------------")

	// Searching
	fmt.Println("Index 'World':", strings.Index(message1, "World"))
	fmt.Println("Index 'Go':", strings.Index(message1, "Go"))

	fmt.Println("Contains 'World':", strings.Contains(message1, "World"))
	fmt.Println("Contains 'Go':", strings.Contains(message1, "Go"))

	message1 = strings.ReplaceAll(message1, "o", "R")
	fmt.Println("ReplaceAll o to R:", message1)

	fmt.Println("--------------------")
}
