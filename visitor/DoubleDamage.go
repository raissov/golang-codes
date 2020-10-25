package main

import "fmt"

type DoubleDamage struct {
	moveSpeed int
	damage int
	name string
}

func (d *DoubleDamage) bafForAgility(a *AgilityHero) {
	fmt.Println("Damage has been increased to: ", a.damage*2, " ", a.getName())
}
func (d *DoubleDamage) bafForStrength(s *StrengthHero){
	fmt.Println("Damage has been increased to: ", s.damage*2, " ", s.getName())
}