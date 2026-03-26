package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/singhparbjyot/banking/errors"
	"github.com/singhparbjyot/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, error) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "Select customer_id,name,city,zipcode,date_of_birth,status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "	Select customer_id,name,city,zipcode,date_of_birth,status from customers where status = $1"

		err = d.client.Select(&customers, findAllSql, status)
	}
	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, err
	}
	//err = sqlx.StructScan(rows, &customers)
	//Here we are looping the rows of the query return to show the customers with their information
	//if err != nil {
	//	logger.Error("Error while scanning customer table " + err.Error())
	//	return nil, err
	//}
	return customers, nil

}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errors.AppError) {
	customerById := "Select customer_id,name,city,zipcode,date_of_birth::TEXT,status from customers where customer_id = $1"

	var c Customer
	err := d.client.Get(&c, customerById, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errors.NewNotUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

// Func which open the connection with the database and set the settings for the db.
func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{client: dbClient}
}
