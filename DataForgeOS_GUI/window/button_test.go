package window_test

import (
	"testing"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func TestMainButton_CallbackExecutates(t *testing.T) {
	clicked := false
	esperado := true

	myApp := app.New()
	win := CreateMainWindow(myApp) // já retorna v2.Window

	btn, ok := win.Content().(*widget.Button)
	if !ok {
		t.Fatal("conteudo da janela não é um botão")
	}

	btn.OnTapped = func() { clicked = true }

	test.Tap(btn)

	if !clicked {
		t.Errorf("esperado clicked == '%v', mas foi '%v'", esperado, clicked)
	}
}
