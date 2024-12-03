package main

import (
	"fmt"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
)

func main() {

	c1, _ := mf.NewClub("TSG Heilbronn")


	
	m2, _ := mf.NewTeam("Mannschaft 2", &c1)
	m3, _ := mf.NewTeam("Mannschaft 3", &c1)
	m5, _ := mf.NewTeam("Mannschaft 5", &c1)
	m7, _:= mf.NewTeam("Mannschaft 7", &c1)
	m8, _ := mf.NewTeam("Mannschaft 8", &c1)

	fmt.Println("")
	fmt.Println("####################### club show")
	c1.Show()
	

	lasse, _  := mf.NewPlayer("Lasse", &c1)
	lasse.SetPlayerAge(20)
	lasse.SetPlayerRanking(1300)
	lasse.SetPlayerMaterial("Victas V22 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	lasse.Show()
	mf.AddPlayerToTeam(lasse, mf.GetName(m5), &c1)
	
	julien, _  := mf.NewPlayer("Julien", &c1)
	julien.SetPlayerAge(27)
	julien.SetPlayerRanking(1585)
	julien.SetPlayerMaterial("Victas V20 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	mf.AddPlayerToTeam(julien, mf.GetName(m2), &c1)
	julien.Show()

	robin, _  := mf.NewPlayer("Robin", &c1)
	leon, _  := mf.NewPlayer("Leon", &c1)
	patrick, _  := mf.NewPlayer("Patrick", &c1)
	jonathan, _  := mf.NewPlayer("Jonathan", &c1)
	sumi, _  := mf.NewPlayer("Sumi", &c1)
	martin, _  := mf.NewPlayer("Martin", &c1)

	mf.AddPlayerToTeam(robin, mf.GetName(m2), &c1)
	mf.AddPlayerToTeam(leon, mf.GetName(m2), &c1)
	mf.AddPlayerToTeam(patrick, mf.GetName(m2), &c1)
	mf.AddPlayerToTeam(jonathan, mf.GetName(m2), &c1)
	mf.AddPlayerToTeam(sumi, mf.GetName(m2), &c1)
	mf.AddPlayerToTeam(martin, mf.GetName(m2), &c1)
	fmt.Println("")
	fmt.Println("####################### team.Show()")
	m2.Show()

	mf.AddPlayerToTeam(julien, mf.GetName(m3), &c1)
	fmt.Println("")
	fmt.Println("####################### Julien 2 teams ")
	julien.Show()
	mf.AddPlayerToTeam(lasse, mf.GetName(m5), &c1)

	mf.AddPlayerToTeam(julien, mf.GetName(m7), &c1)

	fmt.Println("")
	fmt.Println("####################### club show")
	c1.Show()
	

	fmt.Println("")
	fmt.Println("####################### Team 7 delete show ")
	mf.DeleteTeam(mf.GetName(m7), &c1)
	m7.Show()
	fmt.Println("")
	fmt.Println("####################### Julien delete show ")
	mf.DeletePlayer(julien, &c1)
	julien.Show()
	fmt.Println("")
	fmt.Println("####################### Team 2 show ")
	m2.Show()
	fmt.Println("")
	fmt.Println("#######################  Team 5 show ")
	mf.AddPlayerToTeam(robin, mf.GetName(m5), &c1)
	m5.Show()
	mf.RemovePlayerFromTeam(robin, mf.GetName(m5), &c1)
	m5.Show()
	//fmt.Println(m5.PlayerList[1])
	fmt.Println("")
	fmt.Println("#######################  Team 8 show ")
	m8.Show()
	fmt.Println(m8.PlayerList)
	

	

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