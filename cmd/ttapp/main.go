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

	// Charger ou créer une nouvelle base de données
	db, err := md.LoadDatabase(filename)
	if err != nil {
		fmt.Println("Error while loading database :", err)
		return
	}

	c1, _ := mf.NewClub("TSG Heilbronn", db)

	m2, _ := mf.NewTeam("Mannschaft 2", c1, db)
	mf.NewTeam("Mannschaft 3", c1, db)
	m5, _ := mf.NewTeam("Mannschaft 5", c1, db)
	mf.NewTeam("Mannschaft 7", c1, db)
	mf.NewTeam("Mannschaft 8", c1, db)

	lasse, _ := mf.NewPlayer("Lasse", c1, db)
	lasse.SetPlayerAge(20)
	lasse.SetPlayerRanking(1300)
	lasse.SetPlayerMaterial("Victas V22 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	mf.AddPlayerToTeam(lasse, m5, c1)

	julien, _ := mf.NewPlayer("Julien", c1, db)
	julien.SetPlayerAge(27)
	julien.SetPlayerRanking(1585)
	julien.SetPlayerMaterial("Victas V20 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
	mf.AddPlayerToTeam(julien, m2, c1)

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

	// Sauvegarder les modifications
	err = md.SaveDatabase(filename, db)
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde de la base de données :", err)
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
