# text-grocery-store

The purpose of this project was to familiarize myself with the Golang programming language. This project utilizes GO's maps, slices, for loops, flags, type conversions, user input scanner, and other fundamentals of GO. The program keeps track of the items and total cost of the customer's cart, as well as inventory left. The program also records the information about each time a customer checks out. 

**Interesting problems:**

I had to learn to clear a map with a for loop because I wanted to clear the customer's cart after each purchase so I could log each purchase they make. The "DisplayPurchases" option tells the customer everything they've previously purchased, but keeps the inventory available updated. Additionally, I logged the purchases with string concatenation of each purchase details. 

Currently, the program will not let the customer buy more inventory than is available, and if a customer chooses to buy the same item multiple times during the same cart session, it will add the current objects. Additionally, I made it only shows things being purchased in the cart and not the entire inventory, with things not purchased set to 0. 


Menu:

![](Screenshots/MenuDisplay.png)

Introduction/UserInput:

![](Screenshots/Introduction:UserInput.png)

First Checkout:

![](Screenshots/FirstCheckout.png)

Second Checkout (multiple different items):

![](Screenshots/SecondCheckout(multiple%20items).png)

Displaying Purchase History:

![](Screenshots/DisplayingPurchaseHistory.png)

Displaying Available Inventory:

![](Screenshots/DisplayingAvailableInventory.png)


