package adapters

import (
	"GoAir-Accounts/API/Users/domain"
	"GoAir-Accounts/API/core"
	"fmt"
)

type PostgreSQL struct {
	conn *core.ConnPostgreSQL
}

func NewPostgreSQL() *PostgreSQL {
	conn := core.GetDBPool()
	
	if conn.Err != "" {
		fmt.Println("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &PostgreSQL{conn: conn}
}

func (postgres *PostgreSQL) CreateUser(u domain.User)(uint, error) {
	query := "INSERT INTO users (first_name, last_name, email, password) VALUES (?,?,?,?)"

	res, err := postgres.conn.ExecutePreparedQuery(query, u.First_name, u.Last_name, u.Email, u.Password)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	id, _ := res.LastInsertId() 

	return uint(id), nil
}

func (postgres *PostgreSQL) DeleteUser(id_user int) (uint, error) {
	query := "DELETE FROM users WHERE id_user = ?"

	res, err := postgres.conn.ExecutePreparedQuery(query, id_user)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	rows, _ := res.RowsAffected() 

	return uint(rows), nil
}

func (postgres *PostgreSQL) GetUserByEmail(email string) domain.User {
	query := "SELECT * FROM users WHERE email = ?"
	var users []domain.User
	
	rows, _ := postgres.conn.FetchRows(query, email)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return domain.User{}
    }

	defer rows.Close()

	for rows.Next() {
		var u domain.User
		rows.Scan(&u.First_name, &u.Last_name, &u.Email, &u.Password)

		users = append(users, u)
	}
	
	return users[0]
}

func (postgres *PostgreSQL) GetUserById(id_user int) domain.User {
	query := "SELECT * FROM users WHERE id_user = ?"
	var users []domain.User
	
	rows, _ := postgres.conn.FetchRows(query, id_user)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return domain.User{}
    }

	defer rows.Close()

	for rows.Next() {
		var u domain.User
		rows.Scan(&u.First_name, &u.Last_name, &u.Email, &u.Password)

		users = append(users, u)
	}
	
	return users[0]
}