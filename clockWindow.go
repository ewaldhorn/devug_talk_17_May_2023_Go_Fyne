package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"time"
)

func updateTime(clock *widget.Label) {
	formattedTime := time.Now().Format("The current time is 15:04:05.")
	clock.SetText(formattedTime)
}

func createClockWindow(application fyne.App) fyne.Window {
	clock := widget.NewLabel("")
	updateTime(clock)
	clockWindow := application.NewWindow("This is a clock")
	clockWindow.SetContent(clock)
	clockWindow.Resize(fyne.Size{
		Width:  250,
		Height: 20,
	})
	clockWindow.CenterOnScreen()

	// now update the time every second, until the end of time, or the end of the app, at least
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	clockWindow.Show()
	return clockWindow
}
