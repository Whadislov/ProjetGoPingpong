package myapp

import (
	"image/color"
	"log"
	"time"

	//mdb "github.com/Whadislov/TTCompanion/internal/my_db"
	//mt "github.com/Whadislov/TTCompanion/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// StarterPage creates the introduction page to the UI and the starter page
func StarterPage() fyne.App {
	a := app.NewWithID("com.onrender.TTCompanion")

	// Set the icon
	icon, err := fyne.LoadResourceFromPath("Icon.png")
	if err != nil {
		log.Printf("Failed to load icon: %v", err)
	}
	a.SetIcon(icon)

	mainWindow := a.NewWindow("TT Companion")
	mainWindow.Resize(fyne.NewSize(600, 400))

	// Center the window on the monitor
	mainWindow.CenterOnScreen()

	// Welcome page
	themeColor := a.Settings().Theme().Color("foreground", a.Settings().ThemeVariant())
	welcomeText := canvas.NewText("Welcome to TT Companion", themeColor)
	welcomeText.Alignment = fyne.TextAlignCenter
	welcomeText.TextSize = 32
	welcomePage := container.NewCenter(welcomeText)

	// Main page
	//mainPage := MainPage(db, mainWindow, a)

	// Fade
	go func() {
		time.Sleep(1 * time.Second)
		if appStartOption == "local" {
			fadeText(welcomeText, themeColor)
			// go to main page with delay so that the menu is not directly shown
			log.Println("Transitioning to identification page")
			mainWindow.SetContent(AuthentificationPage(mainWindow, a))

		} else if appStartOption == "browser" {
			// No fade because it blinks on the browser and the problem is not yet solved
			log.Println("Transitioning to the authentification page web")
			mainWindow.SetContent(AuthentificationPageWeb(mainWindow, a))
		}

	}()
	log.Println("Displaying welcome page")
	mainWindow.SetContent(welcomePage)
	mainWindow.SetMainMenu(nil)
	mainWindow.ShowAndRun()
	return a
}

func fadeText(text *canvas.Text, textColor color.Color) {
	r, g, b, alp := textColor.RGBA()
	var fadeStep uint8 = 5
	var threshold uint8 = 120

	// >> 8 because color.RGBA can only use values of 8 bits (textColor is 16 bits)
	updateUI := func(alpha uint8) {
		text.Color = color.RGBA{
			R: uint8(r >> 8),
			G: uint8(g >> 8),
			B: uint8(b >> 8),
			A: alpha,
		}
		text.Refresh()
	}

	for alpha := uint8(alp >> 8); alpha >= threshold; alpha -= fadeStep {
		updateUI(alpha)
		text.Refresh()
		time.Sleep(20 * time.Millisecond) // Pause to simulate fade
	}

}
