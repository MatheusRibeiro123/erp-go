# TODO - ERP em GO

## Clients CRUD

* [x] Criar client_repository.go
* [x] Criar client_service.go
* [x] Criar client_handler.go
* [x] Criar rotas de clients
* [x] Testar CRUD de clients

### ClientRepository concluído

* [x] GetAll()
* [x] GetByID()
* [x] Create()
* [x] Update()
* [x] Delete()

### Melhorias Clients

#### Atualizações

* [x] Estudar implementação de PATCH /clients/:id
* [x] Implementar PATCH /clients/:id
* [ ] Testar PATCH /clients/:id

#### Tratamento de erros

* [ ] Implementar retorno 404 para cliente não encontrado
* [ ] Diferenciar erros de validação (400), recurso não encontrado (404) e erros internos (500)
* [ ] Melhorar tratamento de erros do PostgreSQL

#### Validações

* [ ] Adicionar validações de entrada (nome obrigatório, email válido, documento obrigatório)

#### Regras de negócio

* [ ] Verificar duplicidade de email
* [ ] Verificar duplicidade de documento

---

## Products CRUD

* [x] Criar product_repository.go
* [x] Criar product_service.go
* [x] Criar product_handler.go
* [x] Criar rotas de products
* [x] Testar CRUD de products

### ProductRepository concluído

* [x] GetAll()
* [x] GetByID()
* [x] Create()
* [x] Update()
* [x] Delete()

### Melhorias Products

#### Atualizações

* [ ] Implementar PATCH /products/:id
* [ ] Testar PATCH /products/:id

#### Validações

* [ ] Validar nome obrigatório
* [ ] Validar preço maior que zero
* [ ] Validar estoque não negativo

#### Regras de negócio

* [ ] Impedir preço negativo

---

## Melhorias REST

* [ ] Retornar 201 Created nos POSTs
* [ ] Retornar 204 No Content nos DELETEs
* [ ] Padronizar respostas da API

---

## Próximos módulos

* [ ] Categories
* [ ] Suppliers
* [ ] Orders
* [ ] Users

---

## Futuro

* [ ] Middleware
* [ ] Autenticação JWT
* [ ] Paginação
* [ ] Filtros de busca
