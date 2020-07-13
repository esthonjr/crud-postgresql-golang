# crud-postgresql-golang

Este pacote é um exemplo simples de CRUD (Create, read, update and delete) com golang + postgresql.

- [crud-postgresql-golang](#crud-postgresql-golang)
  - [Configuração da conexão ao PostgreSQL](#configuração-da-conexão-ao-postgresql)
  - [Utilização](#utilização)

## Configuração da conexão ao PostgreSQL
Editar o arquivo crud.go:
```bash
PGUSER=esthonjr
PGHOST=localhost
PGDATABASE=testdb
PGPASSWORD=estadmin
PGPORT=5432
```

## Utilização

Gerar executável após editar o crud.go com a configuração da conexão ao PostgreSQL:

```bash
$ go build
```

Comandos disponíveis:

```bash
$ ./crud-psql-golang create # para criar a tabela
$ ./crud-psql-golang read # para ler tabela
$ ./crud-psql-golang update # para fazer update dos dados
$ ./crud-psql-golang insert # para inserir dados
$ ./crud-psql-golang delete # para deletar os dados
$ ./crud-psql-golang drop # para deletar a tabela
```
