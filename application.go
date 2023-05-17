package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func runApplication(title ...string) {
	windowTitle := "Main Window"

	if len(title) > 0 {
		windowTitle = title[0]
	}

	application := app.New()
	mainWindow := application.NewWindow(windowTitle)

	mainWindow.SetContent(widget.NewLabel("This is our super boring label. Yeah, it sucks."))

	mainWindow.ShowAndRun()
}
