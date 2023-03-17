package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewSettingsButton(window fyne.Window) *widget.Button {
	user_home, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	settingsFile := filepath.Join(user_home, ".aysugoi_config")

	usernameData, err := ioutil.ReadFile(settingsFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	username := strings.TrimSpace(string(usernameData))

	settingsButton := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		usernameEntry := widget.NewEntry()
		usernameEntry.SetText(username)
		messageLabel := widget.NewLabel("Sync with friends coming soon!\n\nUsername: " + username + "\n\nFor your username change to take effect, you must restart the app.")

		dialog.NewForm("Settings", "Save", "Cancel", []*widget.FormItem{
			widget.NewFormItem("Username", usernameEntry),
			widget.NewFormItem("", messageLabel),
		}, func(save bool) {
			if save {
				err := ioutil.WriteFile(settingsFile, []byte(usernameEntry.Text), 0644)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}, window).Show()
	})

	return settingsButton
}
