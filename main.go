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

var windowSize = fyne.Size{Width: 800, Height: 600}
var directory binding.String = binding.NewString()
var sameFileListPointer binding.DataList

func oneFolderCheck(dir string) []string {
	fileList := createfilelist(dir)
	fListMapHash := createmapfilehash(fileList)
	sameFileList := comparemap(fListMapHash)
	log.Printf("same file list: %v", sameFileList)
	return sameFileList
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
		sameFileList := oneFolderCheck(folder)
		d.Hide()
		log.Printf("same file list: %v", sameFileList)
		sameFileListPointer = binding.BindStringList(&sameFileList)
	},
	)

	//create a text box with the name of the folder selected
	showDuplicatesList := widget.NewListWithData(
		sameFileListPointer,
		func() fyne.CanvasObject {
			return widget.NewLabel("Liste des fichiers en double dans le dossier: ")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	// generate window content
	w.SetContent(
		container.NewVBox(
			container.NewHBox(
				showTextfolderSelect,
				showSelectedFolder),
			folderSelectionButton,
			checkFolderContent,
			showDuplicatesList,
			//container.NewHBox(
			//	showTextDuplicates,
			//	showDuplicatesList),
		))

	//show window when run
	w.ShowAndRun()
}
