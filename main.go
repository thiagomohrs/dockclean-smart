package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"dockclean-smart/internal/docker"
	"dockclean-smart/internal/i18n"
)

func main() {
	lang := flag.String("lang", "", "Idioma / Language (pt | en)")
	flag.Parse()

	msg := i18n.DetectLanguage(*lang)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cli, err := docker.NovoCliente()
	if err != nil {
		log.Fatalf(msg.ErrDockerConn, err)
	}

	fmt.Println(msg.SearchingImages)
	imagens, err := cli.ListarImagensNaoUtilizadas(ctx)
	if err != nil {
		log.Fatalf(msg.ErrListImages, err)
	}

	if len(imagens) == 0 {
		fmt.Println(msg.NoUnusedFound)
		return
	}

	fmt.Printf(msg.UnusedFoundHeader, len(imagens))

	var tamanhoTotalBytes int64
	for _, img := range imagens {
		tamanhoTotalBytes += img.Size
		tamanhoMB := float64(img.Size) / 1024 / 1024
		tag := "<none>"
		if len(img.RepoTags) > 0 {
			tag = img.RepoTags[0]
		}

		idCurto := img.ID
		if len(idCurto) > 19 && idCurto[:7] == "sha256:" {
			idCurto = idCurto[7:19]
		} else if len(idCurto) > 12 {
			idCurto = idCurto[:12]
		}

		fmt.Printf(msg.ImageItemFormat, idCurto, tag, tamanhoMB)
	}

	tamanhoTotalMB := float64(tamanhoTotalBytes) / 1024 / 1024
	tamanhoTotalGB := tamanhoTotalMB / 1024
	fmt.Printf(msg.TotalReclaimable, tamanhoTotalMB, tamanhoTotalGB)
}
