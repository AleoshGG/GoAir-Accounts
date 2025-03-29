package adapters

import (
	"GoAir-Accounts/API/Applications/domain"
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
	query := `SELECT (a.id_application, u.first_name, u.last_name, a.status_application, a.id_user) 
			  FROM applications a
			  INNER JOIN users u 
			  ON a.id_user = u.id_user 
			  WHERE a.id_application = $1`

	var rmsg domain.RabbitMessage
	rows, err := postgre.conn.FetchRows(query, id_application)

	if err != nil {
		fmt.Errorf("error al ejecutar la consulta: %w", err)
		return domain.RabbitMessage{}, err
	}

	defer rows.Close()

	if !rows.Next() {
        fmt.Println("No se pudieron obtener los datos.")
        return domain.RabbitMessage{}, err
    }

	if err := rows.Scan(&rmsg.Id_Application, &rmsg.First_name, &rmsg.Last_name, &rmsg.Status_application, &rmsg.Id_user); err != nil {
		fmt.Errorf("error al escanear el usuario: %w", err)
        return domain.RabbitMessage{}, err
    }
	fmt.Print(rmsg)
	return rmsg, nil
}
