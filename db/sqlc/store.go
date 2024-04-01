package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface{
}

type SQLStore struct{
	connPool *pgxpool.Pool
	Queries *Queries
}

func NewStore(connPool *pgxpool.Pool) SQLStore{
	return SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
