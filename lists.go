package main

import "fmt"

/*func main() {
	var productNames [4]string
	productNames = [4]string{"A book", "A T-shirt", "A notebook", "A desktop"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices[1])
	fmt.Println(productNames[1])

	// First element in slices is included, second is excluded
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}*/

/*func main() {
	prices := []float64{10.99, 8.99}
	prices[1] = 9.99
	newArr := prices[1:2]
	fmt.Println(newArr)

	fmt.Println(append(newArr, 5.99), prices)
}*/

func main() {
	hobbies := []string{"Programming", "Anime", "Twitch.tv"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	recrHobbies := hobbies[1:3]
	fmt.Println(recrHobbies)
	firstHobbies := hobbies[:2]
	fmt.Println(firstHobbies)
	firstHobbies = firstHobbies[1:3]
	fmt.Println(firstHobbies)
	firstHobbies[1] = "streams"
	firstHobbies = append(firstHobbies, "youtube")
	fmt.Println(firstHobbies)
}
