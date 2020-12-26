package ddd_tutorial

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	Save(user User) (retsult sql.Result, err error)
	Find(name UserName) (*User, error)
}

type UserRepository struct {
	connectionString string
}

func NewUserRepository(connectionString string) *UserRepository {
	return &UserRepository{connectionString: connectionString}
}

func (r UserRepository) Save(user User) (retsult sql.Result, err error) {
	conn, err := sqlx.Open("mysql", r.connectionString)
	if err != nil {
		return nil, err
	}
	stmt, err := conn.Prepare(`
INSERT INTO user (id, name)
VALUES (?, ?)
ON DUPLICATE KEY
UPDATE name = ?
`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return stmt.Exec(user.userId, user.name)
}

func (r UserRepository) Find(userName UserName) (user *User, err error) {
	conn, err := sqlx.Open("mysql", r.connectionString)
	if err != nil {
		return nil, err
	}
	if err = conn.Get(user, "SELECT * FROM user WHERE name = ?"); err != nil {
		return nil, err
	}
	return user, nil
}
