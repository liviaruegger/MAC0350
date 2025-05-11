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
├── .gitignore
├── backend/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── app/
│   │   │   └── user_service.go
│   │   ├── config/
│   │   │   └── database.go
│   │   ├── domain/
│   │   │   ├── activity.go
│   │   │   └── user.go
│   │   ├── handler/
│   │   │   └── user_handler.go
│   │   └── repository/
│   │       ├── activity_repository.go
│   │       └── user_repository.go
│   └── utils/
│       └── curls.txt
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

A aplicação segue uma separação de responsabilidades em camadas, baseada nos princípios de clean architecture e nas convenções da linguagem Go, como descrito abaixo:

### `internal/app/`
#### Camada de regras de negócio (serviços)
Esta camada contém a lógica de negócio da aplicação: toma decisões, valida dados, define o fluxo de operações entre as entidades e os repositórios.

### `internal/config/`
#### Configuração da aplicação
Contém a configuração de serviços e dependências externas, como a conexão com o banco de dados.

### `internal/domain/`
#### Camada de domínio
Definição das entidades que compõem o sistema (como User, Activity, Interval) e tipos auxiliares; essas estruturas representam os objetos "reais" com os quais a aplicação lida.

### `internal/handler/`
#### Camada de manipulação de requisições HTTP
Responsável por receber as requisições HTTP, extrair os dados necessários e repassar essas informações para a camada de aplicação (serviços). Também é responsável por desenvolver uma resposta HTTP adequada.

### `internal/repository/`
#### Camada de persistência de dados
Camada de abstração de acesso ao banco de dados: realiza operações de CRUD (Create, Read, Update, Delete) e cria interfaces para serem utilizadas pelos serviços da aplicação.

## Como executar
### Pré-requisitos
- Linux:
    - [Docker Engine](https://docs.docker.com/engine/install/)
    - [Docker Compose](https://docs.docker.com/compose/install/)

### Rodando
O container pode ser construído e rodado com
```
docker-compose up --build
```

### Comandos úteis
#### Apagar o container e seus volumes
```
docker-compose down -v
```
#### Construir o container do zero
```
docker-compose up build --no-cache
```

## Como testar
Para rodar todos os testes do backend, use o comando
```
go test ./backend/...
```
Para analisar a cobertura de testes, utilize a flag `-cover`:
```
go test ./backend/... -cover
```

---
Este projeto está sendo desenvolvido para a disciplina MAC0350 - Introdução ao Desenvolvimento de Sistemas de Software (2025.1) do IME-USP.
