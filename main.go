package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5) // type, empty slots, capacity(~max size)
	userNames[0] = "Jenifer"
	userNames[1] = "Julia"
	userNames = append(userNames, "John", "Jane")
	fmt.Println(userNames)

	courseRatings := make(floatMap, 3) // type, capacity
	courseRatings["Go"] = 4.7
	courseRatings["React"] = 4.8
	courseRatings["Vue"] = 4.7
	courseRatings.output()
	// fmt.Println(courseRatings)
}
