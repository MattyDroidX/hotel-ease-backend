HotelEase Backend 🏨
=====================

Backend da aplicação HotelEase, feito em Go com PostgreSQL.

---------------------
REQUISITOS
---------------------

- Go instalado: https://go.dev/dl/
- PostgreSQL
- Um arquivo `.env` com a seguinte variável:
  
  DB_URL=postgres://hotelease_user:admin@localhost:5432/hotelease?sslmode=disable
  PORT=8080

---------------------
COMO RODAR O BACKEND
---------------------

1. Clone o repositório:

   git clone https://github.com/seuusuario/hotel-ease-backend.git
   cd hotel-ease-backend

2. Copie o `.env.example` como `.env`:

   cp .env.example .env

3. (Opcional) Instale as dependências:

   go mod tidy

4. Rode o backend com:

   go run main.go

   Ou use:

   ./run.sh

---------------------
BANCO DE DADOS
---------------------

A aplicação cria as tabelas automaticamente ao iniciar.

Certifique-se de que o PostgreSQL esteja rodando localmente com um banco chamado `hotelease`.

Você pode criar manualmente com:

   CREATE DATABASE hotelease;

---------------------
ENDPOINTS DISPONÍVEIS
---------------------

- GET    /funcionarios
- GET    /funcionarios/:id
- POST   /funcionarios
- PUT    /funcionarios/:id
- DELETE /funcionarios/:id

- GET    /tarefas
- GET    /tarefas/:id
- POST   /tarefas
- PUT    /tarefas/:id
- DELETE /tarefas/:id

---------------------
DOCS COM SWAGGER
---------------------

Acesse a documentação completa da API após rodar o projeto em:

  http://localhost:8080/swagger/index.html

---------------------
AUTOR
---------------------

Feito com 💙 por [Seu Nome]
