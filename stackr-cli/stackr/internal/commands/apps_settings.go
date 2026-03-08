package commands

import (
	"fmt"
	"strconv"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func AppsSettings(pos []string, flags map[string]string) {
	id := ResolveID(pos, "apps settings")
	body := map[string]interface{}{}
	if v, ok := flags["name"]; ok {
		body["name"] = v
	}
	if v, ok := flags["memory"]; ok {
		if n, err := strconv.Atoi(v); err == nil {
			body["memoryMb"] = n
		}
	}
	if v, ok := flags["command"]; ok {
		body["command"] = v
	}
	if len(body) == 0 {
		ui.Fail("Informe ao menos um campo: --name, --memory ou --command")
		return
	}
	spin := ui.NewSpinner("Atualizando configurações...")
	var app api.App
	data, err := NewClient().Patch("/apps/"+id+"/settings", body)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &app); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("Configurações atualizadas!")
	ui.Label("Nome", app.Name)
	ui.Label("Memória", fmt.Sprintf("%d MB", app.MemoryMb))
	fmt.Println()
}
