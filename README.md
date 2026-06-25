# ERP em Go

Sistema ERP backend desenvolvido em Go com foco em aprendizado de arquitetura backend profissional, APIs REST, PostgreSQL e boas práticas utilizadas no mercado.

O projeto simula a construção de um sistema ERP real, aplicando conceitos como separação de responsabilidades, tratamento de erros entre camadas, DTOs e organização de código escalável.

---

## 🚀 Tecnologias

* Go (Golang)
* PostgreSQL
* Gin
* Git / GitHub

---

## 🏗️ Arquitetura

O projeto utiliza arquitetura em camadas:


Handler → Service → Repository → PostgreSQL


### Responsabilidades

* Handler → Recebe requisições HTTP e retorna respostas
* Service → Regras de negócio
* Repository → Acesso ao banco de dados
* PostgreSQL → Persistência dos dados

---

## 📁 Estrutura do projeto

```bash
erp-go/
│
├── internal/
│ ├── handlers/
│ ├── services/
│ ├── repositories/
│ ├── dto/
│ ├── models/
│ ├── routes/
│ └── database/
│
├── main.go
├── go.mod
└── go.sum
```


---

## 📌 Objetivos do projeto

Este projeto está sendo desenvolvido para estudar:

* Desenvolvimento backend com Go
* Construção de APIs REST profissionais
* Arquitetura em camadas
* Injeção de dependências
* PostgreSQL
* Organização de projetos escaláveis
* Tratamento de erros entre camadas
* DTOs e validação de entrada
* Boas práticas de mercado

---

## ⚙️ Conceitos aplicados

* Arquitetura em camadas
* Separação de responsabilidades
* DTOs para entrada de dados
* Injeção de dependências
* Tratamento de erros com errors.Is
* Tradução de erros SQL → domínio
* Uso de PATCH com ponteiros
* Uso de RowsAffected() para updates/deletes
* Status HTTP corretos (200, 201, 400, 404, 500)

---

## ✅ Funcionalidades implementadas

### 👤 Clients

CRUD completo de clientes

* GET /clients → listar todos os clientes
* GET /clients/:id → buscar cliente por ID
* POST /clients → criar cliente
* PUT /clients/:id → atualizar cliente
* PATCH /clients/:id → atualização parcial
* DELETE /clients/:id → remover cliente
* Tratamento de erros HTTP (400, 404, 500)
* Uso de DTOs para entrada de dados

---

### 📦 Products

CRUD completo de produtos

* GET /products → listar todos os produtos
* GET /products/:id → buscar produto por ID
* POST /products → criar produto
* PUT /products/:id → atualizar produto
* PATCH /products/:id → atualização parcial
* DELETE /products/:id → remover produto

---

## 📚 Status do projeto

🚧 Em desenvolvimento ativo

### Progresso atual

* Clients CRUD → concluído
* Products CRUD → concluído
* Tratamento de erros → implementado em Clients e em expansão para Products
* Validações de negócio → pendente
* Categories → próximo módulo

---

## 🔥 Próximos passos

* Implementar validações completas (Clients e Products)
* Regras de negócio (duplicidade de email/documento)
* Padronização de responses da API
* Módulo Categories
* Autenticação JWT
* Paginação e filtros de busca

---

## 🎯 Observação

Este projeto não é apenas um CRUD simples. Ele está sendo construído com foco em arquitetura, entendimento profundo de fluxo de dados e boas práticas de backend utilizadas em sistemas reais.