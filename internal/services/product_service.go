package services

import (
	"erp-go/internal/dto"
	"erp-go/internal/models"
	"erp-go/internal/repositories"
)

type ProductService struct {
	Repository *repositories.ProductRepository
}

// serviço para obter todos os produtos, delegando a responsabilidade para o repositório
func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.Repository.GetAll()
}

// serviço para obter um produto por ID, delegando a responsabilidade para o repositório
func (s *ProductService) GetByID(id int) (models.Product, error) {
	return s.Repository.GetByID(id)
}

// serviço para criar um novo produto, delegando a responsabilidade para o repositório
func (s *ProductService) Create(product models.Product) (int, error) {
	return s.Repository.Create(product)
}

// serviço para atualizar um produto existente, delegando a responsabilidade para o repositório
func (s *ProductService) Update(id int, product models.Product) error {
	return s.Repository.Update(id, product)
}

// serviço para excluir um produto, delegando a responsabilidade para o repositório
func (s *ProductService) Delete(id int) error {
	return s.Repository.Delete(id)
}

// serviço para atualizar parcialmente um produto

func (s *ProductService) Patch(id int, input dto.PatchProductInput) error {
	produto, err := s.Repository.GetByID(id)
	if err != nil {
		return err
	}

	if input.Name != nil {
		produto.Name = *input.Name
	}
	if input.Description != nil {
		produto.Description = *input.Description
	}
	if input.Price != nil {
		produto.Price = *input.Price
	}
	if input.StockQuantity != nil {
		produto.StockQuantity = *input.StockQuantity
	}

	return s.Repository.Update(id, produto)
}
