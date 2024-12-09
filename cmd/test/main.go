package main

import (
	"fmt"
	md "github.com/Whadislov/ProjetGoPingPong/internal/my_database"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
)

func main() {

	db, _ := md.LoadDb("database.json")

	c1, _ := mf.NewClub("TSG Heilbronn", db)

	m2, _ := mf.NewTeam("Mannschaft 2", db)
	m3, _ := mf.NewTeam("Mannschaft 3", db)
	m5, _ := mf.NewTeam("Mannschaft 5", db)
	m7, _ := mf.NewTeam("Mannschaft 7", db)
	m8, _ := mf.NewTeam("Mannschaft 8", db)

	mf.AddTeamToClub(m2, c1)
	mf.AddTeamToClub(m3, c1)
	mf.AddTeamToClub(m5, c1)
	mf.AddTeamToClub(m7, c1)
	mf.AddTeamToClub(m8, c1)

	fmt.Println("")
	fmt.Println("####################### club show")
	c1.Show()

	lasse, _ := mf.NewPlayer("Lasse", db)
	lasse.SetPlayerAge(20)
	lasse.SetPlayerRanking(1300)
	lasse.SetPlayerMaterial("Victas V22 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	lasse.Show()
	mf.AddPlayerToTeam(lasse, m5)

	julien, _ := mf.NewPlayer("Julien", db)
	julien.SetPlayerAge(27)
	julien.SetPlayerRanking(1585)
	julien.SetPlayerMaterial("Victas V20 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	mf.AddPlayerToTeam(julien, m2)
	julien.Show()

	robin, _ := mf.NewPlayer("Robin", db)
	leon, _ := mf.NewPlayer("Leon", db)
	patrick, _ := mf.NewPlayer("Patrick", db)
	jonathan, _ := mf.NewPlayer("Jonathan", db)
	sumi, _ := mf.NewPlayer("Sumi", db)
	martin, _ := mf.NewPlayer("Martin", db)

	mf.AddPlayerToClub(julien, c1)
	mf.AddPlayerToClub(lasse, c1)
	mf.AddPlayerToClub(robin, c1)
	mf.AddPlayerToClub(leon, c1)
	mf.AddPlayerToClub(patrick, c1)
	mf.AddPlayerToClub(jonathan, c1)
	mf.AddPlayerToClub(sumi, c1)
	mf.AddPlayerToClub(martin, c1)

	mf.AddPlayerToTeam(robin, m2)
	mf.AddPlayerToTeam(leon, m2)
	mf.AddPlayerToTeam(patrick, m2)
	mf.AddPlayerToTeam(jonathan, m2)
	mf.AddPlayerToTeam(sumi, m2)
	mf.AddPlayerToTeam(martin, m2)

	fmt.Println("")
	fmt.Println("####################### team.Show()")
	m2.Show()

	mf.AddPlayerToTeam(julien, m3)
	fmt.Println("")
	fmt.Println("####################### Julien 2 teams ")
	julien.Show()
	mf.AddPlayerToTeam(lasse, m5)

	mf.AddPlayerToTeam(julien, m7)

	fmt.Println("")
	fmt.Println("####################### club show")
	c1.Show()

	fmt.Println("")
	fmt.Println("####################### Team 7 delete show ")
	mf.DeleteTeam(m7, db)
	m7.Show()
	fmt.Println("")
	fmt.Println("####################### Julien delete show ")
	mf.DeletePlayer(julien, db)
	julien.Show()
	fmt.Println("")
	fmt.Println("####################### Team 2 show ")
	m2.Show()
	fmt.Println("")
	fmt.Println("#######################  Team 5 show ")
	mf.AddPlayerToTeam(robin, m5)
	m5.Show()
	mf.RemovePlayerFromTeam(robin, m5)
	m5.Show()
	//fmt.Println(m5.PlayerList[1])
	fmt.Println("")
	fmt.Println("#######################  Team 8 show ")
	m8.Show()

	/* Objectifs:
	 	Créer des équipes
	  	Créer des joueurs
	   	Attribuer des joueurs à des équipes
	   	Créer des matchs
	   	Créer un calendrier de matchs
	   	Créer un calendrier pour chaque joueur
	   	Ecrire les joueurs, team, calendrier dans un fichier externe et sauvegarder
	   	Créer un exécutable
	   	Créer des goroutines, mutex
		Créer un ID pour chaque joueur, deux joueurs peuvent avoir le meme nom





	*/

}
