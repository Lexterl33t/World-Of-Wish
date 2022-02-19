package main

import (
	"WorldOfSwish/worldofwish"
	"fmt"
)

func banner() {
	fmt.Println("[ Welcome to World Of Wish ]")
	fmt.Println("1) Create Human")
	fmt.Println("2) Create Orc")
	fmt.Println("3) Quit")

	fmt.Println("")
}

func startAttack(you *worldofwish.Races_t, enemy *worldofwish.Races_t) {
	for {
		if you.GetVie() == 0 {
			fmt.Println(enemy.GetName(), "is winner !!")
			break
		} else if enemy.GetVie() == 0 {
			fmt.Println(you.GetName(), "is winner !!")
			break
		} else {
			you.Attack(enemy)
			enemy.Attack(you)
			continue
		}
	}
	return
}

func main() {
	for {
		choice := 0
		banner()
		fmt.Print("> ")
		fmt.Scan(&choice)

		if choice == 1 {

			you := worldofwish.New().NewHuman("Bryton", 80, 12)
			enemy := worldofwish.New().NewOrc("Cedric", 85, 8)

			startAttack(you, enemy)

		} else if choice == 2 {

			you := worldofwish.New().NewOrc("Cedric", 85, 8)
			enemy := worldofwish.New().NewHuman("Bryton", 80, 12)

			startAttack(you, enemy)

		} else if choice == 3 {

			fmt.Println("Goodbye :)")
			break

		} else {

			fmt.Println("Unknow command")
			continue

		}
	}
}
