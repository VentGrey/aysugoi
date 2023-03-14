package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Aysugoi")

	window.SetContent(widget.NewLabel("Aysugoi!"))
	window.SetMaster()
	window.Show()


	app.Run()
}
