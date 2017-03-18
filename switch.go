package main 

import "fmt"

func location(city string)(string, string) {

	var region string
	var continent string

	switch city{
	case "Los Angeles", "LA", "Santa Monica":
			region, continent = "California", "Nort America"

	case "Newyork", "NYC":
			region, continent = "Newyork", "Nort America"

	default:
			region, continent = "Unknown", "Unknown"
	}
	return region, continent
	
}

func main() {
	region, continent := location("istanbul")
	fmt.Printf("Matt Lives in %s, %s", region, continent)
}