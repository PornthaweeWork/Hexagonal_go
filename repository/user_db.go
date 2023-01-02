package repository

import "github.com/jmoiron/sqlx"

type userRepositoryDB struct {
	db *sqlx.DB
}

// constructor
func NewUserRepositoryDB(db *sqlx.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetAll() ([]User, error) {
	users := []User{}
	query := "select * from usercomport"
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepositoryDB) GetById(id int) (*User, error) {
	user := User{}
	query := "select * from usercomport where user_id=$1"
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
