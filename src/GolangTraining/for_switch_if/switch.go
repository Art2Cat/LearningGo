package main

import "fmt"

func main() {
	normalSwitch()
	otherSwitch()
}

func normalSwitch() {
	switch "m" {
	case "t":
		fmt.Println("t")
	case "m":
		fmt.Println("m")
	case "y":
		fmt.Println("y")
	}
}

func otherSwitch() {
	var x string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&x)
	switch x {
	case "Yiming":
		fmt.Println("hello", x)
		fallthrough
	case "world":
		fmt.Println("hello", "world")
	case "you":
		fmt.Println("hello", "you")
		fallthrough
	case "fuck":
		fmt.Println("fuck", "world")
	}
}

func multiple_evalsSwitch()  {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}
}

func no_expression_switch()  {
	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Wassup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")
	}
}