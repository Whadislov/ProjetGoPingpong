package myapp

import (
	"fmt"
	"log"

	mfr "github.com/Whadislov/ProjetGoPingPong/internal/my_frontend"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// AuthentificationPageWeb returns a page that contains a log in button and a sign up button
func AuthentificationPageWeb(w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle("TT Companion", 32)

	authLabel := widget.NewLabel("Please authenticate")
	authLabel.Alignment = fyne.TextAlignCenter

	logInButton := widget.NewButton("Log in", func() {
		w.SetContent(loginPageWeb(w, a))
	})

	signUpButton := widget.NewButton("Sign up", func() {
		w.SetContent(signUpPageWeb(w, a))
	})

	authentificationPageWeb := container.NewVBox(
		pageTitle,
		authLabel,
		logInButton,
		signUpButton,
	)

	w.SetContent(authentificationPageWeb)
	return authentificationPageWeb

}

// signUpPageWeb returns a page that contains a create username and a create password field. Adds a new user in the database
func signUpPageWeb(w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Create your account", 32)

	emailLabel := widget.NewLabel("✉️ Email")
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("abc@def.com")

	usernameLabel := widget.NewLabel("👤 Username")
	usernameEntry := widget.NewEntry()

	passwordLabel := widget.NewLabel("🔒 Password")
	passwordEntry := widget.NewPasswordEntry()

	confirmPasswordLabel := widget.NewLabel("Confirm Password")
	confirmPasswordEntry := widget.NewPasswordEntry()

	validationButton := widget.NewButton("Create", func() {
		log.Println("Creating new User")
		_, err := mf.IsValidEmail(emailEntry.Text)
		_, err2 := mf.IsValidUsername(usernameEntry.Text)
		_, err3 := mf.IsValidPassword(passwordEntry.Text)

		if err != nil {
			dialog.ShowError(err, w)
			log.Println("email is not valid:", err)
			w.SetContent(signUpPageWeb(w, a))
		} else if err2 != nil {
			dialog.ShowError(err, w)
			log.Println("username is not valid:", err)
			w.SetContent(signUpPageWeb(w, a))
		} else if err3 != nil {
			dialog.ShowError(err, w)
			log.Println("password is not valid:", err)
			w.SetContent(signUpPageWeb(w, a))
		} else if passwordEntry.Text != confirmPasswordEntry.Text {
			dialog.ShowError(fmt.Errorf("passwords do not match"), w)
			log.Println("passwords do not match:")
			w.SetContent(signUpPageWeb(w, a))
		} else {
			db, token, err := mfr.SignUp(usernameEntry.Text, passwordEntry.Text, emailEntry.Text)
			jsonWebToken = token
			// Last check if username or email already exist
			if err != nil {
				// Need to recheck this err
				log.Println("Username or email already exist")
				dialog.ShowError(err, w)
				w.SetContent(signUpPageWeb(w, a))
			} else {
				log.Println("Sign up is successfull")
				// Set the user of the session for the profile page on the menu. There is only one user on the Users map
				for _, user := range db.Users {
					if user.Name == usernameEntry.Text {
						userOfSession = user
					} else {
						dialog.ShowError(fmt.Errorf("failed to load your profile: "), w)
					}
				}
				w.SetContent(MainPage(db, w, a))
				w.SetMainMenu(MainMenu(db, w, a))
			}
		}
	})

	cancelButton := widget.NewButton("Cancel", func() {
		w.SetContent(AuthentificationPageWeb(w, a))
	})

	signUpPageWeb := container.NewVBox(
		pageTitle,
		emailLabel,
		emailEntry,
		usernameLabel,
		usernameEntry,
		passwordLabel,
		passwordEntry,
		confirmPasswordLabel,
		confirmPasswordEntry,
		validationButton,
		cancelButton,
	)

	return signUpPageWeb
}

// logInPageWeb returns a page that contains a enter username and a enter password field. Checks if the user is in the database. If yes, sets the variable user_id for the rest of the program
func loginPageWeb(w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Login", 32)

	usernameLabel := widget.NewLabel("👤 Username")
	usernameEntry := widget.NewEntry()

	passwordLabel := widget.NewLabel("🔒 Password")
	passwordEntry := widget.NewPasswordEntry()

	validationButton := widget.NewButton("Connect", func() {
		db, token, err := mfr.Login(usernameEntry.Text, passwordEntry.Text)
		jsonWebToken = token
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to login: %v", err), w)
			// Reset entries
			w.SetContent(loginPageWeb(w, a))
		} else {
			log.Println("Login is successfull")
			// Set the user of the session for the profile page on the menu. There is only one user on the Users map
			for _, user := range db.Users {
				if user.Name == usernameEntry.Text {
					userOfSession = user
				} else {
					dialog.ShowError(fmt.Errorf("failed to load your profile: "), w)
				}
			}
			w.SetContent(MainPage(db, w, a))
			w.SetMainMenu(MainMenu(db, w, a))
		}
	})

	cancelButton := widget.NewButton("Cancel", func() {
		w.SetContent(AuthentificationPageWeb(w, a))
	})

	loginPageWeb := container.NewVBox(
		pageTitle,
		usernameLabel,
		usernameEntry,
		passwordLabel,
		passwordEntry,
		validationButton,
		cancelButton,
	)

	return loginPageWeb
}
