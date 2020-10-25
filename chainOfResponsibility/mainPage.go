package main

import "fmt"

type mainPage struct {
	next stage
}

func (m *mainPage) execute(p *player){
	if p.clickedToPlay{
		fmt.Println("Player clicked to start game button...")
		fmt.Println("Searching for a game")
		m.next.execute(p)
		return
	}
	fmt.Println("Game has been found")
	p.clickedToPlay = true
	m.next.execute(p)
}
func (m *mainPage) setNext(next stage){
	m.next=next
}