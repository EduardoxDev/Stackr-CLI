package commands

import (
	"fmt"
	"os"

	"github.com/stackr-lat/cli/internal/config"
	"github.com/stackr-lat/cli/internal/ui"
)

func Whoami() {
	token := config.GetToken()
	if token == "" {
		ui.Fail("Não autenticado.")
		ui.Hint("Execute: stackr login <TOKEN>")
		return
	}
	ui.Header("Sessão Atual")
	preview := token
	if len(token) > 16 {
		preview = token[:8] + "••••••••" + token[len(token)-4:]
	}
	ui.Label("Token", preview)
	if os.Getenv("STACKR_API_TOKEN") != "" {
		ui.Label("Fonte", "env STACKR_API_TOKEN")
	} else {
		ui.Label("Fonte", "~/.stackr/config.json")
	}
	ui.Label("Versão CLI", "v"+config.Version)
	if localCfg, localPath := config.FindLocalConfig(); localCfg != nil {
		ui.SectionTitle("App Local Detectado")
		ui.Label("Config", localPath)
		if localCfg.ID != "" {
			ui.Label("ID", localCfg.ID)
		}
		if localCfg.Name != "" {
			ui.Label("Nome", localCfg.Name)
		}
		if localCfg.Lang != "" {
			ui.Label("Linguagem", localCfg.Lang)
		}
	}
	fmt.Println()
}
