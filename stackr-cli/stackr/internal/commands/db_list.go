package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func DBList() {
	spin := ui.NewSpinner("Buscando databases...")
	var dbs []api.Database
	data, err := NewClient().Get("/databases")
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &dbs); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("")
	if len(dbs) == 0 {
		ui.Info("Nenhum database encontrado.")
		ui.Hint("Crie com: stackr db create --name meu-db --engine postgresql")
		return
	}
	ui.Header(fmt.Sprintf("Databases  %s(%d)%s", ui.Gray, len(dbs), ui.Reset))
	rows := make([]ui.Row, len(dbs))
	for i, d := range dbs {
		id := d.ID
		if len(id) > 8 {
			id = id[:8] + "…"
		}
		rows[i] = ui.Row{
			"id":     ui.Dim + id + ui.Reset,
			"name":   ui.Bold + d.Name + ui.Reset,
			"engine": d.Engine,
			"status": ui.StatusBadge(d.Status),
			"mem":    fmt.Sprintf("%d MB", d.MemoryMb),
		}
	}
	ui.PrintTable(rows, []ui.Column{
		{Key: "id", Label: "ID"},
		{Key: "name", Label: "Nome"},
		{Key: "engine", Label: "Engine"},
		{Key: "status", Label: "Status"},
		{Key: "mem", Label: "Memória"},
	})
	fmt.Println()
}
