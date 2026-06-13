# Customer Favorites

API REST em Go para gerenciamento de produtos favoritos por cliente.

## Sobre

Serviço que permite a clientes criar e consultar listas de produtos favoritos. Desenvolvido como exercício de API design em Go com foco em simplicidade e estrutura adequada para escalar.

## Stack

- Go 1.25+
- PostgreSQL 17+
- Docker / Docker Compose

## Estrutura

```
├── cmd/            # Entry point da aplicação
├── internal/       # Lógica de negócio e handlers
├── docs/           # Coleção Postman para testes
├── Dockerfile
└── docker-compose.yml
```

## Como executar

Para executar o projeto é necessário ter o docker instalado.

Crie o arquivo .env na raiz do projeto, a partir do arquivo .env.example.

```
cp .env.example .env
```
Altere os valores das variáveis de ambiente, conforme desejado.

Para inicializar o projeto e suas dependências utilize o comando abaixo:

```
docker compose up -d --build
go run ./cmd/...
```

Após iniciar, o app pode ser acessado através da url:

```
http://localhost:3000
```

## Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | `/customers` | Criar cliente |
| POST | `/customers/:id/favorites` | Adicionar produto aos favoritos |
| GET | `/customers/:id/favorites` | Listar favoritos do cliente |
| DELETE | `/customers/:id/favorites/:productId` | Remover produto dos favoritos |
