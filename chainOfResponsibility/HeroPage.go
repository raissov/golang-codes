package main

import "fmt"

type HeroPage struct {
	next stage
}

func (m *HeroPage) execute(p *player){
	if p.chosenHero{
		fmt.Println("Choosing for a hero...")
		m.next.execute(p)
		return
	}
	fmt.Println("Player chose a hero")
	p.chosenHero = true
	m.next.execute(p)
}
func (m *HeroPage) setNext(next stage){
	m.next=next
}