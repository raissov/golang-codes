package main

import "fmt"

type HasteRune struct {
	moveSpeed int
	damage int
	name string
}

func (h *HasteRune) bafForAgility(a *AgilityHero) {
	a.moveSpeed = 550
	fmt.Println("Movement Speed has been increased to: ", a.moveSpeed, " ", a.getName())
}
func (h *HasteRune) bafForStrength(s *StrengthHero){
	s.moveSpeed = 550
	fmt.Println("Movement Speed has been increased to: ", s.moveSpeed, " ", s.getName())
}
