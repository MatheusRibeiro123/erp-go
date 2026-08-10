package services

import (
	"erp-go/internal/dto"
	"erp-go/internal/models"
	"erp-go/internal/repositories"
)

type ClientService struct {
	Repository *repositories.ClientRepository
}

// serviço para obter todos os clientes, delegando a responsabilidade para o repositório
func (s *ClientService) GetAll(page int, limit int) ([]models.Client, error) {
	return s.Repository.GetAll(page, limit)
}

// serviço para obter um cliente por ID, delegando a responsabilidade para o repositório
func (s *ClientService) GetByID(id int) (models.Client, error) {
	return s.Repository.GetByID(id)
}

// serviço para criar um novo cliente, delegando a responsabilidade para o repositório
func (s *ClientService) Create(client models.Client) (int, error) {
	return s.Repository.Create(client)
}

// serviço para atualizar um cliente existente, delegando a responsabilidade para o repositório
func (s *ClientService) Update(id int, client models.Client) error {
	return s.Repository.Update(id, client)
}

// serviço para excluir um cliente, delegando a responsabilidade para o repositório
func (s *ClientService) Delete(id int) error {
	return s.Repository.Delete(id)
}

// serviço para atualizar parcialmente um cliente, delegando a responsabilidade para o repositório
func (s *ClientService) Patch(id int, input dto.PatchClientInput) error {
	client, err := s.Repository.GetByID(id)
	if err != nil {
		return err
	}

	if input.Name != nil {
		client.Name = *input.Name
	}
	if input.Email != nil {
		client.Email = *input.Email
	}
	if input.Phone != nil {
		client.Phone = *input.Phone
	}
	if input.Document != nil {
		client.Document = *input.Document
	}
	return s.Repository.Update(id, client)
}
