package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5) // type, empty slots, capacity(~max size)
	userNames[0] = "Jenifer"
	userNames[1] = "Julia"

	userNames = append(userNames, "John", "Jane")

	fmt.Println(userNames)
}
