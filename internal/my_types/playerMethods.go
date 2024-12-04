package my_types

func (p *Player) SetPlayerAge(age int) {
	p.Age = age
}

func (p *Player) SetPlayerRanking(ranking int) {
	p.Ranking = ranking
}

func (p *Player) SetPlayerMaterial(forehand string, backhand string, blade string) {
	material := []string{forehand, backhand, blade}
	p.Material = material
}

func (p *Player) IsEmpty() bool {
	return p.Name == "" &&
		p.Age == 0 &&
		p.Ranking == 0 &&
		p.Material == nil &&
		p.TeamList == nil
}
