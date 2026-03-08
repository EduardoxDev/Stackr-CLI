package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func DBInfo(pos []string) {
	if len(pos) == 0 {
		ui.Fail("Informe o ID: stackr db info <ID>")
		return
	}
	spin := ui.NewSpinner("Buscando detalhes...")
	var db api.Database
	data, err := NewClient().Get("/databases/" + pos[0])
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &db); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("")
	ui.Header("Database: " + ui.Bold + db.Name + ui.Reset)
	ui.Label("ID", db.ID)
	ui.Label("Engine", ui.EngineBadge(db.Engine))
	ui.Label("Status", ui.StatusBadge(db.Status))
	ui.Label("Memória", fmt.Sprintf("%d MB", db.MemoryMb))
	ui.Label("Host", db.Host)
	ui.Label("Porta", fmt.Sprintf("%d", db.Port))
	ui.Label("Database", db.Database)
	ui.Label("Usuário", db.Username)
	ui.LabelSecret("Senha", db.Password)
	ui.Label("Criado em", db.CreatedAt)
	fmt.Printf("\n  %sConnection String:%s\n  %s%s%s\n\n", ui.Bold, ui.Reset, ui.Cyan, db.ConnectionString, ui.Reset)
}
