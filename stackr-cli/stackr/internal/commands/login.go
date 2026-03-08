package commands

import (
	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/config"
	"github.com/stackr-lat/cli/internal/ui"
)

func Login(args []string) {
	if len(args) == 0 {
		ui.Fail("Token não informado.")
		ui.Hint("stackr login <SEU_TOKEN>")
		ui.Hint("Gere seu token em: https://stackr.lat/dashboard/settings")
		return
	}
	spin := ui.NewSpinner("Validando token...")
	client := api.New(args[0])
	data, err := client.Get("/apps")
	if err != nil || data == nil {
		spin.Fail("Token inválido: " + err.Error())
		return
	}
	cfg := config.Load()
	cfg.Token = args[0]
	if err := config.Save(cfg); err != nil {
		spin.Fail("Erro ao salvar: " + err.Error())
		return
	}
	spin.Stop("Autenticado com sucesso!")
	ui.Hint("Token salvo em ~/.stackr/config.json")
}
