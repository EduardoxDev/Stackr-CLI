package commands

import (
	"fmt"
	"strings"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func FilesRead(pos []string, flags map[string]string) {
	id := ResolveID(pos, "files read")
	path := Flag(flags, "path", "")
	if path == "" {
		ui.Fail("Informe o caminho com --path")
		return
	}
	spin := ui.NewSpinner("Lendo " + path + "...")
	var fc api.FileContent
	data, err := NewClient().Get(fmt.Sprintf("/apps/%s/files/content?path=%s", id, path))
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &fc); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("")
	ui.Header("Arquivo: " + path)
	for i, line := range strings.Split(fc.Content, "\n") {
		fmt.Printf("  %s%4d%s  %s\n", ui.Gray, i+1, ui.Reset, line)
	}
	fmt.Println()
}
