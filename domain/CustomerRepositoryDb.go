package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/leslesnoa/go-microservices-demo/errs"
	"github.com/leslesnoa/go-microservices-demo/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllsql := "select customer_id, name, city, zipcode, date_of_birth, status from customers;"
		err = d.client.Select(&customers, findAllsql)
	} else {
		findAllsql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=?;"
		err = d.client.Select(&customers, findAllsql, status)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			errs.NewNotFoundError("Customer Not Found")
		} else {
			logger.Error("Error while querying customer table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?;"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not Found")
		} else {
			logger.Error("Error while scanning custokmer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
