package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

var windowSize = fyne.Size{Width: 800, Height: 480}
var directory binding.String = binding.NewString()

//func fillInFolder(w fyne.Window, h *widget.Label) {
//	dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
//		if err != nil {
//			dialog.ShowError(err, w)
//			return
//		}
//		directory := dir.Path()
//		log.Println("selected folder in fillInFolder: ", directory)
//		if err != nil {
//			dialog.ShowError(err, w)
//			return
//		}
//	}, w)
//}

//func checkInFolder(w fyne.Window, h *widget.Label, dirCheck string) {
//	log.Printf("selected folder checkInFolder: %s", directory)
//	d := dialog.NewCustom("Recherche en cours", "Annuler", widget.NewProgressBarInfinite(), w)
//	d.Show()
//}

func main() {
	a := app.NewWithID("APPID")
	w := a.NewWindow("Filecheck")
	w.Resize(windowSize)
	//folderWindow := widget.NewLabel("Selection du dossier: ")
	folderSelectionButton := widget.NewButton("Sélection", func() {
		//fillInFolder(w, folderWindow)
		dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			directory.Set(dir.Path())
			log.Println("selected folder in folderSelectionButton: ", dir.Path())
		}, w)
	})
	folderSelectionWindow := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Sélectionner le dossier: ", Widget: folderSelectionButton}},
		OnSubmit: func() {
			folder, _ := directory.Get()
			log.Println("Selected folder in folderSelectionWindow: ", folder)
		},
	}
	//checkFolderContent :=  container.NewVBox(directory, widget.NewButton("Chercher les fichiers en doubles", func() { checkInFolder(w, folderWindow, folderSelectionWindow) })

	w.SetContent(
		container.NewVBox(
			//checkFolderContent,
			folderSelectionWindow,
		),
	)
	//w.SetContent(container.NewVBox(
	//	folderWindow,
	//	widget.NewButton("Sélection", func() {
	//		fillInFolder(w, folderWindow)
	//	}),
	//	widget.NewButton("Chercher les fichiers en doubles", func() {
	//		checkInFolder(w, folderWindow, directory)
	//	}),
	//))

	//selectFolder := widget.NewButton("Sélection", func() { fillInFolder(w, folderWindow) })
	//w.SetContent(
	//	container.NewVSplit(
	//		selectFolder,
	//		checkFolderContent,
	//	),
	//)
	w.ShowAndRun()
}
