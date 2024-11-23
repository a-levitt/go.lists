package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5) // type, empty slots, capacity(~max size)
	userNames[0] = "Jenifer"
	userNames[1] = "Julia"
	userNames = append(userNames, "John", "Jane")
	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3) // type, capacity
	courseRatings["Go"] = 4.7
	courseRatings["React"] = 4.8
	courseRatings["Vue"] = 4.7
	fmt.Println(courseRatings)
}
