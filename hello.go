package main

import "fmt"

func Hello(str string, lang string) string {
	if str == "" {
		return "Hello, World!"
	}

	if lang == "Spanish" {
		return "Hola, " + str
	}
	return "Hello, " + str + "!"
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
