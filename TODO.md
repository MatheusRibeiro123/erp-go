# TODO - ERP EM GO

---

# 🟢 CLIENTS CRUD

* [x] Criar client_repository.go  
* [x] Criar client_service.go  
* [x] Criar client_handler.go  
* [x] Criar rotas de clients  
* [x] Testar CRUD de clients  

## Repository (Clients)

* [x] GetAll()  
* [x] GetByID()  
* [x] Create()  
* [x] Update()  
* [x] Delete()  

---

## 🔥 Melhorias Clients (FASE ATUAL)

### Atualizações

* [x] Estudar implementação de PATCH /clients/:id  
* [x] Implementar PATCH /clients/:id  
* [x] Testar PATCH /clients/:id  

---

### 🚨 Tratamento de erros (EVOLUÇÃO PROFISSIONAL)

* [x] Implementar retorno 404 para cliente não encontrado  
* [x] Diferenciar erros de validação (400), recurso não encontrado (404) e erros internos (500)  
* [ ] Melhorar tratamento de erros do PostgreSQL  
* [ ] Criar padrão central de erros da API (Error Handler global)  
* [ ] Criar estrutura padrão de resposta da API (success/error wrapper)  

---

### ✅ Validações

* [ ] Adicionar validações de entrada (DTO + binding Gin)  
  - nome obrigatório  
  - email válido  
  - documento obrigatório  

---

### 🧠 Regras de negócio

* [ ] Verificar duplicidade de email  
* [ ] Verificar duplicidade de documento  

---

# 🟡 PRODUCTS CRUD

* [x] Criar product_repository.go  
* [x] Criar product_service.go  
* [x] Criar product_handler.go  
* [x] Criar rotas de products  
* [x] Testar CRUD de products  

## Repository (Products)

* [x] GetAll()  
* [x] GetByID()  
* [x] Create()  
* [x] Update()  
* [x] Delete()  

---

## 🔥 Melhorias Products (MESMA EVOLUÇÃO)

### Atualizações

* [x] Implementar PATCH /products/:id  
* [x] Testar PATCH /products/:id  

---

### 🚨 Tratamento de erros

* [ ] Implementar retorno 404 para produto não encontrado  
* [ ] Diferenciar erros de validação (400), recurso não encontrado (404), erros internos (500)  
* [ ] Aplicar padrão central de erros da API (igual Clients)  

---

### ✅ Validações

* [ ] Validar nome obrigatório  
* [ ] Validar preço maior que zero  
* [ ] Validar estoque não negativo  

---

### 🧠 Regras de negócio

* [ ] Impedir preço negativo  

---

# 🟣 PADRONIZAÇÃO DA API (FASE PROFISSIONAL)

* [ ] Padronizar respostas da API (response envelope)
  - success + data  
  - success + error  

* [ ] Retornar 201 Created nos POSTs  
* [ ] Retornar 204 No Content nos DELETEs  
* [ ] Garantir GET sempre consistente (lista vazia ao invés de null)  

---

# 🔵 BASE DE ARQUITETURA (REFORÇO PROFISSIONAL)

* [ ] Criar error handler global (middleware ou helper central)  
* [ ] Criar helper de response (JSON padrão)  
* [ ] Melhorar organização de erros do repository → service → handler  
* [ ] Separar erros de domínio vs erros de infraestrutura  

---

# 🟠 PRÓXIMOS MÓDULOS

* [ ] Categories (PRÓXIMO - ideal agora)  
* [ ] Suppliers  
* [ ] Orders  
* [ ] Users  

---

# 🧩 FUTURO (NÍVEL MERCADO)

* [ ] Middleware (logging, recovery)  
* [ ] Autenticação JWT  
* [ ] Roles (Admin/User)  
* [ ] Paginação  
* [ ] Filtros de busca  
* [ ] Docker + Docker Compose  
* [ ] Migrations  
* [ ] Testes unitários  