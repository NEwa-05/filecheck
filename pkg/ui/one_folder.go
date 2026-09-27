package ui

import (
	"filecheck/pkg/check"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// var srcDirectory binding.String = binding.NewString()
// var sameFileList binding.StringList = binding.NewStringList()
// var oneFWindow fyne.Window

func oneFolderCheck(srcDir string) []string {
	fileList := check.Createfilelist(srcDir)
	fListMapHash := check.Createmapfilehash(fileList)
	sFiles := check.OneFolderDup(fListMapHash)
	log.Printf("same file list: %v", sFiles)
	return sFiles
}

func oneFolderWindow() {

	oneFWindow := FilecheckApp.NewWindow("Filecheck_1_folder")
	oneFWindow.RequestFocus()
	oneFWindow.Resize(WindowSize)
	oneFWindow.SetTitle("Filecheck_1_dossier")

	// create button to select first source folder
	srcFolderSelectionButton := widget.NewButton("Selection", func() {
		dialog.ShowFolderOpen(func(srcDir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, oneFWindow)
				return
			}
			srcDirectory.Set(srcDir.Path())
			log.Println("selected folder in folderSelectionButton: ", srcDir.Path())
		}, oneFWindow)
	})

	//create a text box with text
	srcShowTextFolderSelect := widget.NewLabel("Premier dossier selectionné: ")

	//create a text box with the name of the folder selected
	srcShowSelectedFolders := widget.NewLabelWithData(srcDirectory)

	//create a text box with text
	lbSameFiles := widget.NewLabel("Liste des fichiers identiques: ")

	//button that will start the check process
	checkFoldersContent := widget.NewButton("Chercher les fichiers en doubles", func() {
		d := dialog.NewCustom("Recherche en cours", "Annuler", widget.NewProgressBarInfinite(), oneFWindow)
		d.Show()
		srcFolder, _ := srcDirectory.Get()
		log.Println("Selected folder in checkFolderContent: ", srcFolder)
		sameFiles := oneFolderCheck(srcFolder)
		d.Hide()
		log.Printf("same file list: %v", sameFiles)
		sameFileList.Set(sameFiles)
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
	oneFWindow.SetContent(
		container.NewVBox(
			container.NewCenter(
				srcFolderSelectionButton,
			),
			srcShowTextFolderSelect,
			srcShowSelectedFolders,
			container.NewCenter(
				checkFoldersContent,
			),
			container.NewCenter(
				returnWindow(oneFWindow),
			),
			lbSameFiles,
			showDuplicatesList,
		),
	)
	//show window when run
	oneFWindow.Show()
}
