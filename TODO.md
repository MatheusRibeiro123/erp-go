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

#### Tratamento de erros
* [ ] Implementar retorno 404 para cliente não encontrado
* [ ] Diferenciar erros de validação (400), recurso não encontrado (404) e erros internos (500)
* [ ] Melhorar tratamento de erros do PostgreSQL

#### Validações
* [ ] Adicionar validações de entrada (nome obrigatório, email válido, documento obrigatório)

#### Atualizações
* [ ] Melhorar Update para manter valores atuais quando um campo não for enviado
* [ ] Estudar implementação de PATCH /clients/:id

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
