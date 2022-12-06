package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var alphabet = map[string]string{
		"ą": "a",
		"ć": "c",
		"ę": "e",
		"ł": "l",
		"ń": "n",
		"ó": "o",
		"ś": "s",
		"ź": "z",
		"ż": "z",
	}
	var new_message string = ""
	fmt.Print("Type your Polish message > ")
	in := bufio.NewReader(os.Stdin)
	message, _ := in.ReadString('\n')
	for _, x := range message {
		if alphabet[string(x)] != "" {
			new_message += alphabet[string(x)]
		} else {
			new_message += string(x)
		}
	}
	fmt.Println(new_message)
}
