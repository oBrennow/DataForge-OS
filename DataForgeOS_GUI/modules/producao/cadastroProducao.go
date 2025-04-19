package producao

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/oBrenn0w/Data-Forge-OS/DataForgeOS_GUI/window"
)

// NewTurnoWindow monta a janela de cadastro de produção por turno
func CadastroProducao(a fyne.App) fyne.Window {
	w := a.NewWindow("Cadastro de Produção por turno")

	actions := window.FormActions{
		Include: func() {},
		Search:  func() {},
		Delete:  func() {},
		Confirm: func() {},
		Cancel:  func() {},
	}

	tb := window.FormToolbar(w, actions)

	userEntry := widget.NewEntry()
	userEntry.SetPlaceHolder("Usuario")

	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("Data (DD/MM/AAAA)")

	idLabel := widget.NewLabel("ID gerado automáticamente")

	turnoSelect := widget.NewSelect([]string{"A", "B", "C"}, func(string) {})

	maquinaSelect := widget.NewSelect([]string{"Maq1", "Maq2" /*Viram da API*/}, func(string) {})

	kgEntry := widget.NewEntry()
	kgEntry.SetPlaceHolder("KG Produzidos")

	paradaEntry := widget.NewEntry()
	paradaEntry.SetPlaceHolder("Tempo de Paradas")

	pacotesEntry := widget.NewEntry()
	pacotesEntry.SetPlaceHolder("Pacotes Produzidos")

	fardosLabel := widget.NewLabel("Fardos (calculado)")

	perdasEntry := widget.NewEntry()
	perdasEntry.SetPlaceHolder("Perdas em KG")

	obsEntry := widget.NewMultiLineEntry()
	obsEntry.SetPlaceHolder("Observação (até 200 caracteres)")

	form := container.NewGridWithColumns(2,
		widget.NewLabel("Usuário:"), userEntry,
		widget.NewLabel("Data:"), dateEntry,
		widget.NewLabel("ID:"), idLabel,
		widget.NewLabel("Turno:"), turnoSelect,
		widget.NewLabel("Máquina:"), maquinaSelect,
		widget.NewLabel("KG Produzidos:"), kgEntry,
		widget.NewLabel("Tempo Paradas:"), paradaEntry,
		widget.NewLabel("Pacotes:"), pacotesEntry,
		widget.NewLabel("Fardos:"), fardosLabel,
		widget.NewLabel("Perdas em KG:"), perdasEntry,
		widget.NewLabel("Observação:"), obsEntry,
	)

	content := container.NewVBox(
		tb,
		form,
	)
	w.SetContent(content)

	return w
}
