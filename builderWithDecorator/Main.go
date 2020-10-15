package main

func main(){
	direFaction:=&dire{}
	agilityHero:= &agility{}
	agilityHero.setHeroFaction(direFaction)
	// agilityHero.goGame - without decorator
	someDecorator(agilityHero.goGame)() //with decorator

	radiantFaction:=&radiant{}
	strengthHero:=&strength{}
	strengthHero.setHeroFaction(radiantFaction)

	strengthHero.goGame() // without decorator

}
