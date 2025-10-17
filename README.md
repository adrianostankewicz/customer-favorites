# Customer Favorites

Este é um projeto em que o cliente pode criar, listar, atualizar e excluir os seus
produtos favoritos através de uma API Restfull.

## Tecnologias

- Postgresql 17+
- Golang 1.25+

## Como usar

Para executar o projeto é necessário ter o docker instalado.

Crie o arquivo .env na raiz do projeto, a partir do arquivo .env.example.

```
cp .env.example .env
```
Altere os valores das variáveis de ambiente, conforme desejado.

Para inicializar o projeto e suas dependências utilize o comando abaixo:

```
docker compose up -d
```