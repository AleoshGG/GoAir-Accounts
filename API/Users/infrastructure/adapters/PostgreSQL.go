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
	query := "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id_user"

	var id uint
	err := postgres.conn.DB.QueryRow(query, u.First_name, u.Last_name, u.Email, u.Password).Scan(&id)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	return id, nil
}

func (postgres *PostgreSQL) DeleteUser(id_user int) (uint, error) {
	query := "DELETE FROM users WHERE id_user = $1"

	fmt.Print(id_user)
	_, err := postgres.conn.DB.Exec(query, id_user)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	return uint(1), nil
}

func (postgres *PostgreSQL) GetUserByEmail(email string) domain.User {
	query := "SELECT * FROM users WHERE email = $1"
	var user domain.User
	
	rows, err := postgres.conn.FetchRows(query, email)

	if err != nil {
		fmt.Errorf("error al ejecutar la consulta: %w", err)
		return domain.User{}
	}

	defer rows.Close()

	if !rows.Next() {
        fmt.Println("No se pudieron obtener los datos.")
        return domain.User{}
    }

	if err := rows.Scan(&user.Id_user, &user.First_name, &user.Last_name, &user.Email, &user.Password); err != nil {
		fmt.Errorf("error al escanear el usuario: %w", err)
        return domain.User{}
    }

	
	return user
}

func (postgres *PostgreSQL) GetUserById(id_user int) domain.User {
	query := "SELECT * FROM users WHERE id_user = $1"
	var user domain.User
	
	rows, err := postgres.conn.FetchRows(query, id_user)

	if err != nil {
		fmt.Errorf("error al ejecutar la consulta: %w", err)
		return domain.User{}
	}

	defer rows.Close()

	if !rows.Next() {
        fmt.Println("No se pudieron obtener los datos.")
        return domain.User{}
    }

	if err := rows.Scan(&user.Id_user, &user.First_name, &user.Last_name, &user.Email, &user.Password); err != nil {
		fmt.Errorf("error al escanear el usuario: %w", err)
        return domain.User{}
    }
	
	return user
}