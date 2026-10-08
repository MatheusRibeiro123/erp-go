# ERP em Go

Backend de um sistema ERP desenvolvido em Go, com foco no aprendizado e aplicação de conceitos de arquitetura backend, APIs REST, PostgreSQL, Docker e boas práticas utilizadas no desenvolvimento de software.

O projeto simula a construção de um sistema ERP real, utilizando arquitetura em camadas, separação de responsabilidades, injeção de dependências, DTOs, validação de entradas e tratamento centralizado de erros.

> **Status:** 🚀 v1.0.0 — primeira versão estável do projeto.

---

## 🚀 Tecnologias

* **Go 1.26+**
* **Gin**
* **PostgreSQL**
* **pgx**
* **Docker**
* **Docker Compose**
* **Git / GitHub**

---

## 🏗️ Arquitetura

O projeto utiliza uma arquitetura em camadas, separando as responsabilidades da aplicação:

```text
HTTP Request
     ↓
  Handler
     ↓
  Service
     ↓
 Repository
     ↓
PostgreSQL
```

### Responsabilidades

**Handler**

* Recebe as requisições HTTP.
* Realiza o binding e validação dos dados de entrada.
* Chama a camada de Service.
* Retorna as respostas HTTP.

**Service**

* Contém as regras de negócio da aplicação.
* Faz a comunicação entre Handler e Repository.
* Mantém a lógica de negócio separada da camada HTTP.

**Repository**

* Responsável pelo acesso ao banco de dados.
* Executa consultas, inserções, atualizações e exclusões no PostgreSQL.

**DTO**

* Define os dados esperados nas requisições.
* Auxilia na validação das entradas da API.
* Evita utilizar diretamente os modelos do banco como entrada das requisições.

**Models**

* Representam as entidades utilizadas pela aplicação.

**AppErrors**

* Centraliza erros da aplicação.
* Traduz erros específicos do PostgreSQL.
* Padroniza o tratamento dos erros HTTP.

**Responses**

* Centraliza o formato das respostas de sucesso da API.

---

## 📁 Estrutura do projeto

```text
erp-go/
│
├── internal/
│   ├── apperrors/
│   │   ├── errors.go
│   │   ├── handler.go
│   │   └── postgres.go
│   │
│   ├── database/
│   │   └── connection.go
│   │
│   ├── dto/
│   │   ├── client.go
│   │   └── products.go
│   │
│   ├── handlers/
│   │   ├── client_handler.go
│   │   ├── home.go
│   │   └── product_handler.go
│   │
│   ├── models/
│   │   ├── clients.go
│   │   └── product.go
│   │
│   ├── repositories/
│   │   ├── client_repository.go
│   │   └── product_repository.go
│   │
│   ├── responses/
│   │   └── response.go
│   │
│   ├── routes/
│   │   ├── client_routes.go
│   │   └── product_routes.go
│   │
│   └── services/
│       ├── client_service.go
│       └── product_service.go
│
├── .dockerignore
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── init.sql
├── main.go
├── README.md
└── TODO.md
```

---

## 📌 Funcionalidades

### 👤 Clients

CRUD completo de clientes.

| Método   | Endpoint       | Descrição                        |
| -------- | -------------- | -------------------------------- |
| `GET`    | `/clients`     | Lista clientes                   |
| `GET`    | `/clients/:id` | Busca cliente por ID             |
| `POST`   | `/clients`     | Cria um cliente                  |
| `PUT`    | `/clients/:id` | Atualiza um cliente              |
| `PATCH`  | `/clients/:id` | Atualiza parcialmente um cliente |
| `DELETE` | `/clients/:id` | Remove um cliente                |

### Campos

```text
id
name
email
phone
document
created_at
```

---

### 📦 Products

CRUD completo de produtos.

| Método   | Endpoint        | Descrição                        |
| -------- | --------------- | -------------------------------- |
| `GET`    | `/products`     | Lista produtos                   |
| `GET`    | `/products/:id` | Busca produto por ID             |
| `POST`   | `/products`     | Cria um produto                  |
| `PUT`    | `/products/:id` | Atualiza um produto              |
| `PATCH`  | `/products/:id` | Atualiza parcialmente um produto |
| `DELETE` | `/products/:id` | Remove um produto                |

### Campos

```text
id
name
description
price
stock_quantity
created_at
```

---

## 📄 Paginação

Os endpoints de listagem possuem paginação através dos parâmetros:

```text
?page=1&limit=10
```

Exemplo:

```http
GET /clients?page=1&limit=10
```

Resposta:

```json
{
  "data": [],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 0
  }
}
```

A paginação também possui validação dos parâmetros enviados.

> Filtros de busca não fazem parte da versão 1.0 e estão planejados para uma versão futura.

---

## ✅ Validação de entradas

As entradas da API são validadas utilizando DTOs e o sistema de binding/validation do Gin.

Exemplos de validações aplicadas:

* Campos obrigatórios.
* Tamanho mínimo de strings.
* Valores numéricos maiores ou iguais a zero.
* Validação específica para atualizações parciais utilizando ponteiros.

---

## ⚠️ Tratamento de erros

O projeto possui tratamento centralizado de erros da aplicação.

Os erros do PostgreSQL são traduzidos para erros compreensíveis pela aplicação, permitindo retornar respostas HTTP adequadas.

Entre os cenários tratados estão:

* Recurso não encontrado.
* Dados duplicados.
* Violação de chave estrangeira.
* Dados inválidos.
* Erros internos do servidor.

### Principais status utilizados

| Status | Significado                                   |
| ------ | --------------------------------------------- |
| `200`  | Operação realizada com sucesso                |
| `201`  | Recurso criado                                |
| `204`  | Operação realizada sem conteúdo para retornar |
| `400`  | Dados de entrada inválidos                    |
| `404`  | Recurso não encontrado                        |
| `409`  | Conflito / dado duplicado                     |
| `500`  | Erro interno do servidor                      |

---

## 📤 Padronização das respostas

As respostas de sucesso da API são centralizadas no pacote `responses`.

Exemplo:

```json
{
  "message": "Client created successfully",
  "data": {
    "id": 1
  }
}
```

As listas utilizam `[]` quando não existem registros, evitando retornar `null`.

---

# 🐳 Docker

O projeto possui suporte completo a Docker.

A aplicação pode ser executada junto com o PostgreSQL utilizando Docker Compose.

A composição possui dois serviços:

```text
┌───────────────┐
│   erp-api     │
│    Go/Gin     │
│    :8080      │
└───────┬───────┘
        │
        │ Docker Network
        ↓
┌───────────────┐
│ erp-postgres  │
│  PostgreSQL   │
│    :5432      │
└───────────────┘
```

O PostgreSQL utiliza um volume Docker para manter os dados mesmo após o reinício dos containers.

---

## ⚙️ Como executar o projeto

### Pré-requisitos

Para executar localmente com Docker, é necessário ter:

* Docker
* Docker Compose

Não é necessário instalar o PostgreSQL separadamente para executar a aplicação através do Docker Compose.

---

### 1. Clone o repositório

```bash
git clone https://github.com/MatheusRibeiro123/erp-go.git
```

Entre na pasta:

```bash
cd erp-go
```

---

### 2. Suba os containers

Execute:

```bash
docker compose up --build
```

O Docker irá:

1. Criar o container do PostgreSQL.
2. Inicializar o banco de dados.
3. Executar o `init.sql` em um banco novo.
4. Criar as tabelas `clients` e `products`.
5. Aguardar o PostgreSQL ficar saudável.
6. Construir a imagem da API.
7. Iniciar a API.

A API ficará disponível em:

```text
http://localhost:8080
```

---

## 🗄️ Banco de dados

O banco utilizado pela aplicação é:

```text
PostgreSQL
```

As tabelas iniciais são criadas pelo arquivo:

```text
init.sql
```

O Docker Compose utiliza um volume chamado:

```text
postgres_data
```

Esse volume permite que os dados permaneçam armazenados quando os containers são parados.

### Parar os containers

```bash
docker compose down
```

Os dados do banco permanecem no volume.

### Remover containers e dados

```bash
docker compose down -v
```

> O comando acima remove também o volume do PostgreSQL. Ao executar `docker compose up` novamente, o banco será inicializado do zero e o `init.sql` será executado novamente.

---

## 🔐 Variáveis de ambiente

A aplicação utiliza variáveis de ambiente para configurar a conexão com o PostgreSQL.

Exemplo:

```text
DB_HOST
DB_PORT
DB_USER
DB_PASSWORD
DB_NAME
DB_SSL_MODE
```

No Docker Compose, essas variáveis são configuradas diretamente no serviço da API.

O arquivo `.env` não é incluído na imagem Docker.

---

# 🧪 Testes da API

Durante o desenvolvimento da versão 1.0, os endpoints foram testados manualmente utilizando um cliente HTTP, validando os principais fluxos de CRUD, atualização parcial, validações, erros e paginação.

Os testes incluem cenários como:

* Criação de clientes.
* Consulta de clientes.
* Atualização completa.
* Atualização parcial.
* Exclusão.
* Criação de produtos.
* Consulta de produtos.
* Atualização completa.
* Atualização parcial.
* Exclusão.
* Recursos inexistentes.
* Dados inválidos.
* Conflitos de dados.
* Paginação.

> Testes automatizados com `go test` fazem parte de uma evolução futura do projeto.

---

# 📚 Conceitos aplicados

Durante o desenvolvimento foram aplicados conceitos como:

* Arquitetura em camadas.
* Separação de responsabilidades.
* Injeção de dependências.
* APIs REST.
* DTOs.
* Validação de entradas.
* Tratamento centralizado de erros.
* `errors.Is`.
* Tradução de erros do PostgreSQL.
* Status HTTP apropriados.
* Atualização parcial com `PATCH`.
* Ponteiros em DTOs para diferenciar campos não enviados.
* `RowsAffected()` para operações de atualização e exclusão.
* Paginação.
* Variáveis de ambiente.
* Docker.
* Docker Compose.
* Healthcheck de serviços.
* Persistência através de volumes Docker.

---

# 🎯 Objetivos do projeto

O principal objetivo do ERP em Go é consolidar conhecimentos de desenvolvimento backend através da construção de uma aplicação prática.

O projeto busca desenvolver experiência em:

* Desenvolvimento backend com Go.
* Construção de APIs REST.
* Arquitetura de aplicações.
* PostgreSQL.
* Docker.
* Organização de código.
* Tratamento de erros.
* Validação de dados.
* Separação de responsabilidades.
* Boas práticas de desenvolvimento.

O foco não é apenas fazer a API funcionar, mas compreender o fluxo completo de uma aplicação backend e as responsabilidades de cada camada.

---

# 🔮 Próximos passos

A versão 1.0 representa a primeira versão estável do projeto.

Possíveis evoluções futuras incluem:

* Filtros de busca.
* Testes automatizados.
* Autenticação e autorização com JWT.
* CI/CD.
* Deploy em ambiente cloud.
* Novos módulos do ERP.
* Melhorias de observabilidade.
* Evolução da arquitetura.

Essas funcionalidades não fazem parte do escopo da versão 1.0.

---

# 📌 Status

## v1.0.0

* ✅ Clients CRUD
* ✅ Products CRUD
* ✅ PUT
* ✅ PATCH
* ✅ Validação de entradas
* ✅ Tratamento de erros
* ✅ Respostas padronizadas
* ✅ Paginação
* ✅ PostgreSQL
* ✅ Docker
* ✅ Docker Compose
* ✅ Persistência do banco através de volume
* 🚧 Filtros de busca — planejado para versão futura
* 🚧 Testes automatizados — planejado para versão futura

---

## 👨‍💻 Sobre o projeto

Este projeto foi desenvolvido como parte do processo de aprendizado e construção de experiência prática em desenvolvimento backend com Go.

A prioridade durante o desenvolvimento foi entender os conceitos e decisões por trás da implementação, evitando apenas reproduzir código sem compreender seu funcionamento.

O projeto representa uma aplicação prática dos conhecimentos de arquitetura backend, APIs REST, banco de dados, tratamento de erros, validação e containerização.

---

**ERP em Go — v1.0.0**
