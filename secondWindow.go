package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func createSecondWindow(application fyne.App) {
	secondWindow := application.NewWindow("This is distracting")
	secondWindow.SetContent(widget.NewProgressBarInfinite())
	secondWindow.Resize(fyne.Size{
		Width:  250,
		Height: 20,
	})
	secondWindow.Show()
}
