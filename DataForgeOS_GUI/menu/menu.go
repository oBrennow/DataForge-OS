package menu

import "fyne.io/fyne/v2"

func NewAppMenu() *fyne.MainMenu {
	producaoTurno := fyne.NewMenuItem("Produção", func() {
		// Criar um subtelinha para r3ealizar o cadastro da produção por turno
	})
	cadastroMaquina := fyne.NewMenuItem("Cadastro de Maquinas", func() {
		// Criar um subtelinha para r3ealizar o cadastro de maquinas
	})

	cadastroVendas := fyne.NewMenuItem("Cadastro de vendas", func() {
		// Criar um subtelinha para r3ealizar o cadastro das vendas
	})

	relatorioVendas := fyne.NewMenuItem("Relatorio de vendas", func() {
		// Criar um subtelinha para r3ealizar o cadastro das vendas
	})

	cadastroProduto := fyne.NewMenuItem("Cadastro de Produto", func() {
		// Criar um subtelinha para r3ealizar o cadastro das vendas
	})

	relatorioSaidaProdutos := fyne.NewMenuItem("Saida de Produtos", func() {
		// Criar um subtelinha para r3ealizar o cadastro das vendas
	})

	relatorioEntradaProdutos := fyne.NewMenuItem("Entrada de Produtos", func() {
		// Criar um subtelinha para r3ealizar o cadastro das vendas
	})

	menuProducao := fyne.NewMenu("Produção", producaoTurno, cadastroMaquina)
	menuVendas := fyne.NewMenu("Vendas", cadastroVendas, relatorioVendas)
	menuProdutos := fyne.NewMenu("Produtos", cadastroProduto, relatorioSaidaProdutos, relatorioEntradaProdutos)

	return fyne.NewMainMenu(menuProducao, menuVendas, menuProdutos)
}
