package window

import (
	"filecheck/pkg/filecheck"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var WindowSize = fyne.Size{Width: 800, Height: 600}

func oneFolderCheck(srcDir string) []string {
	fileList := filecheck.Createfilelist(srcDir)
	fListMapHash := filecheck.Createmapfilehash(fileList)
	sFiles := filecheck.OneFolderDup(fListMapHash)
	log.Printf("same file list: %v", sFiles)
	return sFiles
}

func CreateOneFolderWindow() {

	//create window
	a := app.NewWithID("APPID")
	onefolderw := a.NewWindow("Selection du Dossier")
	// duplistw := a.NewWindow("Liste des doublons")
	onefolderw.Resize(windowSize)

	// create button to select first source folder
	srcFolderSelectionButton := widget.NewButton("Selection", func() {
		dialog.ShowFolderOpen(func(srcDir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, onefolderw)
				return
			}
			srcDirectory.Set(srcDir.Path())
			log.Println("selected folder in folderSelectionButton: ", srcDir.Path())
		}, onefolderw)
	})

	//create a text box with text
	srcShowTextFolderSelect := widget.NewLabel("Premier dossier selectionné: ")

	//create a text box with the name of the folder selected
	srcShowSelectedFolders := widget.NewLabelWithData(srcDirectory)

	//button that will start the check process
	checkFoldersContent := widget.NewButton("Chercher les fichiers en doubles", func() {
		d := dialog.NewCustom("Recherche en cours", "Annuler", widget.NewProgressBarInfinite(), onefolderw)
		d.Show()
		srcFolder, _ := srcDirectory.Get()
		log.Println("Selected folder in checkFolderContent: ", srcFolder)
		sameFiles := oneFolderCheck(srcFolder)
		d.Hide()
		log.Printf("same file list: %v", sameFiles)
		sameFileList.Set(sameFiles)
	})

	// func(i binding.DataItem, o fyne.CanvasObject) {
	// 	o.(*widget.Label).Bind(i.(binding.String))
	// })

	// set window content
	onefolderw.SetContent(
		container.NewVBox(
			container.NewHBox(
				srcShowTextFolderSelect,
				srcShowSelectedFolders,
			),
			srcFolderSelectionButton,
			checkFoldersContent,
		),
	)

	//show window when run
	onefolderw.ShowAndRun()
}
