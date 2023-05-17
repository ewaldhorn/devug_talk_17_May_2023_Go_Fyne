package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func createMainWindow(windowTitle string, application fyne.App) {
	mainWindow := application.NewWindow(windowTitle)
	mainWindow.SetContent(widget.NewLabel("This is our super boring label. Yeah, it sucks."))
	mainWindow.SetMaster() // if we close this chap, it's over folks!
	mainWindow.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})
	mainWindow.CenterOnScreen()
	mainWindow.Show()
}
