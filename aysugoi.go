package main

import (
	"fmt"
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Anime struct {
    Image     string     `json:"image"`
    Title     string     `json:"title"`
    Parts     int        `json:"parts"`
    Episodes  [][]int    `json:"episodes"`
}

func main() {
	url, err := url.Parse("https://github.com/VentGrey/aysugoi")

	if err != nil {
		log.Fatal(err)
	}

	/* Application starts here */
	app := app.New()
	window := app.NewWindow("Aysugoi")
	window.Resize(fyne.NewSize(800, 600))

	addAnimeButton := widget.NewButton("Add anime", func() {
		// Open a new dialog with a form to add an anime to the $HOME/.config/aysugoi/anime.json file
		// The form should have the following fields:
		// - Image (file picker)
		// - Title (text input)
		// - Parts (number input)
		// - Episodes (Per part, a list of numbers)
		// The form should be validated before submitting

		imagePicker := widget.NewButton("Choose image", func() {
			dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err == nil && reader != nil {
					// Aquí puedes manejar el archivo que el usuario seleccionó.
					// En este ejemplo, simplemente mostramos el nombre del archivo.
					dialog.ShowInformation("Selected file", reader.URI().Name(), window)
					reader.Close()
				}
			}, window)
		})
		titleInput := widget.NewEntry()
		partsInput := widget.NewEntry()
		episodesInput := widget.NewEntry()


		dialog.NewForm("Add anime", "Submit", "Cancel", []*widget.FormItem{
			widget.NewFormItem("Image: ", imagePicker),
			widget.NewFormItem("Title: ", titleInput),
			widget.NewFormItem("Parts: ", partsInput),
			widget.NewFormItem("Episodes: ", episodesInput),
		}, func(save bool) {

		}, window).Show()
	})

	addMangaButton := widget.NewButton("Add manga", func() {
		fmt.Println("Add manga button pressed")
	})

	themeSwitch := widget.NewCheck("Dark mode", func(on bool) {
		if on {
			app.Settings().SetTheme(theme.DarkTheme())
		} else {
			app.Settings().SetTheme(theme.LightTheme())
		}
	})

	themeSwitch.SetChecked(true)

	buttonsTop := container.NewHBox(
		layout.NewSpacer(),
		addAnimeButton,
		addMangaButton,
		themeSwitch,
	)

	contentCovers := []fyne.CanvasObject{
		// Things
	}

	// If contentCovers is empty, display a message
	if len(contentCovers) == 0 {
		contentCovers = append(contentCovers, widget.NewLabel("You haven't added any content yet! Try adding some by clicking the buttons above."))
	}

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
