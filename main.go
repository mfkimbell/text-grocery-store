package main

import (
	"fmt"
	"strconv"
)

var userInfo = []string{}
var firstN = ""
var lastN = ""
var email = ""
var userPurchases = []string{}
var purchaseLog string
var purchaseCount = 0
var inv = make(map[string]int)
var cart = make(map[string]int)
var prices = make(map[string]float32)

// cd desktop/go/checkout
// go run main.go
func main() {
	//user information
	firstN = updateFirst()
	lastN = updateLast()
	email = updateEmail()
	userInfo = append(userInfo, firstN)
	userInfo = append(userInfo, lastN)
	userInfo = append(userInfo, email)

	//fmt.Print(userInfo[2]) this would call the email

	inv["Carrots"] = 12
	inv["Tomatoes"] = 5
	inv["Onions"] = 9
	inv["Apples"] = 10

	prices["Carrots"] = 0.25
	prices["Tomatoes"] = 0.75
	prices["Onions"] = 1.25
	prices["Apples"] = 0.50

	fmt.Println("Map: ", inv)

	var name = "GO shopping application"

	fmt.Print("\n")
	fmt.Printf("Welcome %v to the %v \n \n", firstN, name)
	menu()

}

func updateFirst() string {

	var fn = ""
	fmt.Println("Please enter your first name:")
	fmt.Scan(&fn)
	return fn
}
func updateLast() string {
	var ln = ""
	fmt.Println("Please enter your last name:")
	fmt.Scan(&ln)
	return ln
}
func updateEmail() string {
	var e = ""
	fmt.Println("Please enter your email:")
	fmt.Scan(&e)
	return e

}
func menu() {
	for k := range cart {
		delete(cart, k)
	}
	userPurchases = nil
	var menuInput int
	var flag1 = true

	for flag1 {

		fmt.Println("\nPlease select one of the following menu options:")
		fmt.Println("(1) Make a purchase")
		fmt.Println("(2) Check available inventory")
		fmt.Println("(3) Display purchase history")
		fmt.Println("(4) Update user information")
		fmt.Println("(5) Quit")

		fmt.Scan(&menuInput)

		if menuInput == 1 {
			fmt.Println("")
			makePurchase()
			flag1 = false

		} else if menuInput == 2 {
			displayInv()
			menu()
			flag1 = false

		} else if menuInput == 3 {
			purchaseHistory()
			flag1 = false

		} else if menuInput == 4 {
			updateFirst()
			updateLast()
			updateEmail()
			menu()
			flag1 = false

		} else if menuInput == 5 {
			fmt.Print("Goodbye!\n")
			break
		} else {
			fmt.Println("Invalid input. Please select one of the valid options.")
		}

	}

}
func displayInv() {
	fmt.Println("\nCurrent Inventory:")
	fmt.Printf("Carrots $0.25 (%v in stock) \n", inv["Carrots"])
	fmt.Printf("Tomatoes $0.75 (%v in stock) \n", inv["Tomatoes"])
	fmt.Printf("Onions $1.25 (%v in stock) \n", inv["Onions"])
	fmt.Printf("Apples $0.50 (%v in stock) \n", inv["Apples"])
}

func makePurchase() {

	var input2 int
	var input3 int
	var flag1 = true

	for flag1 {

		fmt.Print("What would you like to purchase? ")
		fmt.Println("Please select one of the following menu options:")
		fmt.Printf("(1) Carrots $0.25 (%v in stock) \n", inv["Carrots"])
		fmt.Printf("(2) Tomatoes $0.75 (%v in stock) \n", inv["Tomatoes"])
		fmt.Printf("(3) Onions $1.25 (%v in stock) \n", inv["Onions"])
		fmt.Printf("(4) Apples $0.50 (%v in stock) \n", inv["Apples"])

		fmt.Scan(&input2)

		if input2 == 1 {
			var item = "Carrots"
			fmt.Printf("\nHow many %v\n", item)
			fmt.Scan(&input3)
			if validPurchase(input3, item) {
				commitPurchase(input3, item)
			} else {
				fmt.Printf("Cannot purchase %v %v with %v in stock.\n\n", input3, item, inv[item])
				makePurchase()
			}
			flag1 = false

		} else if input2 == 2 {
			var item = "Tomatoes"
			fmt.Printf("\nHow many %v\n", item)
			fmt.Scan(&input3)
			if validPurchase(input3, item) {
				commitPurchase(input3, item)
			} else {
				fmt.Printf("Cannot purchase %v %v with %v in stock.\n\n", input3, item, inv[item])
				makePurchase()
			}
			flag1 = false

		} else if input2 == 3 {
			var item = "Onions"
			fmt.Printf("\nHow many %v\n", item)
			fmt.Scan(&input3)
			if validPurchase(input3, item) {
				commitPurchase(input3, item)
			} else {
				fmt.Printf("Cannot purchase %v %v with %v in stock.\n\n", input3, item, inv[item])
				makePurchase()
			}
			flag1 = false

		} else if input2 == 4 {
			var item = "Apples"
			fmt.Printf("\nHow many %v\n", item)
			fmt.Scan(&input3)
			if validPurchase(input3, item) {
				commitPurchase(input3, item)
			} else {
				fmt.Printf("\nCannot purchase %v %v with %v in stock.\n\n", input3, item, inv[item])
				makePurchase()
			}
			flag1 = false

		} else {
			fmt.Println("\nInvalid input. Please select one of the valid options.")
			fmt.Println("")
		}
	}

}

func commitPurchase(amount int, veg string) {
	var input4 string
	//var amountS = strconv.Itoa(amount) if you want string version
	fmt.Printf("\n%v %v added to cart. ($", amount, veg)
	var amountF = float32(amount)
	fmt.Print(prices[veg] * amountF)
	fmt.Print(")")

	cart[veg] = cart[veg] + amount
	inv[veg] = inv[veg] - amount

	showCart()
	fmt.Println("Which would you like to do?\n(1) Make another purchase\n(2) Checkout")
	fmt.Scan(&input4)
	if input4 == "1" {
		makePurchase()
	} else if input4 == "2" {
		Checkout()
	} else {
		commitPurchase(amount, veg)
	}

}

func validPurchase(amount int, veg string) bool {
	if inv[veg]-amount < 0 {
		return false
	} else {
		return true
	}
}

func showCart() {
	fmt.Printf("\nYour current cart contains the following:\n")
	for key, ele := range cart {
		fmt.Printf("%v: %v \n", key, ele)

	}
	fmt.Println()
}

func Checkout() {
	var input5 string
	var total float32
	var purchase string

	fmt.Println("\nCurrent Cart")
	for key, ele := range cart {
		var ele2 = float32(ele)
		var ele3 = strconv.Itoa(ele)
		fmt.Printf("%v: %v \n", key, ele)
		purchase = key + ": " + ele3
		userPurchases = append(userPurchases, purchase)
		total += prices[key] * ele2
	}

	purchaseCount += 1
	var c2 = strconv.Itoa(purchaseCount)
	purchaseLog = purchaseLog + "\nCart for purchase: " + c2 + "\n"
	for _, ind := range userPurchases {
		purchaseLog += ind
		purchaseLog += "\n"
	}

	fmt.Printf("\nYour total is ($%v)\n", total)
	fmt.Printf("Would you like to receive an email confirmation at %v?\n", email)
	fmt.Printf("(1) Yes\n")
	fmt.Printf("(2) No\n")
	fmt.Printf("(3) Update user email\n")
	fmt.Scan(&input5)
	if input5 == "1" {
		fmt.Println("\nThank you for using my GO shopping application.\n\nReturning to menu...")
		menu()

	} else if input5 == "2" {
		fmt.Println("\nThank you for using my GO shopping application.\n\nReturning to menu...")
		menu()

	} else if input5 == "3" {
		updateEmail()
		Checkout()

	} else {
		fmt.Println("Invalid input. Please enter a number 1-3")
		Checkout()

	}

}

func purchaseHistory() {
	fmt.Print("\n")
	fmt.Println(purchaseLog)
	fmt.Println("\nReturning to menu...")
	menu()

}
