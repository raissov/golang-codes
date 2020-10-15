package main

import "fmt"

type strength struct {
	faction faction
}

func (s *strength) goGame(){
	fmt.Println("Strength Hero")
	s.faction.startGame()
}

func (s *strength) setHeroFaction (f faction){
	s.faction = f
}
