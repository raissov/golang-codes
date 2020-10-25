package main

func main()  {
	strength:=&StrengthHero{"Axe", 100, 280}
	doubleDamageRune:=&DoubleDamage{}
	strength.accept(doubleDamageRune)
	agility:=&AgilityHero{"Phantom Assassin", 120, 320}
	hasteRune:=&HasteRune{}
	agility.accept(hasteRune)
}
