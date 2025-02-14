package myapp

import (
	"fmt"

	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mfr "github.com/Whadislov/ProjetGoPingPong/internal/my_frontend"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// UserPage sets up the page for displaying user info.
func UserPage(user *mt.User, db *mt.Database, w fyne.Window, a fyne.App) {
	pageTitle := setTitle("Your informations", 32)
	if HasChanged {
		pageTitle = setTitle("Your informations* (in edition)", 32)
	} else {
		// Display the correct user info after modifications even before saving
		currentUsername = user.Name
		currentEmail = user.Email
		currentPassword = user.PasswordHash
	}

	usernameLabel1 := widget.NewLabel("👤 Username:")
	usernameLabel2 := widget.NewLabel(currentUsername)
	usernameContent := container.NewGridWithColumns(2, usernameLabel1, usernameLabel2)

	emailLabel1 := widget.NewLabel("✉️ Email:")
	emailLabel2 := widget.NewLabel(currentEmail)
	emailContent := container.NewGridWithColumns(2, emailLabel1, emailLabel2)

	passwordLabel := widget.NewLabel("🔒 Password:")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetText(currentPassword)

	createdAtLabel1 := widget.NewLabel("📅 User since(yyyy-mm-dd): ")
	// Example format: 2025-02-03T09:58:59Z
	// [:10] -> only show the date and not the hours
	readableCreatedAt := user.CreatedAt[:10]
	createdAtLabel2 := widget.NewLabel(readableCreatedAt)
	createdAtContent := container.NewGridWithColumns(2, createdAtLabel1, createdAtLabel2)

	changeUsernameButton := widget.NewButton("Edit username", func() {
		ChangeUsernamePage(user, db, w, a)
	})

	changeEmailButton := widget.NewButton("Edit email", func() {
		ChangeEmailPage(user, db, w, a)
	})

	changePasswordButton := widget.NewButton("Edit password", func() {
		ChangePasswordPage(user, db, w, a)
	})

	passwordContent := container.NewGridWithColumns(2, passwordLabel, passwordEntry)

	saveChangesButton := widget.NewButton("Save changes", func() {
		if !HasChanged {
			dialog.ShowInformation("Information", "There is nothing new to save", w)
			return
		} else {
			var err error
			if appStartOption == "local" {
				err = mdb.SaveDB(db)
			} else if appStartOption == "browser" {
				err = mfr.SaveDB(jsonWebToken, db)
			}

			if err != nil {
				dialog.ShowError(err, w)
			} else {
				dialog.ShowInformation("Success", "Changes successfully saved", w)
				HasChanged = false
			}
			UserPage(user, db, w, a)
		}
	})

	returnToMainPageButton := widget.NewButton("Return to main page", func() {
		if HasChanged {
			dialog.ShowConfirm("Unsaved Changes", "You have unsaved changes. Do you want to save them before returning to the main page?", func(confirm bool) {
				if confirm {
					if appStartOption == "local" {
						err := mdb.SaveDB(db)
						if err != nil {
							dialog.ShowError(err, w)
						}
					} else if appStartOption == "browser" {
						err := mfr.SaveDB(jsonWebToken, db)
						if err != nil {
							dialog.ShowError(err, w)
						}
					}
					HasChanged = false
				}
				w.SetContent(MainPage(db, w, a))
			}, w)
		} else {
			w.SetContent(MainPage(db, w, a))
		}
		w.SetContent(MainPage(db, w, a))
	})

	content := container.NewVBox(
		pageTitle,
		container.NewGridWithColumns(2, usernameContent, changeUsernameButton),
		container.NewGridWithColumns(2, emailContent, changeEmailButton),
		container.NewGridWithColumns(2, passwordContent, changePasswordButton),
		createdAtContent,
		saveChangesButton,
		returnToMainPageButton,
	)

	w.SetContent(content)
}

// ChangeUsernamePage sets up the page to change user's username.
func ChangeUsernamePage(user *mt.User, db *mt.Database, w fyne.Window, a fyne.App) {

	pageTitle := setTitle("Editing 👤 username", 32)

	usernameLabel := widget.NewLabel("Your current username: " + user.Name)
	editUsernameEntry := widget.NewEntry()
	editUsernameEntry.SetPlaceHolder("Enter your new username")

	confirmButton := widget.NewButton("Confirm", func() {
		err := mf.ChangeUsername(user.Name, editUsernameEntry.Text, db)
		if err != nil {
			dialog.ShowError(err, w)
			editUsernameEntry.SetPlaceHolder("Enter your new username")
		} else {
			dialog.ShowInformation("Success", "Username successfully changed", w)
			HasChanged = true
			currentUsername = editUsernameEntry.Text
			UserPage(user, db, w, a)
		}

	})

	cancelButton := widget.NewButton("Cancel", func() {
		UserPage(user, db, w, a)
	})

	content := container.NewVBox(
		pageTitle,
		usernameLabel,
		editUsernameEntry,
		confirmButton,
		cancelButton,
	)
	w.SetContent(content)
}

// ChangeEmailPage sets up the page to change user's email.
func ChangeEmailPage(user *mt.User, db *mt.Database, w fyne.Window, a fyne.App) {

	pageTitle := setTitle("Editing ✉️ email", 32)

	emailLabel := widget.NewLabel("Your current email: " + user.Email)
	editEmailEntry := widget.NewEntry()
	editEmailEntry.SetPlaceHolder("Enter your new email. Example: abc@def.com")

	confirmButton := widget.NewButton("Confirm", func() {
		err := mf.ChangeEmail(editEmailEntry.Text, user)
		if err != nil {
			dialog.ShowError(err, w)
			editEmailEntry.SetPlaceHolder("Enter your new email. Example: abc@def.com")
		} else {
			dialog.ShowInformation("Success", "Email successfully changed", w)
			HasChanged = true
			currentEmail = editEmailEntry.Text
			UserPage(user, db, w, a)
		}

	})

	cancelButton := widget.NewButton("Cancel", func() {
		UserPage(user, db, w, a)
	})

	content := container.NewVBox(
		pageTitle,
		emailLabel,
		editEmailEntry,
		confirmButton,
		cancelButton,
	)
	w.SetContent(content)
}

// ChangePasswordPage sets up the page to change user's email.
func ChangePasswordPage(user *mt.User, db *mt.Database, w fyne.Window, a fyne.App) {

	pageTitle := setTitle("Editing 🔒 password", 32)

	passwordLabel := widget.NewLabel("Enter your current password")

	currentPasswordEntry := widget.NewPasswordEntry()
	editPasswordEntry := widget.NewPasswordEntry()
	confirmEditPasswordEntry := widget.NewPasswordEntry()
	editPasswordLabel := widget.NewLabel("Enter your new password")
	confirmEditPasswordLabel := widget.NewLabel("Confirm your new password")

	confirmButton := widget.NewButton("Confirm", func() {
		if currentPasswordEntry.Text != user.PasswordHash {
			dialog.ShowError(fmt.Errorf("wrong password"), w)
			currentPasswordEntry.SetText("")
		} else if editPasswordEntry.Text != confirmEditPasswordEntry.Text {
			dialog.ShowError(fmt.Errorf("passwords do not match"), w)
			editPasswordEntry.SetText("")
			confirmEditPasswordEntry.SetText("")
		} else {
			err := mf.ChangePassword(editPasswordEntry.Text, user)
			if err != nil {
				dialog.ShowError(err, w)
				editPasswordEntry.SetText("")
				confirmEditPasswordEntry.SetText("")
			} else {
				dialog.ShowInformation("Success", "Password successfully changed", w)
				HasChanged = true
				currentPassword = editPasswordEntry.Text
				UserPage(user, db, w, a)
			}
		}
	})

	cancelButton := widget.NewButton("Cancel", func() {
		UserPage(user, db, w, a)
	})

	content := container.NewVBox(
		pageTitle,
		passwordLabel,
		currentPasswordEntry,
		editPasswordLabel,
		editPasswordEntry,
		confirmEditPasswordLabel,
		confirmEditPasswordEntry,
		confirmButton,
		cancelButton,
	)
	w.SetContent(content)

}
