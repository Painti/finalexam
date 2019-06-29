package model

import "github.com/Painti/finalexam/database"

type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func (c *Customer) Create(conn *database.DBConnection) error {
	row, err := conn.Create(map[string]interface{}{
		"name":   c.Name,
		"email":  c.Email,
		"status": c.Status,
	}, "name", "email", "status")
	if err != nil {
		return err
	}
	sErr := row.Scan(&c.Name, &c.Email, &c.Status, &c.ID)
	if sErr != nil {
		return sErr
	}
	return nil
}

func (c Customer) Save(conn *database.DBConnection) error {
	err := conn.Update(c.ID, map[string]interface{}{
		"name":   c.Name,
		"email":  c.Email,
		"status": c.Status,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Customer) GetData(conn *database.DBConnection) error {
	row, err := conn.ReadByID(c.ID, "id", "name", "email", "status")
	if err != nil {
		return err
	}
	sErr := row.Scan(&c.ID, &c.Name, &c.Email, &c.Status)
	if sErr != nil {
		return sErr
	}
	return nil
}

func (c *Customer) Delete(conn *database.DBConnection) error {
	return conn.Delete(c.ID)
}

func GetAllCustomer(conn *database.DBConnection) ([]Customer, error) {
	var customers []Customer
	rows, err := conn.ReadAllData("id", "name", "email", "status")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := Customer{}
		sErr := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Status)
		if sErr != nil {
			return nil, sErr
		}
		customers = append(customers, c)
	}
	return customers, nil
}
