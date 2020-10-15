package main

import "fmt"

type agility struct {
	faction faction
}

func (a *agility) goGame(){
	fmt.Println("Agility Hero")
	a.faction.startGame()
}

func (a *agility) setHeroFaction (f faction){
	a.faction = f
}
func someDecorator(f func()) func(){
	return func(){
		fmt.Println("This is the Start of the Game")
		f()
		fmt.Println("Yes this is interesting")
	}
}
