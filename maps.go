package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
		"Go":       "https://golang.org"}

	fmt.Println(websites["Go"])
	websites["Amazon Web Services"] = "https://aws.com"
	fmt.Println(websites)
}
