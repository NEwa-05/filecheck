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
var srcDirectory binding.String = binding.NewString()
var srcdupDirectory binding.String = binding.NewString()
var sameFileList binding.StringList = binding.NewStringList()

func oneFolderCheck(srcDir string) []string {
	fileList := createfilelist(srcDir)
	fListMapHash := createmapfilehash(fileList)
	sFiles := oneFolderDup(fListMapHash)
	log.Printf("same file list: %v", sFiles)
	return sFiles
}

func twoFolderCheck(srcDir string, srcdupDir string) []string {
	srcFileList := createfilelist(srcDir)
	dupsrcFileList := createfilelist(srcdupDir)
	srcListMapHash := createmapfilehash(srcFileList)
	dupsrcListMapHash := createmapfilehash(dupsrcFileList)
	sFiles := twoFolderDup(srcListMapHash, dupsrcListMapHash)
	log.Printf("same file list: %v", sFiles)
	return sFiles
}

func main() {

	//create window
	a := app.NewWithID("APPID")
	w := a.NewWindow("Filecheck")
	w.Resize(windowSize)

	// create button to select first source folder
	srcFolderSelectionButton := widget.NewButton("Sélection", func() {
		dialog.ShowFolderOpen(func(srcDir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			srcDirectory.Set(srcDir.Path())
			log.Println("selected folder in folderSelectionButton: ", srcDir.Path())
		}, w)
	})

	// create button to select second source folder
	srcdupFolderSelectionButton := widget.NewButton("Sélection", func() {
		dialog.ShowFolderOpen(func(srcdupDir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			srcdupDirectory.Set(srcdupDir.Path())
			log.Println("selected folder in folderSelectionButton: ", srcdupDir.Path())
		}, w)
	})

	//create a text box with text
	srcShowTextFolderSelect := widget.NewLabel("Premier dossier selectionné: ")

	//create a text box with the name of the folder selected
	srcShowSelectedFolders := widget.NewLabelWithData(srcDirectory)

	//create a text box with text
	srcdupShowTextFolderSelect := widget.NewLabel("Second dossier selectionné: ")

	//create a text box with the name of the folder selected
	srcdupShowSelectedFolders := widget.NewLabelWithData(srcdupDirectory)

	//create a text box with text
	lbSameFiles := widget.NewLabel("Liste des fichiers identiques: ")

	//button that will start the check process
	checkFoldersContent := widget.NewButton("Chercher les fichiers en doubles", func() {
		d := dialog.NewCustom("Recherche en cours", "Annuler", widget.NewProgressBarInfinite(), w)
		d.Show()
		srcFolder, _ := srcDirectory.Get()
		srcdupFolder, _ := srcDirectory.Get()
		if srcdupFolder == "" {
			log.Println("Selected folder in checkFolderContent: ", srcFolder)
			sameFiles := oneFolderCheck(srcFolder)
			d.Hide()
			log.Printf("same file list: %v", sameFiles)
			sameFileList.Set(sameFiles)
		} else {
			log.Println("Selected folder in checkFolderContent: ", srcFolder, srcdupFolder)
			sameFiles := twoFolderCheck(srcFolder, srcdupFolder)
			d.Hide()
			log.Printf("same file list: %v", sameFiles)
			sameFileList.Set(sameFiles)
		}
	})

	// create a text box with the name of the folder selected
	showDuplicatesList := widget.NewListWithData(
		sameFileList,
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},

		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	// set window content
	w.SetContent(
		container.NewBorder(
			container.NewVBox(
				container.NewHBox(
					srcShowTextFolderSelect,
					srcShowSelectedFolders,
					srcdupShowTextFolderSelect,
					srcdupShowSelectedFolders,
				),
				srcFolderSelectionButton,
				srcdupFolderSelectionButton,
				checkFoldersContent,
				lbSameFiles),
			nil,
			nil,
			nil,
			showDuplicatesList))

	//show window when run
	w.ShowAndRun()
}
