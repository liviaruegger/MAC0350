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
  
---
Este projeto está sendo desenvolvido para a disciplina MAC0350 - Introdução ao Desenvolvimento de Sistemas de Software (2025.1) do IME-USP.
