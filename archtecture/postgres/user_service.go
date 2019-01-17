package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/takochuu/sandbox/archtecture"
)

type UserService struct {
	DB *sql.DB
}

func (s *UserService) User(id int) (*archtecture.User, error) {
	var u archtecture.User
	row := s.DB.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id)
	if row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}
