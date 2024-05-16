## 🚀 Como executar

## Pre-requisitos

Antes de iniciar, vamos precisar de algumas ferramentas instaladas na sua máquina como Git e Docker. É bom ter uma IDE para trabalhar com o código como o VSCode

## 🧭 Executando o projeto

# Clonando o repositorio

$ git clone git@github.com:jckonewalik/migrations-golang-yt.git

## Acesse a pasta via terminal

```bash
 cd migrations-golang-yt && cp .env.sample .env
```

## Variáveis de ambiente

Preencha as variaveis de ambiente do arquivo .env

## Iniciar o banco de dados

```bash
docker-compose up -b
```

## Executando as Migrations

```bash
make migrate-up
```
