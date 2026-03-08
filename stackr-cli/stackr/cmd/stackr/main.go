package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/stackr-lat/cli/internal/commands"
	"github.com/stackr-lat/cli/internal/config"
	"github.com/stackr-lat/cli/internal/ui"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		ui.Banner(config.Version)
		printHelp()
		return
	}

	switch args[0] {
	case "help", "--help", "-h":
		ui.Banner(config.Version)
		printHelp()
	case "version", "--version", "-v":
		fmt.Printf("stackr v%s (%s/%s)\n", config.Version, runtime.GOOS, runtime.GOARCH)
	case "login":
		commands.Login(args[1:])
	case "logout":
		commands.Logout()
	case "whoami":
		commands.Whoami()
	case "update":
		commands.Update()
	case "apps":
		routeApps(args[1:])
	case "files":
		routeFiles(args[1:])
	case "db":
		routeDB(args[1:])
	default:
		ui.Fail("Comando desconhecido: " + args[0])
		ui.Hint("Use stackr help para ver todos os comandos")
		os.Exit(1)
	}
}

func routeApps(args []string) {
	if len(args) == 0 {
		printAppsHelp()
		return
	}
	pos, flags := commands.ParseFlags(args[1:])
	switch args[0] {
	case "list":
		commands.AppsList()
	case "info":
		commands.AppsInfo(pos)
	case "stats":
		commands.AppsStats(pos)
	case "logs":
		commands.AppsLogs(pos, flags)
	case "start":
		commands.AppsStart(pos)
	case "stop":
		commands.AppsStop(pos)
	case "restart":
		commands.AppsRestart(pos)
	case "rebuild":
		commands.AppsRebuild(pos)
	case "delete":
		commands.AppsDelete(pos)
	case "settings":
		commands.AppsSettings(pos, flags)
	case "upload":
		commands.AppsUpload(pos)
	case "backup":
		commands.AppsBackup(pos, flags)
	default:
		ui.Fail("Subcomando desconhecido: apps " + args[0])
		printAppsHelp()
	}
}

func routeFiles(args []string) {
	if len(args) == 0 {
		printFilesHelp()
		return
	}
	pos, flags := commands.ParseFlags(args[1:])
	switch args[0] {
	case "list":
		commands.FilesList(pos, flags)
	case "read":
		commands.FilesRead(pos, flags)
	case "write":
		commands.FilesWrite(pos, flags)
	case "mkdir":
		commands.FilesMkdir(pos, flags)
	case "delete":
		commands.FilesDelete(pos, flags)
	default:
		ui.Fail("Subcomando desconhecido: files " + args[0])
		printFilesHelp()
	}
}

func routeDB(args []string) {
	if len(args) == 0 {
		printDBHelp()
		return
	}
	pos, flags := commands.ParseFlags(args[1:])
	switch args[0] {
	case "list":
		commands.DBList()
	case "create":
		commands.DBCreate(flags)
	case "info":
		commands.DBInfo(pos)
	case "start":
		commands.DBStart(pos)
	case "stop":
		commands.DBStop(pos)
	case "restart":
		commands.DBRestart(pos)
	case "delete":
		commands.DBDelete(pos)
	case "logs":
		commands.DBLogs(pos, flags)
	case "stats":
		commands.DBStats(pos)
	default:
		ui.Fail("Subcomando desconhecido: db " + args[0])
		printDBHelp()
	}
}

func line(label, desc string) {
	fmt.Printf("    %s%-48s%s  %s%s%s\n", ui.Cyan, label, ui.Reset, ui.Gray, desc, ui.Reset)
}

func printHelp() {
	fmt.Printf("  %suso%s  stackr <comando> [subcomando] [flags]\n\n", ui.Bold, ui.Reset)

	ui.SectionTitle("autenticacao")
	line("stackr login <TOKEN>",   "salva o api token em ~/.stackr/config.json")
	line("stackr logout",          "encerra a sessao")
	line("stackr whoami",          "exibe sessao atual e app local detectado")

	ui.SectionTitle("apps")
	fmt.Printf("    %s  ID e opcional se houver stackr.config no diretorio%s\n\n", ui.Gray, ui.Reset)
	line("stackr apps list",                                   "lista todos os apps")
	line("stackr apps info [ID]",                              "detalhes do app")
	line("stackr apps stats [ID]",                             "cpu, ram e rede em tempo real")
	line("stackr apps logs [ID] [--tail N]",                   "logs com highlight por nivel")
	line("stackr apps start [ID]",                             "ligar app")
	line("stackr apps stop [ID]",                              "desligar app")
	line("stackr apps restart [ID]",                           "reiniciar app")
	line("stackr apps rebuild [ID]",                           "reconstruir container do zero")
	line("stackr apps delete [ID]",                            "deletar permanentemente")
	line("stackr apps settings [ID] --name --memory --command","atualizar configuracoes")
	line("stackr apps upload <arquivo.zip>",                   "deploy de novo app")
	line("stackr apps backup [ID] [--output arquivo.zip]",     "baixar backup do app")

	ui.SectionTitle("arquivos")
	line("stackr files list [ID] --path /app",             "listar com tamanhos")
	line("stackr files read [ID] --path /app/index.js",    "ver conteudo com linhas")
	line("stackr files write [ID] --path P --content C",   "escrever arquivo")
	line("stackr files write [ID] --path P --file local",  "enviar arquivo local")
	line("stackr files mkdir [ID] --path /app/data",       "criar diretorio")
	line("stackr files delete [ID] --path /app/temp",      "remover arquivo")

	ui.SectionTitle("databases")
	line("stackr db list",                                        "listar todos")
	line("stackr db create --name X --engine E [--memory MB]",   "provisionar  [ postgresql / mysql / mongodb / redis ]")
	line("stackr db info <ID>",                                   "detalhes + connection string")
	line("stackr db start <ID>",                                  "iniciar")
	line("stackr db stop <ID>",                                   "parar")
	line("stackr db restart <ID>",                                "reiniciar")
	line("stackr db delete <ID>",                                 "deletar permanentemente")
	line("stackr db logs <ID> [--tail N]",                        "logs do container")
	line("stackr db stats <ID>",                                  "metricas de uso")

	ui.SectionTitle("cli")
	line("stackr update",  "atualiza para a versao mais recente automaticamente")
	line("stackr version", "exibe versao e plataforma")
	fmt.Println()
	fmt.Printf("  %sdica:%s  export STACKR_API_TOKEN=sk_live_xxx  para autenticar via env\n\n", ui.Gray, ui.Reset)
}

func printAppsHelp() {
	ui.SectionTitle("stackr apps")
	fmt.Printf("    %slist  info  stats  logs  start  stop  restart  rebuild  delete  settings  upload  backup%s\n\n", ui.Gray, ui.Reset)
}

func printFilesHelp() {
	ui.SectionTitle("stackr files")
	fmt.Printf("    %slist  read  write  mkdir  delete%s\n\n", ui.Gray, ui.Reset)
}

func printDBHelp() {
	ui.SectionTitle("stackr db")
	fmt.Printf("    %slist  create  info  start  stop  restart  delete  logs  stats%s\n\n", ui.Gray, ui.Reset)
}
