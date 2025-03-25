package adapters

import (
	"GoAir-Accounts/API/Places/domain"
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

func (postgres *PostgreSQL) CreatePlace(p domain.Place)(uint, error) {
	query := "INSERT INTO places (name) VALUES (?)"

	res, err := postgres.conn.ExecutePreparedQuery(query, p.Name)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	id, _ := res.LastInsertId() 

	return uint(id), nil
}

func (postgres *PostgreSQL) DeletePlace(id_place int) (uint, error) {
	query := "DELETE FROM places WHERE id_place = ?"

	res, err := postgres.conn.ExecutePreparedQuery(query, id_place)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	rows, _ := res.RowsAffected() 

	return uint(rows), nil
}
