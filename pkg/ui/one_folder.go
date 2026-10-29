package ui

import (
	"filecheck/pkg/check"
	"fmt"
	"log"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var oneFWindow fyne.Window

func oneFolderCheck(srcDir string) []string {
	srcFileList := check.Createfilelist(srcDir)
	srcListMapHash := check.Createmapfilehash(srcFileList)
	sFiles := check.OneFolderDup(srcListMapHash)
	log.Printf("same file list: %v", sFiles)
	return sFiles
}

func oneFolderWindow() {

	oneFWindow := FilecheckApp.NewWindow("Filecheck_1_folder")
	oneFWindow.RequestFocus()
	oneFWindow.Resize(WindowSize)
	oneFWindow.SetTitle("Filecheck_1_dossier")

	// create button to select first source folder
	srcFolderSelectionButton := widget.NewButton("Sélection du dossier à vérifier", func() {
		dialog.ShowFolderOpen(func(srcDir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, oneFWindow)
			}
			SrcDirectory.Set(srcDir.Path())
			log.Println("selected folder in folderSelectionButton: ", srcDir.Path())
		}, oneFWindow)
	})

	//create a text box with text
	srcShowTextFolderSelect := widget.NewLabel("Dossier à vérifier: ")

	//create a text box with the name of the folder selected
	srcShowSelectedFolders := widget.NewLabelWithData(SrcDirectory)

	//create a text box with text
	lbSameFiles := widget.NewLabel("Liste des fichiers identiques: ")

	showDuplicatesList := widget.NewCheckGroup(
		SameFiles,
		func(fileSelected []string) {
			fmt.Println("Selected files:", fileSelected)
			if len(fileSelected) >= 1 && fileSelected[0] != "tous" {
				check.ToRemoveFileList = fileSelected
				fmt.Println("Selected files:", fileSelected)
			} else if slices.Contains(fileSelected, "tous") {
				check.ToRemoveFileList = SameFiles
				fmt.Println("Selected files:", check.ToRemoveFileList)
			}
		},
	)

	suppressButton := widget.NewButton("Sélection des fichiers à Supprimer", func() {
		dialog.ShowConfirm(
			"Confirmation",
			"Voulez vous supprimer les doublons sélectionné ?",
			func(confirmed bool) {
				if confirmed {

				} else {
					widget.NewLabel("Status: Suppression annulée.")
					fmt.Println("User clicked No/Cancel")
				}
			},
			oneFWindow,
		)
	})

	suppressButton.Hide()

	//button that will start the check process
	checkFoldersContent := widget.NewButton("Chercher les fichiers en doubles", func() {
		d := dialog.NewCustom("Recherche en cours", "Annuler", widget.NewProgressBarInfinite(), oneFWindow)
		d.Show()
		srcFolder, _ := SrcDirectory.Get()
		log.Println("Selected folder in checkFolderContent: ", srcFolder)
		SameFiles = oneFolderCheck(srcFolder)
		d.Hide()
		log.Printf("same file list: %v", SameFiles)
		if len(SameFiles) != 0 {
			showDuplicatesList.Append("tous")
			for _, v := range SameFiles {
				showDuplicatesList.Append(v)
			}
			showDuplicatesList.Refresh()
			showDuplicatesList.Show()
			suppressButton.Show()
		}
	})

	if SrcDirectory == nil {
		checkFoldersContent.Hide()
	} else {
		checkFoldersContent.Show()
	}

	// set window content
	oneFWindow.SetContent(
		container.NewBorder(
			container.NewVBox(
				container.NewPadded(
					srcFolderSelectionButton,
				),
				srcShowTextFolderSelect,
				srcShowSelectedFolders,
			),
			container.NewVBox(
				container.NewPadded(
					returnWindow(oneFWindow),
				),
			),
			nil,
			nil,
			container.NewBorder(
				container.NewBorder(
					container.NewPadded(
						checkFoldersContent,
					),
					lbSameFiles,
					nil,
					nil,
					nil,
				),
				suppressButton,
				nil,
				nil,
				showDuplicatesList,
			),
		),
	)

	//show window when run
	oneFWindow.Show()
}
