package main
import "fmt"

func main() {

	var accountBalance float64 = 1000
	
	fmt.Println("Welcome to bank")

	for  {
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
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ",accountBalance)
		} else if choice == 3{
			fmt.Print("Enter amount to withdraw: ")
			var withrawalAmount float64
			fmt.Scan(&withrawalAmount)

			if withrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				return
			}
			
			if withrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have")
				return
			}

			accountBalance -= withrawalAmount
			fmt.Println("Balance updated! new amount: ",accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break;
		}
	}

	fmt.Println("Thanks for choosing our bank")

}