# Swim Tracker (nome provisório)

#### Autores
- Ana Lívia Rüegger Saldanha (NUSP: 8686691)
- Gustavo Mota Bastos (NUSP: 10284389)

## Descrição geral
A proposta é desenvolver um sistema para registro de treinos de natação, com feed de atividades e possibilidade de compartilhamento (como o Strava, mas permitindo a inserção manual das informações, sem necessidade de um dispositivo externo).

## Escopo
As funcionalidades incluídas, a princípio, serão:
- Cadastro do usuário com informações pessoais básicas;
- Registro de atividades: cada atividade registrada poderá incluir título, data, local, distâncias (possibilitando discriminar por estilos), horário, duração total, esforço percebido e comentários;
- Perfil com feed para compartilhamento de atividades;
- Geração de resumo com estatísticas por período (semanal, mensal).

## Organização do projeto
```
MAC0350/
├── .env
├── .gitignore
├── backend/
│   ├── cmd/
│   │   └── main.go
│   ├── config/
│   │   └── database.go
│   ├── internal/
│   │   ├── app/
│   │   │   ├── interval_service_test.go
│   │   │   ├── interval_service.go
│   │   │   ├── user_service_test.go
│   │   │   └── user_service.go
│   │   ├── domain/
│   │   │   ├── activity_test.go
│   │   │   ├── activity.go
│   │   │   ├── duration_test.go
│   │   │   ├── duration.go
│   │   │   ├── interval_test.go
│   │   │   ├── interval.go
│   │   │   └── user.go
│   │   ├── handler/
│   │   │   ├── interval_handler_test.go
│   │   │   ├── interval_handler.go
│   │   │   ├── response.go
│   │   │   ├── user_handler_test.go
│   │   │   └── user_handler.go
│   │   └── repository/
│   │       ├── activity_repository_test.go
│   │       ├── activity_repository.go
│   │       ├── interval_repository_test.go
│   │       ├── interval_repository.go
│   │       ├── user_repository_test.go
│   │       └── user_repository.go
│   └── utils/
│       ├── curls.txt
│       └── wait-for-it.sh
├── docker-compose.yml
├── Dockerfile
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
```
A árvore acima foi gerada utilizando [esta ferramenta](https://project-tree-generator.netlify.app/generate-tree).

A aplicação segue uma separação de responsabilidades em camadas, baseada nos princípios de clean architecture e nas convenções da linguagem Go, como descrito abaixo:

### `config/` — Configuração da aplicação
Contém a configuração de serviços e dependências externas, como a conexão com o banco de dados.

### `internal/app/` — Camada de regras de negócio (serviços)
Esta camada contém a lógica de negócio da aplicação: toma decisões, valida dados, define o fluxo de operações entre as entidades e os repositórios.

### `internal/domain/` — Camada de domínio
Definição das entidades que compõem o sistema (como User, Activity, Interval) e tipos auxiliares; essas estruturas representam os objetos "reais" com os quais a aplicação lida.

### `internal/handler/` — Camada de manipulação de requisições HTTP
Responsável por receber as requisições HTTP, extrair os dados necessários e repassar essas informações para a camada de aplicação (serviços). Também é responsável por desenvolver uma resposta HTTP adequada.

### `internal/repository/` — Camada de persistência de dados
Camada de abstração de acesso ao banco de dados: realiza operações de CRUD (Create, Read, Update, Delete) e cria interfaces para serem utilizadas pelos serviços da aplicação.

## Como executar
### Pré-requisitos
#### Docker
- Linux:
    - [Docker Engine](https://docs.docker.com/engine/install/)
    - [Docker Compose](https://docs.docker.com/compose/install/)

#### Swag
Para testar a API utilizando a UI do Swagger, instale utilizando o comando: 
``` bash
go install github.com/swaggo/swag/cmd/swag@latest
```
Se necessário, exporte o path:
``` bash
export PATH=$PATH:$HOME/go/bin
```

### Rodando
Para construir o container e rodar o projeto:
```
make run
```
Para apagar o container e seus volumes:
```
make docker-down
```

## Como testar
### Backend
Para rodar todos os testes do backend:
```
make test
```
Para analisar a cobertura de testes:
```
make coverage
```
Para gerar um relatório detalhando a cobertura de testes em cada arquivo:
```
make test-report
```

### API
Para testar a API, execute o projeto com `make run` e acesse a [UI do Swagger](http://localhost:8080/swagger/index.html).

## Outros comandos úteis
Para atualizar as dependências Go (`go.mod` e `go.sum`):
```
go mod tidy
```

---
Este projeto está sendo desenvolvido para a disciplina MAC0350 - Introdução ao Desenvolvimento de Sistemas de Software (2025.1) do IME-USP.
