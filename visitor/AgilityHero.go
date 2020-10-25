package main

type AgilityHero struct {
	name string
	damage int
	moveSpeed int
}
func (a *AgilityHero) accept(b baf){
	b.bafForAgility(a)
}
func (a *AgilityHero) setName(name string){
	a.name=name
}
func (a *AgilityHero) getName() string{
	return a.name
}
func (a *AgilityHero) setDamage(damage int ){
	a.damage=damage
}
