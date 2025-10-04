package sqlite

import (
	"api/internal/entity"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	conn, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	queryusers := `
		CREATE TABLE IF NOT EXISTS users(
			ID TEXT,
			Name TEXT,
			Number TEXT
		);
	`
	_, err = conn.Exec(queryusers)
	if err != nil {
		log.Fatal(err)
	}
	return &UserRepository{db: conn}
}

func (repo *UserRepository) Create(user *entity.User) error {

	return nil
}

func (repo *UserRepository) GetByID(id int64) (*entity.User, error) {
	return &entity.User{ID: 1, Name: "test", Number: "+323423423"}, nil
}

func (repo *UserRepository) DeleteByID(id int64) error {
	return nil
}
