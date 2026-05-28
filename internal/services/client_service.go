package services

import (
	"erp-go/internal/models"
	"erp-go/internal/repositories"
)

type ClientService struct {
	Repository *repositories.ClientRepository
}

// serviço para obter todos os clientes, delegando a responsabilidade para o repositório
func (s *ClientService) GetAll() ([]models.Client, error) {
	return s.Repository.GetAll()
}
