package myapp

import (
	"image/color"
	"strconv"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PlayerInfos(player *mt.Player) *fyne.Container {
	nameStr := "Name: " + player.Name
	ageStr := "Age: " + strconv.Itoa(player.Age)
	materialStr := "Material: " + strHelper(player.Material)
	rankingStr := "Ranking: " + strconv.Itoa(player.Ranking)
	teams := []string{}
	clubs := []string{}
	for _, team := range player.TeamIDs {
		teams = append(teams, team)
	}
	for _, club := range player.ClubIDs {
		clubs = append(clubs, club)
	}
	teamsStr := "Teams: " + strHelper(teams)
	clubsStr := "Clubs: " + strHelper(clubs)
	item := container.NewVBox(
		widget.NewLabel(nameStr),
		widget.NewLabel(ageStr),
		widget.NewLabel(materialStr),
		widget.NewLabel(rankingStr),
		widget.NewLabel(teamsStr),
		widget.NewLabel(clubsStr),
	)
	return item
}

func PlayerMenu(w fyne.Window, players map[int]*mt.Player) {

	label := canvas.NewText("Players:", color.Black)
	ac := widget.NewAccordion()

	for _, player := range players {
		item := widget.NewAccordionItem(
			player.Name,
			PlayerInfos(player),
		)

		ac.Append(item)
	}

	w.SetContent(container.NewVBox(
		label,
		ac),
	)
}
