package myapp

import (
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
)

// strHelper is a helper fonction that takes from example ["ok1", "ok2" , "ok3"] and returns "ok1, ok2, ok3"
func strHelper(list []string) string {
	str := ""
	for _, word := range list {
		str += word + ", "
	}
	// Remove extra ", "
	if len(str) > 2 {
		str = str[:len(str)-2]
	}
	return str
}

func SortMap[T ~map[int]V, V mt.Entity](m T) []struct {
	Key   int
	Value V
} {
	// This function takes as parameter a variable of type : mt.Player or mt.Team or mt.Club
	// It returns a struct that contains the keys and the name of player/team/club.
	// The slices are alphabeticaly sorted
	// Tip: it does not return a map, because maps can't be sorted

	// Get keys from the map
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// Sort keys alphabeticaly per value using m.Lastname. Keys is the sorted slice
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]].GetName() < m[keys[j]].GetName()
	})

	// Build a new slice with the pair key/value
	sorted := make([]struct {
		Key   int
		Value V
	}, len(keys))
	// The new slice is sorted thanks to the slice keys
	for i, k := range keys {
		sorted[i] = struct {
			Key   int
			Value V
		}{
			Key:   k,
			Value: m[k],
		}
	}

	return sorted
}

// Create a confirmation dialog
func ShowConfirmationDialog(w fyne.Window, message string, onConfirm func()) {
	d := dialog.NewConfirm("Confirm deletion", message, func(confirm bool) {
		if confirm {
			onConfirm()
		}
	}, w)
	d.Show()
}

// Reinit the text of a widget entry
func ReinitWidgetEntryText(entry *widget.Entry, entryHolder string) {
	entry.SetText("")
	entry.SetPlaceHolder(entryHolder)
}

// Verify if the string is letters only
func IsLettersOnly(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'z' {
			return false
		}
	}
	return true
}

// Verify if the string is numbers only
func IsNumbersOnly(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// setTitle sets the string as a title for the page. The string is centered, respects dark/light mode and has its size
func setTitle(s string, size float32) *canvas.Text {
	a := fyne.CurrentApp()
	themeColor := a.Settings().Theme().Color("foreground", a.Settings().ThemeVariant())
	title := canvas.NewText(s, themeColor)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = size
	return title
}

// setText sets the string as a text for the page. The string respects dark/light mode and has its size
func sexText(s string, size float32) *canvas.Text {
	a := fyne.CurrentApp()
	themeColor := a.Settings().Theme().Color("foreground", a.Settings().ThemeVariant())
	title := canvas.NewText(s, themeColor)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = size
	return title
}
