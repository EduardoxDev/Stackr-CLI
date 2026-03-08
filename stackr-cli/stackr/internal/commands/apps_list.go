package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func AppsList() {
	spin := ui.NewSpinner("Buscando apps...")
	var apps []api.App
	data, err := NewClient().Get("/apps")
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &apps); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("")
	if len(apps) == 0 {
		ui.Info("Nenhum app encontrado.")
		ui.Hint("Faça deploy com: stackr apps upload <arquivo.zip>")
		return
	}
	ui.Header(fmt.Sprintf("Seus Apps  %s(%d)%s", ui.Gray, len(apps), ui.Reset))
	rows := make([]ui.Row, len(apps))
	for i, a := range apps {
		id := a.ID
		if len(id) > 8 {
			id = id[:8] + "…"
		}
		rows[i] = ui.Row{
			"id":     ui.Dim + id + ui.Reset,
			"name":   ui.Bold + a.Name + ui.Reset,
			"lang":   a.Language,
			"status": ui.StatusBadge(a.Status),
			"ram":    a.RAM,
		}
	}
	ui.PrintTable(rows, []ui.Column{
		{Key: "id", Label: "ID"},
		{Key: "name", Label: "Nome"},
		{Key: "lang", Label: "Linguagem"},
		{Key: "status", Label: "Status"},
		{Key: "ram", Label: "RAM"},
	})
	fmt.Println()
}
