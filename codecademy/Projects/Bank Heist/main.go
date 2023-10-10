package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGiards := rand.Intn(100)

	if eludedGiards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if openedVault >= 70 && isHeistOn == true {
		fmt.Println("Grab and GO!")
	} else if isHeistOn == true {
		isHeistOn = false
		fmt.Println("You didn't bring the right tools")
	}

	leftSafley := rand.Intn(5)
	if isHeistOn == true {
		switch leftSafley {
		case 0:
			isHeistOn = false
			fmt.Println("The cops showed up")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("You were working for The Joker ...")
		default:
			isHeistOn = true
			fmt.Println("You made it to the getaway car!")
		}
	}

	if isHeistOn == true {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("You managed to get away with £%v", amtStolen)
	}

	fmt.Printf("Is the heist on? %+v", isHeistOn)

}
