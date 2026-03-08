package commands

import (
	"github.com/stackr-lat/cli/internal/config"
	"github.com/stackr-lat/cli/internal/ui"
)

func Logout() {
	cfg := config.Load()
	cfg.Token = ""
	_ = config.Save(cfg)
	ui.Ok("Sessão encerrada.")
}
