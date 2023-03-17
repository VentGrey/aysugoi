package main

import (
	"fmt"
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	url, err := url.Parse("https://github.com/VentGrey/aysugoi")

	if err != nil {
		log.Fatal(err)
	}

	/* Application starts here */
	app := app.New()
	window := app.NewWindow("Aysugoi")
	window.Resize(fyne.NewSize(800, 600))

	// "Add anime" button
	addAnimeButton := widget.NewButton("Add anime", func() {
		fmt.Println("Add anime button pressed")
	})

	// "Add manga" button
	addMangaButton := widget.NewButton("Add manga", func() {
		fmt.Println("Add manga button pressed")
	})

	// Light / Dark mode switch button
	themeSwitch := widget.NewCheck("Dark mode", func(on bool) {
		if on {
			app.Settings().SetTheme(theme.DarkTheme())
		} else {
			app.Settings().SetTheme(theme.LightTheme())
		}
	})

	// Set themeSwitch to dark mode by default
	themeSwitch.SetChecked(true)

	// Top buttons container
	buttonsTop := container.NewHBox(
		layout.NewSpacer(),
		addAnimeButton,
		addMangaButton,
		themeSwitch,
	)

	// Create a list of portrait images (manga and anime)
	contentCovers := []fyne.CanvasObject{
		// Things
	}

	// If contentCovers is empty, display a message
	if len(contentCovers) == 0 {
		contentCovers = append(contentCovers, widget.NewLabel("You haven't added any content yet! Try adding some by clicking the buttons above."))
	}

	// Create a horizontal container for lower buttons send it to the bottom of the window
	buttonsBottom := container.NewHBox(
		layout.NewSpacer(),
		NewSettingsButton(window),
		widget.NewHyperlink("GitHub", url),
		NewAboutButton(window),
	)

	// Create a vertical container for the whole window
	mainContent := container.NewVBox(
		buttonsTop,
		container.NewVScroll(container.NewVBox(contentCovers...)),
		buttonsBottom,
	)

	window.SetContent(mainContent)
	window.SetMaster()
	window.Show()
	app.Run()
}
