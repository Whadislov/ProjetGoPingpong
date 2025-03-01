package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type myDarkTheme struct {
	IsActivated bool
}
type myLightTheme struct {
	IsActivated bool
}

func (t *myDarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, theme.VariantDark)
}

func (t *myLightTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, theme.VariantLight)
}

func (t *myDarkTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (t *myDarkTheme) Font(name fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(name)
}

func (t *myDarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *myLightTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (t *myLightTheme) Font(name fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(name)
}

func (t *myLightTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}
