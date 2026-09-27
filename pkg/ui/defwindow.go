package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var WindowSize = fyne.Size{Width: 800, Height: 600}
var FilecheckApp fyne.App = app.NewWithID("filecheck")
var defWindow fyne.Window

func closeWindow(windowName fyne.Window) *widget.Button {
	closeButton := widget.NewButton("Quitter", func() {
		windowName.Close()
		log.Println("Closing Default Window")

	})

	return closeButton
}

func returnWindow(windowName fyne.Window) *widget.Button {
	returnButton := widget.NewButton("Retour", func() {
		windowName.Close()
		defWindow.Show()
		log.Println("Return to main Window")
	})

	return returnButton
}

// create the window when starting app
func CreateDefWindow() {

	defWindow = FilecheckApp.NewWindow("Filecheck")
	defWindow.SetMaster()
	defWindow.Resize(WindowSize)
	defWindow.SetTitle("Filecheck")

	// create button to select 1 folder mode
	oneFolderButton := widget.NewButton("1 Dossier", func() {
		oneFolderWindow()
		defWindow.Hide()
	})

	// create button to select 2 folders mode
	twoFoldersButton := widget.NewButton("2 Dossiers", func() {
		twoFolderWindow()
		defWindow.Hide()
	})

	// set window content
	defWindow.SetContent(
		container.NewVBox(
			container.NewCenter(
				oneFolderButton,
			),
			container.NewCenter(
				twoFoldersButton,
			),
			container.NewCenter(
				closeWindow(defWindow),
			),
		),
	)
	//show window when run
	defWindow.ShowAndRun()
}
