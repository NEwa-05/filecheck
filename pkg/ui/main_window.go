package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var WindowSize = fyne.Size{Width: 800, Height: 600}
var FilecheckApp fyne.App = app.NewWithID("filecheck")
var mainWindow fyne.Window
var SrcDirectory binding.String = binding.NewString()
var SrcdupDirectory binding.String = binding.NewString()
var SameFileList binding.StringList = binding.NewStringList()

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
		mainWindow.Show()
		log.Println("Return to main Window")
		SrcDirectory.Set("")
		SrcdupDirectory.Set("")
		SameFileList.Set(nil)
	})

	return returnButton
}

// create the window when starting app
func CreateMainWindow() {

	mainWindow = FilecheckApp.NewWindow("Filecheck")
	mainWindow.SetMaster()
	mainWindow.Resize(WindowSize)
	mainWindow.SetTitle("Filecheck")

	// create button to select 1 folder mode
	oneFolderButton := widget.NewButton("1 Dossier", func() {
		oneFolderWindow()
		mainWindow.Hide()
	})

	// create button to select 2 folders mode
	twoFoldersButton := widget.NewButton("2 Dossiers", func() {
		twoFolderWindow()
		mainWindow.Hide()
	})

	// set window content
	mainWindow.SetContent(
		container.NewBorder(
			container.NewBorder(
				container.NewHSplit(
					container.NewPadded(
						oneFolderButton,
					),
					container.NewPadded(
						twoFoldersButton,
					),
				),
				nil,
				nil,
				nil,
			),
			container.NewPadded(
				closeWindow(mainWindow),
			),
			nil,
			nil,
			nil,
		),
	)
	//show window when run
	mainWindow.ShowAndRun()
}
