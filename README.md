# ERP em Go

Sistema ERP desenvolvido em Go com foco em aprendizado de arquitetura backend profissional, APIs REST, PostgreSQL e boas práticas utilizadas no mercado.

## 🚀 Tecnologias

* Go (Golang)
* PostgreSQL
* Gin
* Git/GitHub

## 🏗️ Arquitetura

O projeto utiliza arquitetura em camadas:

```text
Handler
↓
Service
↓
Repository
↓
PostgreSQL
```

Responsabilidades:

* Handler → Recebe e responde requisições HTTP
* Service → Regras de negócio
* Repository → Acesso ao banco de dados
* PostgreSQL → Persistência dos dados

## 📁 Estrutura do projeto

```bash
erp-go/
│
├── internal/
│   ├── handlers/
│   ├── services/
│   ├── repositories/
│   ├── models/
│   ├── routes/
│   └── database/
│
├── main.go
├── go.mod
└── go.sum
```

## 📌 Objetivos do projeto

Este projeto está sendo desenvolvido para estudar:

* Desenvolvimento backend com Go
* APIs REST
* Arquitetura em camadas
* Injeção de dependências
* PostgreSQL
* Organização profissional de projetos
* Tratamento de erros
* Autenticação e autorização
* Boas práticas de mercado

## ✅ Funcionalidades implementadas

### Clientes

* Buscar todos os clientes (`GET /clients`)
* Buscar cliente por ID (`GET /clients/:id`)

## ⚙️ Funcionalidades planejadas

### Clientes

* Criar cliente
* Atualizar cliente
* Remover cliente

### Produtos

* Cadastro de produtos
* Controle de estoque

### Vendas

* Registro de vendas
* Movimentação de estoque

### Financeiro

* Controle de contas
* Fluxo de caixa

### Segurança

* Login
* JWT
* Controle de permissões

## 📚 Status do projeto

🚧 Em desenvolvimento

Projeto criado para fins de estudo, portfólio e preparação para oportunidades na área de desenvolvimento backend com Go.
