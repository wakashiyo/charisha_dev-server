package database

import (
	"fmt"

	"github.com/wakashiyo/charisha_dev-server/domain"
)

//UserRepository repository for user
type UserRepository struct {
	DbHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Excute(fmt.Sprintf("INSERT INTO users (name, mail) VALUES ('%s', '%s')", u.Name, u.Mail))
	if err != nil {
		return
	}
	_id, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(_id)
	return
}

func (repo *UserRepository) FindById(identifier int) (u domain.User, err error) {
	row, err := repo.Query(fmt.Sprintf("SELECT id, name, mail FROM users WHERE id = '%d'", identifier))
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var name string
	var mail string
	row.Next()
	if err = row.Scan(&id, &name, &mail); err != nil {
		return
	}
	u.ID = id
	u.Name = name
	u.Mail = mail
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, name, mail FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		var mail string
		if err := rows.Scan(&id, &name, &mail); err != nil {
			continue
		}
		user := domain.User{
			ID:   id,
			Name: name,
			Mail: mail,
		}
		users = append(users, user)
	}
	return
}
