package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leslesnoa/go-microservices-demo/errs"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllsql := "select customer_id, name, city, zipcode, date_of_birth, status from customers;"
		rows, err = d.client.Query(findAllsql)
	} else {
		findAllsql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=?;"
		rows, err = d.client.Query(findAllsql, status)
	}

	// findAllsql := "select customer_id, name, city, zipcode, date_of_birth, status from customers;"

	// rows, err := d.client.Query(findAllsql)
	if err != nil {
		if err == sql.ErrNoRows {
			errs.NewNotFoundError("Customer Not Found")
		} else {
			log.Println("Error while querying customer table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("database scanning customer error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?;"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	if err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status); err != nil {
		if err == sql.ErrNoRows {
			// return nil, errors.New("Customer not Found")
			return nil, errs.NewNotFoundError("Customer not Found")
		} else {
			log.Println("Error while scanning custokmer " + err.Error())
			// return nil, errors.New("Unexpected database error")
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
