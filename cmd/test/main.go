package main

import (
	"fmt"
	md "github.com/Whadislov/ProjetGoPingPong/internal/my_database"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
)

func main() {

	db, _ := md.LoadDatabase("database.json")

	c1, _ := mf.NewClub("TSG Heilbronn", db)

	m2, _ := mf.NewTeam("Mannschaft 2", c1, db)
	m3, _ := mf.NewTeam("Mannschaft 3", c1, db)
	m5, _ := mf.NewTeam("Mannschaft 5", c1, db)
	m7, _ := mf.NewTeam("Mannschaft 7", c1, db)
	m8, _ := mf.NewTeam("Mannschaft 8", c1, db)

	fmt.Println("")
	fmt.Println("####################### club show")
	c1.Show()

	lasse, _ := mf.NewPlayer("Lasse", c1, db)
	lasse.SetPlayerAge(20)
	lasse.SetPlayerRanking(1300)
	lasse.SetPlayerMaterial("Victas V22 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	lasse.Show()
	mf.AddPlayerToTeam(lasse, m5, c1)

	julien, _ := mf.NewPlayer("Julien", c1, db)
	julien.SetPlayerAge(27)
	julien.SetPlayerRanking(1585)
	julien.SetPlayerMaterial("Victas V20 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	mf.AddPlayerToTeam(julien, m2, c1)
	julien.Show()

	robin, _ := mf.NewPlayer("Robin", c1, db)
	leon, _ := mf.NewPlayer("Leon", c1, db)
	patrick, _ := mf.NewPlayer("Patrick", c1, db)
	jonathan, _ := mf.NewPlayer("Jonathan", c1, db)
	sumi, _ := mf.NewPlayer("Sumi", c1, db)
	martin, _ := mf.NewPlayer("Martin", c1, db)

	mf.AddPlayerToTeam(robin, m2, c1)
	mf.AddPlayerToTeam(leon, m2, c1)
	mf.AddPlayerToTeam(patrick, m2, c1)
	mf.AddPlayerToTeam(jonathan, m2, c1)
	mf.AddPlayerToTeam(sumi, m2, c1)
	mf.AddPlayerToTeam(martin, m2, c1)
	fmt.Println("")
	fmt.Println("####################### team.Show()")
	m2.Show()

	mf.AddPlayerToTeam(julien, m3, c1)
	fmt.Println("")
	fmt.Println("####################### Julien 2 teams ")
	julien.Show()
	mf.AddPlayerToTeam(lasse, m5, c1)

	mf.AddPlayerToTeam(julien, m7, c1)

	fmt.Println("")
	fmt.Println("####################### club show")
	c1.Show()

	fmt.Println("")
	fmt.Println("####################### Team 7 delete show ")
	mf.DeleteTeam(m7, c1)
	m7.Show()
	fmt.Println("")
	fmt.Println("####################### Julien delete show ")
	mf.DeletePlayer(julien, c1)
	julien.Show()
	fmt.Println("")
	fmt.Println("####################### Team 2 show ")
	m2.Show()
	fmt.Println("")
	fmt.Println("#######################  Team 5 show ")
	mf.AddPlayerToTeam(robin, m5, c1)
	m5.Show()
	mf.RemovePlayerFromTeam(robin, m5, c1)
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
