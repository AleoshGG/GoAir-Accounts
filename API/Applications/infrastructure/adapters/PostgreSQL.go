package adapters

import (
	"GoAir-Accounts/API/Applications/domain"
	"GoAir-Accounts/API/core"
	"database/sql"
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

func (postgre *PostgreSQL) CreateApplication(id_user int) (domain.RabbitMessage, error) {
	query := "INSERT INTO applications (status_application, id_user) VALUES ('requested', $1) RETURNING id_application"

	var id uint
	err := postgre.conn.DB.QueryRow(query, id_user).Scan(&id)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return domain.RabbitMessage{}, err
	}

	dataMessage, err := getDataForRabbit(postgre, int(id))
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return domain.RabbitMessage{}, err
	}

	return dataMessage, nil
	
}

func getDataForRabbit(postgre *PostgreSQL, id_application int) (domain.RabbitMessage, error) {
	query := `SELECT a.id_application, u.first_name, u.last_name, a.status_application, a.id_user
			  FROM applications a
			  INNER JOIN users u 
			  ON a.id_user = u.id_user 
			  WHERE a.id_application = $1`

	var rmsg domain.RabbitMessage
	err := postgre.conn.DB.QueryRow(query, id_application).Scan(
        &rmsg.Id_application,
        &rmsg.First_name,
        &rmsg.Last_name,
        &rmsg.Status_application,
        &rmsg.Id_user,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return domain.RabbitMessage{}, fmt.Errorf("aplicación %d no encontrada", id_application)
        }
        return domain.RabbitMessage{}, fmt.Errorf("error al escanear datos: %w", err)
    }

    return rmsg, nil
}

func (postgres *PostgreSQL) GetApplicationByUser(id_user int) []domain.Application {
	query := `SELECT * FROM applications WHERE id_user = $1`

	var applications []domain.Application

	rows, err := postgres.conn.DB.Query(query, id_user)

	if err != nil {
        fmt.Println("No se pudieron obtener los datos.", err)
        return []domain.Application{}
    }

	defer rows.Close()

	for rows.Next() {
		var a domain.Application
		
		// Escanear los valores de la fila
		err := rows.Scan(&a.Id_application, &a.Status_application, &a.Id_user)
		if err != nil {
			// Manejar error al escanear la fila
			fmt.Println("Error al escanear la fila:", err)
			return []domain.Application{}
		}
		applications = append(applications, a)
	}

	// Verifica errores después de iterar
    if err = rows.Err(); err != nil {
        fmt.Println("Error al recorrer las filas:", err)
        return nil
    }

	return applications
}