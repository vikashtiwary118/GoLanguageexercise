package main 


import "fmt"

func main() {
	age:=16
	switch age{
		case 16: fmt.Println("Your age is 16")

		case 18: fmt.Println("You are able to vote")

		case 20: fmt.Println("Your Age is 20")

		default: fmt.Println("You are not able for vote")
	}
}