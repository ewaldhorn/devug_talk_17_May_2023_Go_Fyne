package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/png"
	"os"
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

func getGoLogoImage() image.Image {
	imgFile, err := os.Open("./resources/go_logo_png.png")

	defer func(imgFile *os.File) {
		err := imgFile.Close()
		if err != nil {
			fmt.Println("Closing the file went quite meh!", err)
		}
	}(imgFile)

	if err != nil {
		fmt.Println("Oh snap! Where's the image dude?", err)
	}

	imgData, err := png.Decode(imgFile)

	if err != nil {
		fmt.Println("Ah no. The PNG is weird!", err)
	}

	return imgData.(image.Image)
}
func createGoLogoCanvasObject() fyne.CanvasObject {
	logoImage := canvas.NewImageFromImage(getGoLogoImage())
	logoImage.FillMode = canvas.ImageFillOriginal
	return logoImage
}
func createMainWindow(windowTitle string, application fyne.App) {
	clockButton := makeClockToggleButton(application)

	mainWindow := application.NewWindow(windowTitle)
	mainWindow.SetMaster() // if we close this chap, it's over folks!

	mainWindow.SetContent(container.NewVBox(
		widget.NewLabel("This is our super boring label. Yeah, it sucks."),
		clockButton,
		createGoLogoCanvasObject(),
	))

	// Let's set the size and position of the window
	mainWindow.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})
	mainWindow.CenterOnScreen()

	mainWindow.Show()
}
