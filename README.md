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

```text
Handler → Service → Repository → PostgreSQL
```

### Responsabilidades

* Handler → Recebe requisições HTTP, valida entradas e retorna respostas.
* Service → Contém as regras de negócio da aplicação.
* Repository → Responsável pelo acesso ao banco de dados.
* PostgreSQL → Persistência dos dados.

---

## 📁 Estrutura do projeto

```bash
erp-go/
│
├── internal/
│   ├── apperrors/
│   ├── handlers/
│   ├── services/
│   ├── repositories/
│   ├── dto/
│   ├── models/
│   ├── routes/
│   └── database/
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
* Tratamento de erros com `errors.Is`
* Tradução de erros do PostgreSQL para erros da aplicação
* Error Handler global para respostas HTTP
* Centralização do tratamento de erros
* Uso de PATCH com ponteiros
* Uso de `RowsAffected()` para updates e deletes
* Status HTTP corretos (200, 201, 400, 404, 409 e 500)

---

## ✅ Funcionalidades implementadas

### 👤 Clients

CRUD completo de clientes

* GET `/clients` → listar todos os clientes
* GET `/clients/:id` → buscar cliente por ID
* POST `/clients` → criar cliente
* PUT `/clients/:id` → atualizar cliente
* PATCH `/clients/:id` → atualização parcial
* DELETE `/clients/:id` → remover cliente
* DTOs para entrada de dados
* Tratamento centralizado de erros
* Tradução de erros do PostgreSQL
* Respostas HTTP (200, 201, 400, 404, 409 e 500)

---

### 📦 Products

CRUD completo de produtos

* GET `/products` → listar todos os produtos
* GET `/products/:id` → buscar produto por ID
* POST `/products` → criar produto
* PUT `/products/:id` → atualizar produto
* PATCH `/products/:id` → atualização parcial
* DELETE `/products/:id` → remover produto
* DTOs para entrada de dados
* Tratamento centralizado de erros
* Tradução de erros do PostgreSQL
* Respostas HTTP (200, 201, 400, 404, 409 e 500)

---

## 📚 Status do projeto

🚧 Em desenvolvimento ativo

### Progresso atual

* ✅ Clients CRUD
* ✅ Products CRUD
* ✅ Tradução de erros do PostgreSQL
* ✅ Error Handler global
* 🔄 Padronização das respostas da API
* ⏳ Validações com Gin Binding
* ⏳ Regras de negócio
* ⏳ Módulo Categories

---

## 🔥 Próximos passos

* Padronizar respostas da API
* Implementar validações com Gin Binding
* Adicionar regras de negócio
* Criar módulo Categories
* Implementar autenticação JWT
* Adicionar paginação e filtros de busca

---

## 🎯 Observação

Este projeto não é apenas um CRUD simples. Ele está sendo desenvolvido com foco em arquitetura backend, entendimento do fluxo completo entre as camadas da aplicação e aplicação de boas práticas utilizadas em projetos profissionais.

O objetivo é construir um backend escalável e bem organizado, priorizando a compreensão dos conceitos por trás de cada implementação, em vez de apenas desenvolver funcionalidades.