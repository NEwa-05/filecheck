package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var windowSize = fyne.Size{Width: 800, Height: 480}
var directory binding.String = binding.NewString()
var file binding.String = binding.NewString()

func oneFolderCheck(dir string) []string {
	filelist := createfilelist(dir)
	log.Printf("file list: %v", filelist)
	return filelist
}

func main() {
	//create window
	a := app.NewWithID("APPID")
	w := a.NewWindow("Filecheck")
	w.Resize(windowSize)

	// create button to select folder
	folderSelectionButton := widget.NewButton("Sélection", func() {
		dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			directory.Set(dir.Path())
			log.Println("selected folder in folderSelectionButton: ", dir.Path())
		}, w)
	})

	//create a text box with text
	showTextfolderSelect := widget.NewLabel("Dossier selectionné: ")

	//create a text box with the name of the folder selected
	showSelectedFolder := widget.NewLabelWithData(directory)

	//button that will start the check process
	checkFolderContent := widget.NewButton("Chercher les fichiers en doubles", func() {
		d := dialog.NewCustom("Recherche en cours", "Annuler", widget.NewProgressBarInfinite(), w)
		d.Show()
		folder, _ := directory.Get()
		log.Println("Selected folder in checkFolderContent: ", folder)
		log.Printf("filelist: %v")
		oneFolderCheck(folder)
	},
	)

	// generate window content
	w.SetContent(
		container.NewVBox(
			container.NewHBox(
				showTextfolderSelect,
				showSelectedFolder),
			folderSelectionButton,
			checkFolderContent))

	//show window when run
	w.ShowAndRun()
}
