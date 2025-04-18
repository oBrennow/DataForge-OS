package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateMainWindow(app fyne.App) fyne.Window {
	win := app.NewWindow("DataForgeOS")

	btn := widget.NewButton("Clique aqui", func() {
		println("Deu certo!")
	})

	win.SetContent(btn)
	win.Resize(fyne.NewSize(400, 300))

	return win
}
