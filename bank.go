package main
import "fmt"

func main() {

	var accountBalance float64 = 1000

	fmt.Println("Welcome to bank")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	fmt.Println("Your choice is: ",choice)

	if choice == 1 {
		fmt.Println("Your balance is:",accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greathre than 0")
			return
		}
		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount: ",accountBalance)
	} else if choice == 3{
		fmt.Print("Enter amount to withdraw: ")
		var amountWithdraw float64
		fmt.Scan(&amountWithdraw)
		accountBalance -= amountWithdraw
		fmt.Println("Balance updated! new amount: ",accountBalance)
	} else {
		fmt.Print("Goodbye!")
	}


}