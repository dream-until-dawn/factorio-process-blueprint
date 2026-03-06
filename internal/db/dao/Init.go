package dao

import (
	"database/sql"

	"processBlueprint/internal/db"
)

type DBDAO struct {
	db *sql.DB
}

func NewDBDAO() *DBDAO {
	dbDao := &DBDAO{
		db: db.Get().DB,
	}
	return dbDao
}
