package repository

import (
    // "database/sql"
    "github.com/liviaruegger/MAC0350/backend/internal/domain"
)

var Profiles = []domain.User{
	{ID: 1, Name: "João da Silva", Email: "joao@example.com", City: "São Paulo", Phone: "(00) 0 0000-0000", Activities: Activities_1},
	{ID: 2, Name: "Maria Souza", Email: "maria@example.com", City: "Não-Me-Toque", Phone: "(00) 0 0000-0000", Activities: Activities_2},
	{ID: 3, Name: "Ana Costa", Email: "ana@example.com", City: "Vitória", Phone: "(00) 0 0000-0000", Activities: Activities_3},
}

var Activities_1 = []domain.Activity{}

var Activities_2 = []domain.Activity{}

var Activities_3 = []domain.Activity{}