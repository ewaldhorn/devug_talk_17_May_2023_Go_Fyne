package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"time"
)

func createAnotherWindow(application fyne.App) {
	anotherWindow := application.NewWindow("This is another window")
	anotherWindow.Resize(fyne.Size{
		Width:  640,
		Height: 480,
	})

	drawingCanvas := anotherWindow.Canvas()
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	rect := canvas.NewRectangle(blue)
	drawingCanvas.SetContent(rect)
	isBlue := true

	go func() {
		for {
			time.Sleep(time.Second)
			if isBlue {
				rect.FillColor = green
				isBlue = false
			} else {
				rect.FillColor = blue
				isBlue = true
			}
			rect.Refresh()
		}
	}()

	anotherWindow.Show()
}
