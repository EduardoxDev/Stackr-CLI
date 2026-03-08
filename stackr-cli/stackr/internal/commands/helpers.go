package commands

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/stackr-lat/cli/internal/api"
	"github.com/stackr-lat/cli/internal/config"
	"github.com/stackr-lat/cli/internal/ui"
)

func NewClient() *api.Client {
	token := config.GetToken()
	if token == "" {
		ui.Fail("Você não está autenticado.")
		ui.Hint("Execute: stackr login <SEU_TOKEN>")
		ui.Hint("Ou defina: export STACKR_API_TOKEN=sk_live_xxx")
		os.Exit(1)
	}
	return api.New(token)
}

func ResolveID(pos []string, cmd string) string {
	if len(pos) > 0 && pos[0] != "" {
		return pos[0]
	}
	id, cfgPath := config.FindLocalAppID()
	if id != "" {
		ui.Info(fmt.Sprintf("App detectado via %s%s%s", ui.Bold, cfgPath, ui.Reset))
		return id
	}
	ui.Fail("ID do app não encontrado.")
	ui.Hint(fmt.Sprintf("Passe o ID: stackr %s <ID>", cmd))
	ui.Hint("Ou rode dentro de um diretório com stackr.config")
	os.Exit(1)
	return ""
}

func Confirm(prompt, expected string) bool {
	fmt.Printf("\n  %s%s:%s ", ui.Bold, prompt, ui.Reset)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input) == expected
}

func ParseFlags(args []string) ([]string, map[string]string) {
	flags := make(map[string]string)
	var pos []string
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--") {
			key := args[i][2:]
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				flags[key] = args[i+1]
				i++
			} else {
				flags[key] = "true"
			}
		} else {
			pos = append(pos, args[i])
		}
	}
	return pos, flags
}

func Flag(flags map[string]string, key, def string) string {
	if v, ok := flags[key]; ok {
		return v
	}
	return def
}

func FlagInt(flags map[string]string, key string, def int) int {
	if v, ok := flags[key]; ok {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}
