package commands

import (
	"fmt"
	"strings"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func DBCreate(flags map[string]string) {
	name := Flag(flags, "name", "")
	engine := Flag(flags, "engine", "")
	if name == "" || engine == "" {
		ui.Fail("Campos obrigatórios: --name e --engine")
		ui.Hint("Engines: postgresql, mysql, mongodb, redis")
		return
	}
	valid := map[string]bool{"postgresql": true, "mysql": true, "mongodb": true, "redis": true}
	if !valid[strings.ToLower(engine)] {
		ui.Fail("Engine inválida: " + engine)
		ui.Hint("Use: postgresql, mysql, mongodb ou redis")
		return
	}
	body := map[string]interface{}{"name": name, "engine": engine}
	if mem := FlagInt(flags, "memory", 0); mem > 0 {
		body["memoryMb"] = mem
	}
	spin := ui.NewSpinner(fmt.Sprintf("Provisionando %s %s...", ui.EngineBadge(engine), name))
	var db api.Database
	data, err := NewClient().Post("/databases", body)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &db); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("Database criado! 🎉")
	ui.Header("Database: " + ui.Bold + db.Name + ui.Reset)
	ui.Label("ID", db.ID)
	ui.Label("Engine", ui.EngineBadge(db.Engine))
	ui.Label("Status", ui.StatusBadge(db.Status))
	ui.Label("Host", db.Host)
	ui.Label("Porta", fmt.Sprintf("%d", db.Port))
	ui.Label("Database", db.Database)
	ui.Label("Usuário", db.Username)
	ui.LabelSecret("Senha", db.Password)
	fmt.Printf("\n  %sConnection String:%s\n  %s%s%s\n\n", ui.Bold, ui.Reset, ui.Cyan, db.ConnectionString, ui.Reset)
	ui.Hint("Guarde as credenciais! A senha não será exibida novamente.")
	fmt.Println()
}
