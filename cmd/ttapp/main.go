package main

import (
	"fmt"
	md "github.com/Whadislov/ProjetGoPingPong/internal/my_database"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	//"io"
	//"qt"
)

//https://golangr.com/gui

func main() {
	filename := "database.json"
	fmt.Println("ok")

	// Charger ou créer une nouvelle base de données
	db, err := md.LoadDb(filename)
	if err != nil {
		fmt.Println("Error while loading database:", err)
		return
	}

	c1, _ := mf.NewClub("TSG Heilbronn", db)

	m2, _ := mf.NewTeam("Mannschaft 2", c1, db)
	mf.NewTeam("Mannschaft 3", c1, db)
	m5, _ := mf.NewTeam("Mannschaft 5", c1, db)
	mf.NewTeam("Mannschaft 7", c1, db)
	mf.NewTeam("Mannschaft 8", c1, db)

	lasse, _ := mf.NewPlayer("Lasse", db)
	lasse.SetPlayerAge(20)
	lasse.SetPlayerRanking(1300)
	lasse.SetPlayerMaterial("Victas V22 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	mf.AddPlayerToTeam(lasse, m5)

	julien, _ := mf.NewPlayer("Julien", db)
	julien.SetPlayerAge(27)
	julien.SetPlayerRanking(1585)
	julien.SetPlayerMaterial("Victas V20 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	mf.AddPlayerToTeam(julien, m2)

	robin, _ := mf.NewPlayer("Robin", db)
	leon, _ := mf.NewPlayer("Leon", db)
	patrick, _ := mf.NewPlayer("Patrick", db)
	jonathan, _ := mf.NewPlayer("Jonathan", db)
	sumi, _ := mf.NewPlayer("Sumi", db)
	martin, _ := mf.NewPlayer("Martin", db)

	mf.AddPlayerToTeam(robin, m2)
	mf.AddPlayerToTeam(leon, m2)
	mf.AddPlayerToTeam(patrick, m2)
	mf.AddPlayerToTeam(jonathan, m2)
	mf.AddPlayerToTeam(sumi, m2)
	mf.AddPlayerToTeam(martin, m2)

	// Sauvegarder les modifications
	err = md.SaveDb(filename, db)
	if err != nil {
		fmt.Println("Error while saving database:", err)
		return
	}

	// Afficher les données
	db.Show()
}

/*
Enregistrer les nouveaux clubs/équipes/joueurs dans un fichier externe
Créer une liste de commandes pour l'utilisateur, il doit pouvoir :
créer clubs/équipes/joueurs
supprimer équipes/joueurs
voir clubs/équipes/joueurs
ajouter retirer joueurs aux équipes
*/
