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

* [ ] Implementar retorno 404 para cliente não encontrado
* [ ] Adicionar validações de entrada (nome, email, documento)
* [ ] Melhorar Update para manter valores atuais quando um campo não for enviado
* [ ] Estudar implementação de PATCH /clients/:id
* [ ] Melhorar tratamento de erros do banco

---

## Products CRUD

* [ ] Criar product_repository.go
* [ ] Criar product_service.go
* [ ] Criar product_handler.go
* [ ] Criar rotas de products
* [ ] Testar CRUD de products
