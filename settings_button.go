package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewSettingsButton(window fyne.Window) *widget.Button {
	settingsButton := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		usernameEntry := widget.NewEntry()
		messageLabel := widget.NewLabel("Sync with friends coming soon!")

		dialog.NewForm("Settings", "Save", "Cancel", []*widget.FormItem{
			widget.NewFormItem("Username", usernameEntry),
			widget.NewFormItem("", messageLabel),
		}, func(save bool) {
			if save {
				fmt.Println("Username:", usernameEntry.Text)
			}
		}, window).Show()
	})

	return settingsButton
}
