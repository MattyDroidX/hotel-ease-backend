![Tela inicial](public/image.png)

# Hotel Ease – Back-end (Go + PostgreSQL)`

## CURSO DE TECNÓLOGO EM ANÁLISE E DESENVOLVIMENTO DE SISTEMAS 

#### EFRAIM ALVES 
#### FERNANDO MATIAS DUARTE 
#### CLERTON ALMEIDA

> Este guia ensina, **passo a passo e do zero**, como preparar o ambiente,
> compilar e executar o back-end do Hotel Ease em **Windows**, **macOS** ou
> **Ubuntu/Linux** – mesmo que você nunca tenha usado Go ou PostgreSQL.

---

## 0. Visão geral

1. Instalar **Go 1.22 LTS**  
2. Instalar **PostgreSQL 13+** e definir a senha do super‑usuário `postgres`  
3. Clonar (ou descompactar) o repositório  
4. Rodar `go mod tidy` → baixa dependências  
5. Executar `backend/run.sh`  
   * cria banco/usuário;  
   * sobe a API em `http://localhost:8080`;  
   * abre o **Swagger** no navegador.  
6. (Opcional) Inspecionar tabelas e registros pelo **psql** ou GUI.  
7. Em outro terminal, executar `make frontend` → abre a interface React.  

---

## 1. Instalando Go 1.22

| Sistema   | Passos rápidos | Teste |
|-----------|----------------|-------|
| **Windows** | 1. Baixe o instalador em <https://go.dev/dl/> (x86_64).<br>2. `Next → Next → Finish`. | **Prompt** → `go version` |
| **macOS** | 1. Baixe o **pkg** no mesmo link.<br>2. Clique duas vezes e conclua. | `go version` no Terminal |
| **Ubuntu** | 
```bash
curl -LO https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
``` 

| `go version` | Você deve ver algo como **go version go1.22.5**.

---

## 2. Instalando PostgreSQL

### Windows / macOS

1. Acesse <https://www.postgresql.org/download>  
2. Baixe o instalador da versão **13 ou superior**  
3. Informe uma **senha para `postgres`** (anote).  
4. Mantenha porta **5432** → Next → Finish.

### Ubuntu / Debian

```bash
sudo apt update && sudo apt install postgresql -y
sudo systemctl enable --now postgresql
```

Se quiser definir senha:

```bash
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD 'minhaSenha';"
```

---

## 3. Baixando o projeto

```bash
git clone https://github.com/MattyDroidX/hotel-ease-backend.git
cd hotel-ease-backend
```

### *(ou descompacte o arquivo zip/rar fornecido)*

#### O mesmo tera a estrutura de diretorios do apartado 10 final.

```bash
hotel-ease/
├─ hotel-ease-backend/
│  ├─ .env
│  ├─ run.sh
│  └─ …
├─ hotel-ease-frontend/
│  └─ …
└─ Makefile
```

---

## 4. Instalando dependências Go

```bash
cd hotel-ease-backend
go mod tidy        # baixa gin, sqlx, swag, etc.
```

---

## 5. Configurando e executando o back‑end

### 5.1 `.env` (opcional)

`hotel-ease-backend/.env`:

```dotenv
DB_URL=postgres://hotelease_user:admin@localhost:5432/hotelease?sslmode=disable
PORT=8080

# Se seu Postgres exige senha p/ super‑usuário:
SUPERUSER=postgres
SUPERPASS=minhaSenhaDoPasso2
```

### 5.2 Rodando

```bash
cd hotel-ease-backend
chmod +x ./run.sh
./run.sh        # Linux/macOS ou Git Bash (Windows)

```

Saída típica:

```
✅ Banco e usuário criados com sucesso!
▶ API em http://localhost:8080  (Ctrl+C para sair)
```

### 5.2 Rodando

Para ver a documentacao da API em Swagger

```bash
   http://localhost:8080/swagger/index.html
```

---

## 6. Explorando o banco

### 6.1 CLI `psql`

```bash
make db             # usa DB_URL
```

Comandos:

```sql
\dt                 -- lista tabelas
SELECT * FROM funcionarios;
SELECT * FROM tarefas;
\q                 -- sair
```

### 6.2 GUI

* **pgAdmin 4** (Windows)  
* **DBeaver**, **TablePlus** …

Host `localhost` | Porta `5432` | DB `hotelease` | User `hotelease_user` | Senha `admin`

---

## 7. Executando o front‑end

```bash
make frontend      # abre React em http://localhost:3000 ou http://127.0.0.1:3000
```

---

## 8. Parando

* **Back‑end**: `Ctrl + C` (ou `make stop`)  
* **Front‑end**: `Ctrl + C`  
* Postgres continua rodando.

---

## 9. Problemas comuns

| Sintoma | Causa & solução |
|---------|-----------------|
| `password authentication failed for user "postgres"` | Senha errada ou `SUPERPASS` faltando no `.env`. |
| `port 8080 already in use` | Outra aplicação usa 8080 → troque `PORT` no `.env`. |
| `psql: could not connect to server` | Postgres inativo. Windows → *pgAdmin* › Start Server. Linux → `sudo systemctl start postgresql`. |
| Front‑end “Cannot GET /funcionarios” | API ainda subindo – aguarde 2 s ou recarregue. |

---

## 10. Estrutura de diretórios

```
hotel-ease/
├─ hotel-ease-backend/
│  ├─ .env
│  ├─ run.sh
│  └─ …
├─ hotel-ease-frontend/
│  └─ …
└─ Makefile
```

---
