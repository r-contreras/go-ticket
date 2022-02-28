package infrastructure

import "database/sql"

type DbModel struct {
	DB *sql.DB
}

func GetDbModel(db *sql.DB) *DbModel {
	return &DbModel{
		DB: db,
	}
}
