package i18n

import (
	"os"
	"strings"
)

// Messages define os textos traduzíveis da aplicação.
type Messages struct {
	SearchingImages   string
	ErrDockerConn     string
	ErrListImages     string
	NoUnusedFound     string
	UnusedFoundHeader string
	ImageItemFormat   string
	TotalReclaimable  string
	FlagLang          string
	FlagDryRun        string
}

var ptMessages = Messages{
	SearchingImages:   "Buscando imagens não utilizadas na sua máquina...",
	ErrDockerConn:     "Erro ao conectar ao Docker Daemon: %v",
	ErrListImages:     "Erro ao listar imagens: %v",
	NoUnusedFound:     "Nenhuma imagem sem uso encontrada.",
	UnusedFoundHeader: "\nForam encontradas %d imagens não associadas a contêineres rodando:\n",
	ImageItemFormat:   "- ID: %s | Tag: %s | Tamanho: %.2f MB\n",
	TotalReclaimable:  "\nEspaço total recuperável: %.2f MB (%.2f GB)\n",
	FlagLang:          "Idioma das mensagens do CLI (pt|en)",
	FlagDryRun:        "Modo de simulação (simulacao)",
}

var enMessages = Messages{
	SearchingImages:   "Searching for unused Docker images on your system...",
	ErrDockerConn:     "Error connecting to Docker Daemon: %v",
	ErrListImages:     "Error listing images: %v",
	NoUnusedFound:     "No unused images found.",
	UnusedFoundHeader: "\nFound %d images not associated with running containers:\n",
	ImageItemFormat:   "- ID: %s | Tag: %s | Size: %.2f MB\n",
	TotalReclaimable:  "\nTotal reclaimable space: %.2f MB (%.2f GB)\n",
	FlagLang:          "CLI output language (en|pt)",
	FlagDryRun:        "Simulation mode (dry-run)",
}

// DetectLanguage detecta o idioma a partir da flag informada ou da variável de ambiente LANG/LC_ALL.
func DetectLanguage(langFlag string) Messages {
	langFlag = strings.ToLower(strings.TrimSpace(langFlag))
	if langFlag == "en" {
		return enMessages
	}
	if langFlag == "pt" {
		return ptMessages
	}

	envLang := strings.ToLower(os.Getenv("LANG"))
	envLcAll := strings.ToLower(os.Getenv("LC_ALL"))
	if strings.HasPrefix(envLang, "en") || strings.HasPrefix(envLcAll, "en") {
		return enMessages
	}

	return ptMessages
}
