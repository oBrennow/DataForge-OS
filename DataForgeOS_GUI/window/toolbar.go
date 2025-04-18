package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type FormActions struct {
	Include func()
	Search  func()
	Delete  func()
	Confirm func()
	Cancel  func()
}

func FormToolbar(w fyne.Window, actions FormActions) *widget.Toolbar {
	tb := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), actions.Include),
		widget.NewToolbarAction(theme.SearchIcon(), actions.Search),
		widget.NewToolbarAction(theme.DeleteIcon(), actions.Delete),
		widget.NewToolbarAction(theme.ConfirmIcon(), actions.Confirm),
		widget.NewToolbarAction(theme.CancelIcon(), actions.Cancel),
	)

	shortcuts := []struct {
		key     fyne.KeyName
		handler func()
	}{
		{fyne.KeyF2, actions.Include},
		{fyne.KeyF5, actions.Search},
		{fyne.KeyF4, actions.Delete},
		{fyne.KeyF10, actions.Confirm},
		{fyne.KeyF6, actions.Cancel},
	}

	canvas := w.Canvas()
	for _, sc := range shortcuts {
		sc := sc
		cs := &desktop.CustomShortcut{
			KeyName:  sc.key,
			Modifier: 0,
		}
		canvas.AddShortcut(cs, func(_ fyne.Shortcut) {
			sc.handler()
		})
	}

	return tb
}
