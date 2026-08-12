package docker

import (
	"context"
	"errors"
	"testing"

	"github.com/docker/docker/api/types/image"
)

type MockDockerAPI struct {
	ImagensRetorno []image.Summary
	ErroRetorno    error
}

func (m *MockDockerAPI) ImageList(ctx context.Context, options image.ListOptions) ([]image.Summary, error) {
	if m.ErroRetorno != nil {
		return nil, m.ErroRetorno
	}
	return m.ImagensRetorno, nil
}

func TestListarImagensNaoUtilizadas(t *testing.T) {
	mockAPI := &MockDockerAPI{
		ImagensRetorno: []image.Summary{
			{ID: "sha256:111111111111", Containers: 0, RepoTags: []string{"alpine:latest"}, Size: 5000000},
			{ID: "sha256:222222222222", Containers: 2, RepoTags: []string{"nginx:latest"}, Size: 20000000},
			{ID: "sha256:333333333333", Containers: 0, RepoTags: []string{"<none>:<none>"}, Size: 15000000},
		},
	}

	cliente := NovoClienteComAPI(mockAPI)
	naoUtilizadas, err := cliente.ListarImagensNaoUtilizadas(context.Background())
	if err != nil {
		t.Fatalf("Erro inesperado ao listar imagens: %v", err)
	}

	if len(naoUtilizadas) != 2 {
		t.Fatalf("Esperava 2 imagens não utilizadas, mas obteve %d", len(naoUtilizadas))
	}

	if naoUtilizadas[0].ID != "sha256:111111111111" {
		t.Errorf("ID inesperado para a primeira imagem: %s", naoUtilizadas[0].ID)
	}

	if naoUtilizadas[1].ID != "sha256:333333333333" {
		t.Errorf("ID inesperado para a segunda imagem: %s", naoUtilizadas[1].ID)
	}
}

func TestListarImagensNaoUtilizadasErroAPI(t *testing.T) {
	mockAPI := &MockDockerAPI{
		ErroRetorno: errors.New("falha de conexão com o daemon"),
	}

	cliente := NovoClienteComAPI(mockAPI)
	_, err := cliente.ListarImagensNaoUtilizadas(context.Background())
	if err == nil {
		t.Error("Esperava erro ao listar imagens com API indisponível, mas obteve nil")
	}
}
