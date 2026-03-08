package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func AppsInfo(pos []string) {
	id := ResolveID(pos, "apps info")
	spin := ui.NewSpinner("Buscando detalhes...")
	var app api.App
	data, err := NewClient().Get("/apps/" + id)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &app); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("")
	ui.Header("App: " + ui.Bold + app.Name + ui.Reset)
	ui.Label("ID", app.ID)
	ui.Label("Nome", app.Name)
	ui.Label("Linguagem", app.Language)
	ui.Label("Tipo", app.Type)
	ui.Label("Status", ui.StatusBadge(app.Status))
	ui.Label("Memória", fmt.Sprintf("%d MB", app.MemoryMb))
	if app.OOMKilled {
		ui.Label("OOM Killed", ui.Red+ui.Bold+"⚠ Sim — app encerrado por falta de memória"+ui.Reset)
		ui.Hint("Aumente a memória: stackr apps settings --memory 512")
	} else {
		ui.Label("OOM Killed", ui.Green+"Não"+ui.Reset)
	}
	fmt.Println()
}
