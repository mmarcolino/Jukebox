# 🎵 Jukebox

**Jukebox** é uma API em Go para gerenciamento de faixas musicais e playlists, com execução das faixas por meio de uma fila SQS (simulada com LocalStack), representando uma abordagem orientada a eventos.

Este projeto serve como base de aprendizado para desenvolvedores iniciantes se familiarizarem com uma stack moderna e profissional, incluindo documentação automatizada via **OpenAPI**.

---

## ⚙️ Stack Utilizada

- **Go** — Linguagem principal
- **OpenAPI 3.0** — Especificação e documentação da API
- **Docker** — Ambientes isolados
- **PostgreSQL** — Banco de dados relacional
- **SQLC** — Geração de código Go a partir de queries SQL
- **DBMate** — Migrations controladas
- **LocalStack** — Simulação da AWS (SQS)
- **Arquitetura Hexagonal** — Separação entre domínio, aplicação e infraestrutura

---

## 📦 Funcionalidades

### 🎧 Faixas musicais (`/tracks`)
- Criar, listar, editar e deletar faixas musicais.
- Atributos: `title`, `artist`, `album`, `genre`, `duration`.

### 📃 Playlists (`/playlists`)
- Cria uma playlist com uma lista de faixas.
- Enfileira cada faixa na **SQS** como parte de uma simulação de "execução musical".

### 🎼 Execução com Worker
- Um processo Go separado (`cmd/worker`) consome a fila e "toca" uma faixa por vez.

---

## 📂 Estrutura de Pastas

```bash
jukebox/
├── api/
│   └── openapi.yaml                   # Especificação OpenAPI da API
├── gen/
│   ├── openapi/                       # Código gerado com oapi-codegen (tipos e handlers)
│   └── sqlc/                          # Código gerado pelo SQLC a partir das queries
├── sql/
│   ├── queries/                       # Queries SQL utilizadas pelo SQLC
│   ├── migrations/                    # Migrations controladas com DBMate
├── cmd/
│   ├── app/                           # API HTTP (main.go)
│   └── worker/                        # Worker que consome a fila e processa faixas
├── internal/
│   ├── domain/
│   │   ├── entity/                    # Entidades de domínio (Track, Playlist)
│   │   ├── gateway/                   # Interfaces (ports)
│   │   └── service/                   # Regras de negócio do domínio
│   ├── api/                           # Definições de URIs e handlers
│   ├── service/                       # Casos de uso da aplicação
│   └── resources/                     # Adapters (infraestrutura)
│       ├── database/postgres/         # Integração com banco de dados via SQLC
│       └── queue/aws/                 # Cliente SQS via AWS SDK / LocalStack
├── Dockerfile
├── docker-compose.yml
├── dbmate.toml
├── sqlc.yaml
└── README.md
``` 
# 🧪 Como Usar
## 1. Clonar o projeto
```bash
git clone https://github.com/seu-usuario/jukebox.git
cd jukebox
```
## 2. Subir os serviços (Postgres + LocalStack)
```bash
docker-compose up -d
```
## 3. Executar migrations com DBMate
``` bash
dbmate up
```
## 4. Gerar código SQLC e API via OpenAPI
``` bash
go generate ./...
```
## 5. Rodar a API
``` bash
go run cmd/app/main.go
```
## 6. Rodar o worker
``` bash
go run cmd/worker/main.go
```

# 📘 Glossário

- Worker: processo que roda em paralelo à API e consome faixas da fila para simular execução musical.

- OpenAPI: especificação que define as rotas, tipos e respostas da API em formato YAML.

- LocalStack: simula serviços da AWS localmente (como SQS).

- SQLC: gera código Go a partir de queries SQL escritas manualmente.

- DBMate: ferramenta para versionar e aplicar migrations no banco de dados.
