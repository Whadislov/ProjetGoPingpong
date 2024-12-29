package myapp

import (
	"image/color"
	"time"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// StarterPage creates the introduction page to the UI and the starter page
func StarterPage(db *mt.Database) fyne.App {
	a := app.New()
	mainWindow := a.NewWindow("TTapp")
	mainWindow.Resize(fyne.NewSize(600, 400))

	// Center the window on the monitor
	mainWindow.CenterOnScreen()

	// Welcome page
	welcomeText := canvas.NewText("Welcome to TTapp 🏓", color.RGBA{R: 0, G: 0, B: 0, A: 255})
	welcomeText.Alignment = fyne.TextAlignCenter
	welcomeText.TextSize = 32
	welcomePage := container.NewCenter(welcomeText)

	// Main page

	mainPage := MainPage(db, mainWindow, a)

	// Fade
	go func() {
		time.Sleep(1 * time.Second)
		for alpha := 255; alpha >= 0; alpha -= 5 {
			welcomeText.Color = color.RGBA{R: 0, G: 0, B: 0, A: uint8(alpha)} // Opacity
			time.Sleep(20 * time.Millisecond)                                 // Pause to simulate fade
			welcomeText.Refresh()
		}

		// go to main page with delay so that the menu is not directly shown
		mainWindow.SetContent(mainPage)
		mainMenu := MainMenu(db, mainWindow, a)
		mainWindow.SetMainMenu(mainMenu)
	}()
	mainWindow.SetContent(welcomePage)
	mainWindow.SetMainMenu(nil)
	mainWindow.ShowAndRun()
	return a
}
