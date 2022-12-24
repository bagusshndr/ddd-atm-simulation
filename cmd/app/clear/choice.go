package clear

import "fmt"

func Choice(choice int) int {

	fmt.Printf("\nEnter any option to be served!\n\n")
	fmt.Printf("1. Check Amount\n")
	fmt.Printf("2. Desposit\n")
	fmt.Printf("3. Withdraw\n")
	fmt.Scan(&choice)

	return choice
}
