package adapters

import (
	"GoAir-Accounts/API/Admin/domain/entities"
	"GoAir-Accounts/API/core"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
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

func (postgres *PostgreSQL) GetAdmin() entities.Admin {
	query := "SELECT * FROM admin"
	var admin []entities.Admin

	rows, _ := postgres.conn.FetchRows(query)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return entities.Admin{}
    }

	defer rows.Close()

	for rows.Next() {
		var a entities.Admin
		rows.Scan(&a.Password, &a.Email)

		admin = append(admin, a)
	}
	
	return admin[0]
}

func (postgres *PostgreSQL) CreatePlace(name string, id_user int) (uint, error){
	query := "INSERT INTO places (name) VALUES (?)"

	res, err := postgres.conn.ExecutePreparedQuery(query, name)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta 1: %v", err)
		return 0, err
	}

	id, _ := res.LastInsertId() 

	query = "INSERT INTO users_places (id_place, id_user) VALUES (?,?)"

	res, err = postgres.conn.ExecutePreparedQuery(query, id, id_user)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta 2: %v", err)
		return 0, err
	} 

	if err = postgres.CreateId(int(id)); err != nil {
		fmt.Println("Error: %v", err)
		return 0, err
	} 

	return uint(id), nil
}

func (postgres *PostgreSQL) SearchUser(last_name string) entities.User {
	query := "SELECT * FROM users WHERE last_name LIKE CONCAT('%', ?, '%')"
	var users []entities.User

	rows, _ := postgres.conn.FetchRows(query)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return entities.User{}
    }

	defer rows.Close()

	for rows.Next() {
		var u entities.User
		rows.Scan(&u.Id_user, &u.First_name, &u.Last_name, &u.Email, &u.Password)

		users = append(users, u)
	}
	
	return users[0]
}

func (postgres *PostgreSQL) CreateId(id_place int) (error) {
	id_mq135a, err := gonanoid.New() 
	if err != nil {
        fmt.Println("No se pudo generar el id de: MQ135A")
        return err
    }

	id_mq135b, err := gonanoid.New() 
	if err != nil {
        fmt.Println("No se pudo generar el id de: MQ135B")
        return err
    }

	id_dh11a, err := gonanoid.New() 
	if err != nil {
        fmt.Println("No se pudo generar el id de: DH11A")
        return err
    }

	id_dh11b, err := gonanoid.New() 
	if err != nil {
        fmt.Println("No se pudo generar el id de: DH11B")
        return err
    }

	query := "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES (?,?,?,?)"
	_, err = postgres.conn.ExecutePreparedQuery(query, id_mq135a, id_place, "air_quality", "MQ135")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaA: %v", err)
		return err
	}

	query = "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES (?,?,?,?)"
	_, err = postgres.conn.ExecutePreparedQuery(query, id_mq135b, id_place, "air_quality", "MQ135")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaB: %v", err)
		return err
	}

	query = "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES (?,?,?,?)"
	_, err = postgres.conn.ExecutePreparedQuery(query, id_dh11a, id_place, "temperature", "DH11")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaC: %v", err)
		return err
	}

	query = "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES (?,?,?,?)"
	_, err = postgres.conn.ExecutePreparedQuery(query, id_dh11b, id_place, "temperature", "DH11")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaD: %v", err)
		return err
	}

	query = "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES (?,?,?,?)"
	_, err = postgres.conn.ExecutePreparedQuery(query, id_dh11a, id_place, "humidity", "DH11")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaE: %v", err)
		return err
	}

	query = "INSERT INTO sensors (id_sensor, id_place, sensor_type, model) VALUES (?,?,?,?)"
	_, err = postgres.conn.ExecutePreparedQuery(query, id_dh11b, id_place, "humidity", "DH11")
	if err != nil {
		fmt.Println("Error al ejecutar la consultaF: %v", err)
		return err
	}

	return nil
}

func (postgres *PostgreSQL) GetIds(id_place int) []entities.Sensor {
	query := "SELECT * FROM sensors WHERE id_place = ?"
	var sensors []entities.Sensor

	rows, _ := postgres.conn.FetchRows(query, id_place)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return []entities.Sensor{}
    }

	defer rows.Close()

	for rows.Next() {
		var s entities.Sensor
		rows.Scan(&s.Id_sensor, &s.Id_place, &s.Sensor_type, &s.Model, &s.Installation_date)

		sensors = append(sensors, s)
	}
	
	return sensors
}

func (postgres *PostgreSQL) GetPlaces(id_user int) []entities.Place {
	query := "SELECT * FROM users_places WHERE id_user = ?"
	var users_places []entities.UsersPlaces
	var places []entities.Place

	rows, _ := postgres.conn.FetchRows(query)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return []entities.Place{}
    }

	defer rows.Close()

	for rows.Next() {
		var up entities.UsersPlaces
		rows.Scan(&up.Id_place, &up.Id_user)

		users_places = append(users_places, up)
	}
	
	for _, up := range users_places {
		query = "SELECT * FROM places WHERE id_place = ?"
		rows, _ := postgres.conn.FetchRows(query, up.Id_place)

		if rows == nil {
			fmt.Println("No se pudieron obtener los datos.")
			return []entities.Place{}
		}
	
			var p entities.Place
		for rows.Next() {
			rows.Scan(&p.Id_place, &p.Name, &p.Timestamp)
	
			places = append(places, p)
		}
	}

	return places
}

func (postgres *PostgreSQL) DeletePlace(id_place int) (uint, error) {
	query := "DELETE FROM places WHERE id_place = ?"

	res, err := postgres.conn.ExecutePreparedQuery(query, id_place)

	if err != nil {
		fmt.Println("Error al ejecutar la consultaF: %v", err)
		return 0, err
	}

	rowsAffected, _ := res.RowsAffected()
	
	return uint(rowsAffected), nil
}
