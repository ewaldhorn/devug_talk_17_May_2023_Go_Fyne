package main

import (
	"fyne.io/fyne/v2/app"
)

func runApplication(title ...string) {
	// Go's default parameter handling is meh
	windowTitle := "Main Window"

	if len(title) > 0 {
		windowTitle = title[0]
	}

	// We need an application, to host all the windows...
	// That's why it's a desktop APPLICATION doh! 🤣
	application := app.New()

	// Let's create a main window, because you need one, right?
	createMainWindow(windowTitle, application)

	// What about another window? Gotta distract the users...
	createSecondWindow(application)

	// And for giggles, something that updates itself
	createClockWindow(application)

	// And now that everything is ready, launch the app!
	application.Run()
}
