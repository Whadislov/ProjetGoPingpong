package myapp

import (
	"fmt"
	"log"

	mfr "github.com/Whadislov/TTCompanion/internal/my_frontend"
	mf "github.com/Whadislov/TTCompanion/internal/my_functions"

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

	emailLabel := widget.NewLabel("✉️ E-mail")
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
			log.Println("e-mail is not valid:", err)
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
				log.Println("Username or e-mail already exist")
				dialog.ShowError(fmt.Errorf("failed to sign up: %v", err), w)
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

	forgotPasswordButton := widget.NewButtonWithIcon("Forgot your password?", nil, func() {
		w.SetContent(reinitPasswordPageWeb(w, a))
	})
	forgotPasswordButton.Importance = widget.LowImportance

	cancelButton := widget.NewButton("Cancel", func() {
		w.SetContent(AuthentificationPageWeb(w, a))
	})

	loginPageWeb := container.NewVBox(
		pageTitle,
		usernameLabel,
		usernameEntry,
		container.NewBorder(nil, nil, passwordLabel, forgotPasswordButton),
		passwordEntry,
		validationButton,
		cancelButton,
	)

	return loginPageWeb
}

// reinitPasswordPage offers the user the ability to be sent an Email for a password reset
func reinitPasswordPageWeb(w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Forgot your password ?", 32)
	pageText := sexText("Enter your e-mail address and click on the button below to reset your password", 12)
	emailString := "Enter your e-mail address"

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder(emailString)

	validationButton := widget.NewButton("Reset your password", func() {
		b, err := mf.IsValidEmail(emailEntry.Text)
		if !b {
			dialog.ShowError(err, w)
			w.SetContent(reinitPasswordPageWeb(w, a))
		} else {

			// E-mail found
			var resetEmailDialogSuccess *dialog.CustomDialog
			cautionMessage := "⚠️ This feature is currently not working ⚠️\n\n"
			resetEmailMessage := widget.NewLabel(fmt.Sprintf(cautionMessage+"A message has been sent to %v.\nIf you can't find it in your inbox, check your junk mail.\nOtherwise, check that you have correctly entered the e-mail address you used to log in.", emailEntry.Text))
			returnToLoginPageWeb := widget.NewButton("Return to the login screen", func() {
				resetEmailDialogSuccess.Hide()
				w.SetContent(loginPageWeb(w, a))
			})
			resetEmailContent := container.NewVBox(
				resetEmailMessage,
				returnToLoginPageWeb,
			)

			resetEmailDialogSuccess = dialog.NewCustomWithoutButtons("Success", resetEmailContent, w)
			resetEmailDialogSuccess.Show()
		}

	})

	cancelButton := widget.NewButton("Cancel", func() {
		w.SetContent(loginPageWeb(w, a))
	})

	reinitPasswordPageWeb := container.NewVBox(
		pageTitle,
		pageText,
		emailEntry,
		validationButton,
		cancelButton,
	)

	return reinitPasswordPageWeb
}
