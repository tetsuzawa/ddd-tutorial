package ddd_tutorial

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	Save(user User) (err error)
	Find(name UserName) (u *User, err error)
}

type UserRepository struct {
	connectionString string
}

func NewUserRepository(connectionString string) *UserRepository {
	return &UserRepository{connectionString: connectionString}
}

func (r UserRepository) Save(user User) (err error) {
	conn, err := sqlx.Open("mysql", r.connectionString)
	if err != nil {
		return err
	}
	stmt, err := conn.Prepare(`
INSERT INTO user (id, name)
VALUES (?, ?)
ON DUPLICATE KEY
UPDATE name = ?
`)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	_, err = stmt.Exec(user.userId, user.name)
	return err
}

func (r UserRepository) Find(userName UserName) (u *User, err error) {
	conn, err := sqlx.Open("mysql", r.connectionString)
	if err != nil {
		return nil, err
	}
	if err = conn.Get(u, "SELECT * FROM user WHERE name = ?"); err != nil {
		return nil, err
	}
	return u, nil
}
