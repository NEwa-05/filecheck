package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
)

var windowSize = fyne.Size{Width: 800, Height: 480}

//var dialogSize = fyne.Size{Width: 780, Height: 460}

func fillInFolder(button *widget.Button, w fyne.Window) {
	fd := dialog.NewFileOpen(func(rc fyne.URIReadCloser, err error) {
		if err != nil {
			showSubmenuError(fmt.Errorf("Could not open file: %w.", err), w)
			return
		}
		if rc != nil {
			if in != nil {
				in.Close()
			}
			in = rc
			button.SetText(rc.URI().Name())
		} else {
			button.SetText("Select input file")
		}
	}, w)
	fd.Resize(dialogSize)
	fd.SetOnClosed(func() {
		w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) { handleSubmenuKey(e, w) })
	})
	w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) { handleDialogKey(e, w, fd) })
	fd.Show()
}

func dirForm(w fyne.Window) fyne.CanvasObject {
	inFolderLabel := widget.NewLabel("Séléction des dossiers:")
	var inFolder *widget.Button
	if inFolder == nil {
		inFolder = widget.NewButton("Dossier source !", nil)
	}
	inFolder.OnTapped = func() { fillInFolder(inFolder, w) }

	pad := strings.Repeat(" ", 32)
	submit := widget.NewButton(pad+"Valider"+pad, func() { validate(w) })
	form := container.New(layout.NewFormLayout(), inFolderLabel, inFolder)
	return container.NewPadded(container.NewCenter(container.NewVBox(form, submit)))

}

func main() {
	a := app.New()
	w := a.NewWindow("Filecheck")
	w.Resize(windowSize)
	dirForm(w)
	w.ShowAndRun()
}
