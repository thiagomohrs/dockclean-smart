package i18n

import (
	"os"
	"testing"
)

func TestDetectLanguageExplicitFlag(t *testing.T) {
	msgEn := DetectLanguage("en")
	if msgEn.NoUnusedFound != "No unused images found." {
		t.Errorf("Esperava tradução em inglês, obteve: %q", msgEn.NoUnusedFound)
	}

	msgPt := DetectLanguage("pt")
	if msgPt.NoUnusedFound != "Nenhuma imagem sem uso encontrada." {
		t.Errorf("Esperava tradução em português, obteve: %q", msgPt.NoUnusedFound)
	}
}

func TestDetectLanguageEnvFallback(t *testing.T) {
	origLang := os.Getenv("LANG")
	defer os.Setenv("LANG", origLang)

	os.Setenv("LANG", "en_US.UTF-8")
	msg := DetectLanguage("")
	if msg.NoUnusedFound != "No unused images found." {
		t.Errorf("Esperava fallback para inglês via LANG env, obteve: %q", msg.NoUnusedFound)
	}

	os.Setenv("LANG", "pt_BR.UTF-8")
	msgPt := DetectLanguage("")
	if msgPt.NoUnusedFound != "Nenhuma imagem sem uso encontrada." {
		t.Errorf("Esperava fallback para português via LANG env, obteve: %q", msgPt.NoUnusedFound)
	}
}
