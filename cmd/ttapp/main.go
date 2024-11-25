package main

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
)

func main() {
	
	m2 := mf.NewTeam("Mannschaft 2")
	m3 := mf.NewTeam("Mannschaft 3")
	m5 := mf.NewTeam("Mannschaft 5")

	lasse := mf.NewPlayer("Lasse")
	lasse.SetPlayerAge(20)
	lasse.SetPlayerRanking(1300)
	lasse.SetPlayerMaterial("Victas V22 double Extra", "Victas V 20 double Extra", "Koki Niwa Wood")
	lasse.SetPlayerTeam(m5)
	

	
	var julien mt.Player 
	julien.Name = "Julien"
	julien.Age = 27
	julien.Ranking = 1599
	julien.Material = []string{"Victas V 20 double Extra", "Victas V 20 double Extra", "Koki Niwa Wood"}
	julien.Teams = []mt.Team{m2}

	mf.AddPlayerToTeam(&julien, &m3)

/* Objectifs:
Créer des équipes
Créer des joueurs
Attribuer des joueurs à des équipes
Créer des matchs
Créer un calendrier de matchs 
Créer






*/

}