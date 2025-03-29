package adapters

import (
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

func (postgre *PostgreSQL) CreateApplication(id_user int) (uint, error) {
	query := "INSERT INTO applications (status_application, id_user) VALUES ($1, $2) RETURNING id_application"

	var id uint
	err := postgre.conn.DB.QueryRow(query, id_user).Scan(&id)
	
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	return id, nil
}