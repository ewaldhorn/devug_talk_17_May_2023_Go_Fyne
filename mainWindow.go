package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeClockToggleButton(application fyne.App) *widget.Button {
	var clockWindow fyne.Window
	clockWindowHidden := true

	button := widget.NewButton("Clockit", func() {
		// And for giggles, something that updates itself
		if clockWindow == nil {
			clockWindow = createClockWindow(application)
			clockWindowHidden = false

			// we have a new window, let's intercept the close action so that we can handle that
			clockWindow.SetOnClosed(func() {
				// if we close this, reset the state correctly
				clockWindow = nil
				clockWindowHidden = true
			})
		} else {
			if clockWindowHidden {
				clockWindow.Show()
			} else {
				clockWindow.Hide()
			}
			// toggle the clock window state
			clockWindowHidden = !clockWindowHidden
		}
	})

	return button
}

func createMainWindow(windowTitle string, application fyne.App) {
	clockButton := makeClockToggleButton(application)

	mainWindow := application.NewWindow(windowTitle)
	mainWindow.SetMaster() // if we close this chap, it's over folks!

	mainWindow.SetContent(container.NewVBox(
		widget.NewLabel("This is our super boring label. Yeah, it sucks."),
		clockButton,
	))

	// Let's set the size and position of the window
	mainWindow.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})
	mainWindow.CenterOnScreen()

	mainWindow.Show()
}
