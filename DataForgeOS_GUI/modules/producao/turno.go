package producao

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// … outros imports
)

// NewTurnoWindow monta a janela de cadastro de produção por turno
func NewTurnoWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Produção por Turno")

	lbl := widget.NewLabel("Turno:")
	entry := widget.NewEntry()
	btn := widget.NewButton("Salvar", func() { /* handler futuro */ })

	w.SetContent(container.NewVBox(lbl, entry, btn))
	return w
}
