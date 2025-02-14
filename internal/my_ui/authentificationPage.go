package myapp

import (
	"fmt"
	"log"

	mdb "github.com/Whadislov/TTCompanion/internal/my_db"
	mf "github.com/Whadislov/TTCompanion/internal/my_functions"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"golang.org/x/crypto/bcrypt"
)

// AuthentificationPage returns a page that contains a log in button and a sign up button
func AuthentificationPage(w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("TT Companion", 32)

	authLabel := widget.NewLabel("Please authenticate")
	authLabel.Alignment = fyne.TextAlignCenter

	var db *mt.Database
	var err error
	db, err = mdb.LoadUsersOnly()

	if err != nil {
		panic(err)
	}
	logInButton := widget.NewButton("Log in", func() {
		w.SetContent(loginPage(db, w, a))
	})

	signUpButton := widget.NewButton("Sign up", func() {
		w.SetContent(signUpPage(db, w, a))
	})

	if len(db.Users) > 0 {
		authentificationPage := container.NewVBox(
			pageTitle,
			authLabel,
			logInButton,
			signUpButton,
		)
		return authentificationPage
	} else {

		quitButtonInSignUpPage := widget.NewButton("Quit", func() {
			a.Quit()
			log.Println("Application exited.")
		})
		content := container.NewVBox(
			signUpPage(db, w, a),
			quitButtonInSignUpPage,
		)
		// No user in the database, go directly to sign up page
		w.SetContent(content)
		return signUpPage(db, w, a)
	}
}

// signUpPage returns a page that contains a create username and a create password field. Adds a new user in the database
func signUpPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
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
		newUser, err := mf.NewUser(usernameEntry.Text, emailEntry.Text, passwordEntry.Text, confirmPasswordEntry.Text, db)
		if err != nil {
			switch err.Error() {
			case "e-mail cannot be empty":
				log.Println("e-mail is empty")
				dialog.ShowError(err, w)
				emailEntry.SetPlaceHolder("abc@def.com")
			case "e-mail is already used":
				log.Println("e-mail is already used")
				dialog.ShowError(err, w)
				emailEntry.SetPlaceHolder("abc@def.com")
			case "e-mail must be valid. Example: abc@def.com":
				log.Println("e-mail is not valid")
				dialog.ShowError(err, w)
				emailEntry.SetPlaceHolder("abc@def.com")
			case "username cannot be empty":
				log.Println("username is empty")
				dialog.ShowError(err, w)
				usernameEntry.SetPlaceHolder("")
			case "username must be valid (only letters and figures are allowed, spaces are not allowed)":
				log.Println("username is not valid")
				dialog.ShowError(err, w)
				usernameEntry.SetPlaceHolder("")
			case "username is already taken":
				log.Println("username is already taken")
				dialog.ShowError(err, w)
				usernameEntry.SetPlaceHolder("")
			case "password cannot be empty":
				log.Println("password cannot be empty")
				dialog.ShowError(err, w)
				passwordEntry.SetText("")
				confirmPasswordEntry.SetText("")
			case "password must be valid (spaces are not allowed)":
				log.Println("password is not valid")
				dialog.ShowError(err, w)
				passwordEntry.SetText("")
				confirmPasswordEntry.SetText("")
			case "passwords do not match":
				log.Println("passwords do not match")
				dialog.ShowError(err, w)
				passwordEntry.SetText("")
				confirmPasswordEntry.SetText("")
			default:
				log.Println("unknown error")
				dialog.ShowError(err, w)
				w.SetContent(signUpPage(db, w, a))
			}
		} else {
			var err error
			dialog.ShowInformation("Success", "Your user account has been created !", w)
			mdb.SetUserIDOfSession(newUser.ID)

			// Save the new user in the database
			mdb.SaveDB(db)
			log.Println("Sign up is successfull")

			// Now load the whole database
			db, err = mdb.LoadDB()

			if err != nil {
				dialog.ShowError(err, w)
			} else {
				userOfSession = db.Users[newUser.ID]
				w.SetContent(MainPage(db, w, a))
				w.SetMainMenu(MainMenu(db, w, a))
			}
		}
	})

	cancelButton := widget.NewButton("Cancel", func() {
		w.SetContent(AuthentificationPage(w, a))
	})

	signUpPage := container.NewVBox(
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

	return signUpPage
}

// loginPage returns a page that contains a enter username and a enter password field. Checks if the user is in the database. If yes, sets the variable user_id for the rest of the program
func loginPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Login", 32)

	usernameLabel := widget.NewLabel("👤 Username")
	usernameEntry := widget.NewEntry()

	passwordLabel := widget.NewLabel("🔒 Password")
	passwordEntry := widget.NewPasswordEntry()

	validationButton := widget.NewButton("Connect", func() {
		// Verify if username and password match
		log.Println("Verifying username and password")
		for _, user := range db.Users {
			if usernameEntry.Text == user.Name {
				err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(passwordEntry.Text))
				if err != nil {
					dialog.ShowError(fmt.Errorf("internal error: %v", err), w)
					return
				} else {
					log.Println("Login is successfull")
					mdb.SetUserIDOfSession(user.ID)
					// Now load the corresponding database of the user
					var err error

					mdb.SetUserIDOfSession(user.ID)
					db, err = mdb.LoadDB()

					if err != nil {
						dialog.ShowError(err, w)
					} else {
						userOfSession = db.Users[user.ID]
						w.SetContent(MainPage(db, w, a))
						w.SetMainMenu(MainMenu(db, w, a))
						return
					}
				}
			}
		}
		// Username and password mismatch
		log.Println("Login failed : Username and password mismatch")
		dialog.ShowError(fmt.Errorf("username and password mismatch"), w)
		// Reset entries
		ReinitWidgetEntryText(usernameEntry, "")
		passwordEntry.SetText("")

	})

	forgotPasswordButton := widget.NewButtonWithIcon("Forgot your password?", nil, func() {
		w.SetContent(reinitPasswordPage(db, w, a))
	})
	forgotPasswordButton.Importance = widget.LowImportance

	cancelButton := widget.NewButton("Cancel", func() {
		w.SetContent(AuthentificationPage(w, a))
	})

	loginPage := container.NewVBox(
		pageTitle,
		usernameLabel,
		usernameEntry,
		container.NewBorder(nil, nil, passwordLabel, forgotPasswordButton),
		passwordEntry,
		validationButton,
		cancelButton,
	)

	return loginPage
}

// reinitPasswordPage offers the user the ability to be sent an Email for a password reset
func reinitPasswordPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Forgot your password ?", 32)
	pageText := sexText("Enter your e-mail address and click on the button below to reset your password", 12)
	emailString := "Enter your e-mail address"

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder(emailString)

	validationButton := widget.NewButton("Reset your password", func() {
		b, err := mf.IsValidEmail(emailEntry.Text)
		if !b {
			dialog.ShowError(err, w)
			w.SetContent(reinitPasswordPage(db, w, a))
		} else {
			for _, user := range db.Users {
				if emailEntry.Text == user.Email {
					// E-mail found
					var resetEmailDialogSuccess *dialog.CustomDialog
					resetEmailMessage := widget.NewLabel(fmt.Sprintf("A message has been sent to %v.\nIf you can't find it in your inbox, check your junk mail.\nOtherwise, check that you have correctly entered the e-mail address you used to log in.", emailEntry.Text))
					returnToLoginPage := widget.NewButton("Return to the login screen", func() {
						resetEmailDialogSuccess.Hide()
						w.SetContent(loginPage(db, w, a))
					})
					resetEmailContent := container.NewVBox(
						resetEmailMessage,
						returnToLoginPage,
					)

					resetEmailDialogSuccess = dialog.NewCustomWithoutButtons("Success", resetEmailContent, w)
					resetEmailDialogSuccess.Show()
				}
			}
			// E-mail valid, but not found
			var resetEmailDialogFail *dialog.CustomDialog
			goToSignupPage := widget.NewButton("Create your account", func() {
				resetEmailDialogFail.Hide()
				w.SetContent(signUpPage(db, w, a))
			})
			returnToLoginPage := widget.NewButton("Try again", func() {
				resetEmailDialogFail.Hide()
				w.SetContent(reinitPasswordPage(db, w, a))
			})
			resetEmailMessageFail := widget.NewLabel(fmt.Sprintf("Your e-mail %v has not been found in our database.\nPlease check that you have correctly entered the e-mail address you used to log in.", emailEntry.Text))
			resetEmailContentFail := container.NewVBox(
				resetEmailMessageFail,
				returnToLoginPage,
				goToSignupPage,
			)

			resetEmailDialogFail = dialog.NewCustomWithoutButtons("Attention", resetEmailContentFail, w)
			resetEmailDialogFail.Show()
		}

	})

	cancelButton := widget.NewButton("Cancel", func() {
		w.SetContent(loginPage(db, w, a))
	})

	reinitPasswordPage := container.NewVBox(
		pageTitle,
		pageText,
		emailEntry,
		validationButton,
		cancelButton,
	)

	return reinitPasswordPage
}
