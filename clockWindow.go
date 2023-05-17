package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"time"
)

func updateTime(clock *widget.Label) {
	formattedTime := time.Now().Format("Current time 03:04:05")
	clock.SetText(formattedTime)
}

func createClockWindow(application fyne.App) {
	clock := widget.NewLabel("")
	updateTime(clock)
	clockWindow := application.NewWindow("This is a clock")
	clockWindow.SetContent(clock)
	clockWindow.Resize(fyne.Size{
		Width:  250,
		Height: 20,
	})

	// now update the time every second
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	clockWindow.Show()
}
