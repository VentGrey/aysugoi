package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewAboutButton(window fyne.Window) *widget.Button {
	return widget.NewButtonWithIcon("", theme.HelpIcon(), func() {
        dialog.ShowInformation("About",
            "Aysugoi 0.1.0\nAysugoi is a simple anime and manga watchlist tracker made with Go + Fyne.\nAuthor: VentGrey\nLicense: GPL-3.0",
            window)
    })
}
