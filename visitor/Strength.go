package main

type StrengthHero struct {
	name string
	damage int
	moveSpeed int
}
func (s *StrengthHero) accept(b baf){
	b.bafForStrength(s)
}
func (s *StrengthHero) setName(name string){
	s.name=name
}
func (s *StrengthHero) getName() string{
	return s.name
}
func (s *StrengthHero) setDamage(damage int ){
	s.damage=damage
}
func (s *StrengthHero) setMoveSpeed(moveSpeed int){
	s.moveSpeed=moveSpeed
}
