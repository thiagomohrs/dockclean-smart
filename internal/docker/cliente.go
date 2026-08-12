package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

// DockerAPIClient define a interface de interação com o Docker SDK para permitir testes unitários sem dependência do daemon real.
type DockerAPIClient interface {
	ImageList(ctx context.Context, options image.ListOptions) ([]image.Summary, error)
}

// ClienteDocker encapsula a API do Docker.
type ClienteDocker struct {
	cli DockerAPIClient
}

// Client mantido para retrocompatibilidade com o código boilerplate inicial.
type Client = ClienteDocker

// NovoCliente inicializa a conexão com o Docker Daemon do ambiente local.
func NovoCliente() (*ClienteDocker, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("falha ao inicializar o cliente Docker: %w", err)
	}
	return &ClienteDocker{cli: cli}, nil
}

// NewClient mantido para retrocompatibilidade com o código boilerplate inicial.
func NewClient() (*Client, error) {
	return NovoCliente()
}

// NovoClienteComAPI permite a injeção de dependência da interface DockerAPIClient (para testes unitários).
func NovoClienteComAPI(api DockerAPIClient) *ClienteDocker {
	return &ClienteDocker{cli: api}
}

// ListarImagensNaoUtilizadas busca todas as imagens que não possuem contêineres associados (Containers == 0).
func (c *ClienteDocker) ListarImagensNaoUtilizadas(ctx context.Context) ([]image.Summary, error) {
	imagens, err := c.cli.ImageList(ctx, image.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var naoUtilizadas []image.Summary
	for _, img := range imagens {
		if img.Containers == 0 {
			naoUtilizadas = append(naoUtilizadas, img)
		}
	}

	return naoUtilizadas, nil
}

// ListUnusedImages mantido para retrocompatibilidade com o código boilerplate inicial.
func (c *ClienteDocker) ListUnusedImages(ctx context.Context) ([]image.Summary, error) {
	return c.ListarImagensNaoUtilizadas(ctx)
}
