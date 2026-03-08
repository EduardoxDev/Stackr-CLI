package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/ui"
)

func FilesDelete(pos []string, flags map[string]string) {
	id := ResolveID(pos, "files delete")
	path := Flag(flags, "path", "")
	if path == "" {
		ui.Fail("Informe o caminho com --path")
		return
	}
	spin := ui.NewSpinner("Removendo " + path + "...")
	_, err := NewClient().Post(
		fmt.Sprintf("/apps/%s/files/delete", id),
		map[string]string{"path": path},
	)
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("Removido: " + path)
}
