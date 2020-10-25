package main

import "fmt"

//Observable object struct
type Hero struct {
	ID int
}

//Observable object interface
type HeroInterface interface {
	execute(m string)
}

// Typing some message with its ID
func (l *Hero) execute(m string) {
	fmt.Printf("%q message receiver for id %d \n", m, l.ID)
}

//Struct for observer
type Subject struct {
	heroes []HeroInterface
}

//Adding objects to slice
func (s *Subject) addHero(l HeroInterface) {
	s.heroes = append(s.heroes, l)
}

// showing them with iteration
func (s *Subject) notify(m string) {
	for _, l := range s.heroes {
		if l != nil {
			l.execute(m)
		}
	}
}

var iter int

func newHero() *Hero {
	l := Hero{iter}
	iter++
	return &l
}
func main() {
	iter = 0
	s := Subject{heroes: make([]HeroInterface, iter)}
	l := newHero()
	s.addHero(l)

	for i := 0; i < 5; i++ {
		l = newHero()
		s.addHero(l)
	}

	s.notify("Game has been started")
	s.notify("Ready for the game....")
	// or some other messages even voice message from game client
}
