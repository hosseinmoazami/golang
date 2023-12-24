//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
		fmt.Println("Hit to Max Health")
	}
}

func (player *Player) addEnergy(amount uint) {
	player.energy += amount
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
		fmt.Println("Hit to Max Energy")
	}
}

func main() {
	hossein := Player{
		"Hossein",
		100, 100,
		100, 100,
	}

	fmt.Println("Before fight", hossein)

	hossein.addHealth(10)
	hossein.addEnergy(95)

	fmt.Println("After fight", hossein)

}
