package main

import (
	"fmt"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
)

func main() {
	
	m2 := mf.NewTeam("Mannschaft 2")
	m3 := mf.NewTeam("Mannschaft 3")
	m5 := mf.NewTeam("Mannschaft 5")
	m7 := mf.NewTeam("Mannschaft 7")

	lasse := mf.NewPlayer("Lasse")
	lasse.SetPlayerAge(20)
	lasse.SetPlayerRanking(1300)
	lasse.SetPlayerMaterial("Victas V22 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	lasse.Show()
	mf.AddPlayerToTeam(&lasse, &m5)
	
	julien := mf.NewPlayer("Julien")
	julien.SetPlayerAge(27)
	julien.SetPlayerRanking(1585)
	julien.SetPlayerMaterial("Victas V20 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	mf.AddPlayerToTeam(&julien, &m2)
	julien.Show()

	robin := mf.NewPlayer("Robin")
	leon := mf.NewPlayer("Leon")
	patrick := mf.NewPlayer("Patrick")
	jonathan := mf.NewPlayer("Jonathan")
	sumi := mf.NewPlayer("Sumi")
	martin := mf.NewPlayer("Martin")

	mf.AddPlayerToTeam(&robin, &m2)
	mf.AddPlayerToTeam(&leon, &m2)
	mf.AddPlayerToTeam(&patrick, &m2)
	mf.AddPlayerToTeam(&jonathan, &m2)
	mf.AddPlayerToTeam(&sumi, &m2)
	mf.AddPlayerToTeam(&martin, &m2)
	m2TeamComp := map[int]string{0:robin.Name, 1:leon.Name, 2:patrick.Name, 3:jonathan.Name, 4:sumi.Name, 5:martin.Name}
	fmt.Println("#######################")
	fmt.Println("team.Show() ")
	m2.SetTeamComposition(m2TeamComp)
	m2.Show()

	mf.AddPlayerToTeam(&julien, &m3)
	fmt.Println("#######################")
	fmt.Println("Julien 2 teams ")
	julien.Show()
	mf.AddPlayerToTeam(&lasse, &m5)

	mf.AddPlayerToTeam(&julien, &m7)

	fmt.Println("#######################")
	fmt.Println("Team 7 delete show ")
	mf.DeleteTeam(&m7)
	m7.Show()

	fmt.Println("#######################")
	fmt.Println("Julien delete show ")
	mf.DeletePlayer(&julien)
	julien.Show()
	fmt.Println("#######################")
	fmt.Println("Team 2 show ")
	m2.Show()
	fmt.Println(m2.PlayerList)
	fmt.Println(m5.PlayerList)
	

	

/* Objectifs:
Créer des équipes
Créer des joueurs
Attribuer des joueurs à des équipes
Créer des matchs
Créer un calendrier de matchs 
Créer un calendrier pour chaque joueur 
Ecrire les joueurs, team, calendrier dans un fichier externe et sauvegarder





*/

}