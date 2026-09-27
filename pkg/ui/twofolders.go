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

var twoFWindow fyne.Window
var srcdupDirectory binding.String = binding.NewString()

func twoFolderCheck(srcDir string, srcdupDir string) []string {
	srcFileList := check.Createfilelist(srcDir)
	dupsrcFileList := check.Createfilelist(srcdupDir)
	srcListMapHash := check.Createmapfilehash(srcFileList)
	dupsrcListMapHash := check.Createmapfilehash(dupsrcFileList)
	sFiles := check.TwoFolderDup(srcListMapHash, dupsrcListMapHash)
	log.Printf("same file list: %v", sFiles)
	return sFiles
}

func twoFolderWindow() {

	twoFWindow := FilecheckApp.NewWindow("Filecheck_2_folders")
	twoFWindow.RequestFocus()
	twoFWindow.Resize(WindowSize)
	twoFWindow.SetTitle("Filecheck_2_dossiers")

	// create button to select first source folder
	srcFolderSelectionButton := widget.NewButton("Selection", func() {
		dialog.ShowFolderOpen(func(srcDir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, twoFWindow)
				return
			}
			srcDirectory.Set(srcDir.Path())
			log.Println("selected folder in folderSelectionButton: ", srcDir.Path())
		}, twoFWindow)
	})

	// create button to select second source folder
	srcdupFolderSelectionButton := widget.NewButton("Selection", func() {
		dialog.ShowFolderOpen(func(srcdupDir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, twoFWindow)
				return
			}
			srcdupDirectory.Set(srcdupDir.Path())
			log.Println("selected folder in srcdupfolderSelectionButton: ", srcdupDir.Path())
		}, twoFWindow)
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
		d := dialog.NewCustom("Recherche en cours", "Annuler", widget.NewProgressBarInfinite(), twoFWindow)
		d.Show()
		srcFolder, _ := srcDirectory.Get()
		srcdupFolder, _ := srcDirectory.Get()
		if srcdupFolder == "" {
			log.Printf("missing dup folder.")
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
	twoFWindow.SetContent(
		container.NewVBox(
			container.NewCenter(
				srcFolderSelectionButton,
			),
			srcShowTextFolderSelect,
			srcShowSelectedFolders,
			container.NewCenter(
				srcdupFolderSelectionButton,
			),
			srcdupShowTextFolderSelect,
			srcdupShowSelectedFolders,
			container.NewCenter(
				checkFoldersContent,
			),
			container.NewCenter(
				returnWindow(twoFWindow),
			),
			lbSameFiles,
			showDuplicatesList,
		),
	)

	// 	//show window when run
	twoFWindow.Show()
}
