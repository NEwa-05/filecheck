package window

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var windowSize = fyne.Size{Width: 800, Height: 600}
var srcDirectory binding.String = binding.NewString()
var srcdupDirectory binding.String = binding.NewString()
var sameFileList binding.StringList = binding.NewStringList()

// CreateInitWindow create the window when starting app
func CreateInitWindow() {

	//create window
	a := app.NewWithID("APPID")
	initw := a.NewWindow("Filecheck")
	onefolderw := a.NewWindow("Selection du Dossier")
	twofolderw := a.NewWindow("Selection des Dossiers")
	// duplistw := a.NewWindow("Liste des doublons")
	initw.Resize(windowSize)

	// create button to select 1 folder mode
	oneFolderButton := widget.NewButton("1 Dossier", func() {
		initw.Hide()
		onefolderw.Show()
	})

	// create button to select 2 folders mode
	twoFoldersButton := widget.NewButton("2 Dossiers", func() {
		dialog.ShowFolderOpen(func(srcdupDir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, twofolderw)
				return
			}
			srcdupDirectory.Set(srcdupDir.Path())
			log.Println("selected folder in folderSelectionButton: ", srcdupDir.Path())
		}, twofolderw)
	})

	// create button to select 2 folders mode
	closeButton := widget.NewButton("Quitter", func() {
		initw.Close()
		log.Println("closing app")

	})

	// set window content
	initw.SetContent(
		container.NewVBox(
			container.NewCenter(
				oneFolderButton,
			),
			container.NewCenter(
				twoFoldersButton,
			),
			container.NewCenter(
				closeButton,
			),
		),
	)
	//show window when run
	initw.ShowAndRun()

}
