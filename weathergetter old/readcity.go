package weathergetter

import "fmt"

// Read city from user
func ReadCity() string {
	fmt.Println("Enter city name. Remember to start with Capital letter.")
	var input string

	fmt.Scanln(&input) // Does not work with whitespaces. Horrible.
	fmt.Println("Your input was:", input)
	return input
}
