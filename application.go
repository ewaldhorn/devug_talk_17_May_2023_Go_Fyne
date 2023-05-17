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

	// We need an application, to host all the windows...
	// That's why it's a desktop APPLICATION doh! 🤣
	application := app.New()
	mainWindow := application.NewWindow(windowTitle)

	mainWindow.SetContent(widget.NewLabel("This is our super boring label. Yeah, it sucks."))
	mainWindow.Show()

	application.Run()
}
