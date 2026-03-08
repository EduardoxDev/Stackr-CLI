package commands

import (
	"fmt"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/ui"
)

func FilesList(pos []string, flags map[string]string) {
	id := ResolveID(pos, "files list")
	path := Flag(flags, "path", "/app")
	spin := ui.NewSpinner("Listando " + path + "...")
	var files []api.FileEntry
	data, err := NewClient().Get(fmt.Sprintf("/apps/%s/files?path=%s", id, path))
	if err != nil {
		spin.Fail(err.Error())
		return
	}
	if err := api.Decode(data, &files); err != nil {
		spin.Fail(err.Error())
		return
	}
	spin.Stop("")
	ui.Header("Arquivos em " + path)
	if len(files) == 0 {
		ui.Info("Diretório vazio.")
		return
	}
	for _, f := range files {
		if f.Type == "directory" {
			fmt.Printf("  %s📁 %s%s%s/\n", ui.Blue+ui.Bold, ui.Cyan, f.Name, ui.Reset)
			continue
		}
		size := ""
		switch {
		case f.Size < 1024:
			size = fmt.Sprintf("  %s%d B%s", ui.Gray, f.Size, ui.Reset)
		case f.Size < 1024*1024:
			size = fmt.Sprintf("  %s%.1f KB%s", ui.Gray, float64(f.Size)/1024, ui.Reset)
		default:
			size = fmt.Sprintf("  %s%.1f MB%s", ui.Gray, float64(f.Size)/1024/1024, ui.Reset)
		}
		fmt.Printf("  %s📄%s %s%s\n", ui.Gray, ui.Reset, f.Name, size)
	}
	fmt.Println()
}
