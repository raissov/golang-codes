package main

import "fmt"

type GamePage struct {
	next stage
}

func (g *GamePage) execute(p *player){
	if p.goToLine{
		fmt.Println("Choosing for a line...")
	}
	fmt.Println("Player chose a line")
}
func (g *GamePage) setNext(next stage){
	g.next=next
}
