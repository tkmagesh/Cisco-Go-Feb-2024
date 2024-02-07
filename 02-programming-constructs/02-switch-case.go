package main

import "fmt"

func main() {
	/*
		rank by score
		0-3 => poor
		4-7 => not bad
		8-9 => good
		10 => excellent
		otherwise => invalid score
	*/

	// score := 6
	/*
		switch score := 6; score {
		case 0:
			fmt.Println("poor")
		case 1:
			fmt.Println("poor")
		case 2:
			fmt.Println("poor")
		case 3:
			fmt.Println("poor")
		case 4:
			fmt.Println("not bad")
		case 5:
			fmt.Println("not bad")
		case 6:
			fmt.Println("not bad")
		case 7:
			fmt.Println("not bad")
		case 8:
			fmt.Println("good")
		case 9:
			fmt.Println("good")
		case 10:
			fmt.Println("excellent")
		default:
			fmt.Println("invalid score")
		}
	*/

	switch score := 6; score {
	case 0, 1, 2, 3:
		fmt.Println("poor")
	case 4, 5, 6, 7:
		fmt.Println("not bad")
	case 8, 9:
		fmt.Println("good")
	case 10:
		fmt.Println("excellent")
	default:
		fmt.Println("invalid score")
	}

	// use switch case for nested if
	// no := 9
	switch no := 9; {
	case no%2 == 0:
		fmt.Printf("%d is divisible by 2\n", no)
	case no%3 == 0:
		fmt.Printf("%d is divisible by 3\n", no)
	}

	// using fallthrough
	switch n := 4; n {
	case 0:
		fmt.Println("n == 0")
		fallthrough
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
		// fallthrough
	case 7:
		fmt.Println("n <= 7")
	}

	// fallthrough applied
	switch subscription := "Premium"; subscription {
	case "Super":
		fmt.Println("[Super] : Licence for a family of 3")
		fallthrough
	case "Premium":
		fmt.Println("[Premium] : Private playlists")
		fallthrough
	case "Pro":
		fmt.Println("[Pro] : No ads")
		fallthrough
	case "Free":
		fmt.Println("[Free] : listen to songs")
	}
}
