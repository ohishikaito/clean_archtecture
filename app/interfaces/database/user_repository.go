package database

import (
	"app/app/domain"
	"fmt"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, *user.Build())
	}
	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	fmt.Println(row)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var firstName string
	var lastName string
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	user.Build()
	return
}

func (repo *UserRepository) DeleteById(ID int) error {
	_, err := repo.Query("DELETE FROM users WHERE id = ?", ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Update(user domain.User) (domain.User, error) {
	// SQL文を生で書け！
	_, err := repo.Query("UPDATE users SET first_name = ?, last_name = ? WHERE id = ?", user.FirstName, user.LastName, user.ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Update error")
		return domain.User{}, err
	}
	user.Build()
	return user, nil
}
