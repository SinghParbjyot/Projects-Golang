package domain

import (
	"strconv"

	"github.com/singhparbjyot/banking/errors"

	"github.com/jmoiron/sqlx"
	"github.com/singhparbjyot/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

// Save implements [AccountRepository].
func (d AccountRepositoryDb) Save(a Account) (*Account, *errors.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values ($1, $2, $3, $4, $5) RETURNING account_id"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errors.NewNotUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errors.NewNotUnexpectedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errors.AppError) {
	// Iniciar transacción
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errors.NewNotUnexpectedError("Unexpected database error")
	}

	// En caso de cualquier error, hacemos rollback
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Insertar transacción y obtener el id con RETURNING
	var transactionId int64
	err = tx.QueryRow(
		`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date)
         VALUES ($1,$2,$3,$4)
         RETURNING transaction_id`,
		t.AccountId, t.Amount, t.TransactionType, t.TransactionDate,
	).Scan(&transactionId)
	if err != nil {
		logger.Error("Error while inserting transaction: " + err.Error())
		return nil, errors.NewNotUnexpectedError("Unexpected database error")
	}

	// Actualizar saldo de la cuenta
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - $1 where account_id = $2`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + $1 where account_id = $2`, t.Amount, t.AccountId)
	}
	if err != nil {
		logger.Error("Error while updating account balance: " + err.Error())
		return nil, errors.NewNotUnexpectedError("Unexpected database error")
	}

	// Commit de la transacción
	err = tx.Commit()
	if err != nil {
		logger.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errors.NewNotUnexpectedError("Unexpected database error")
	}
	// A partir de aquí, err debe considerarse nil para que el defer no haga rollback
	err = nil

	// Obtener la información actualizada de la cuenta
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	// Asignar el ID de la transacción en el struct
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// OJO: aquí estás sobreescribiendo el amount de la transacción con el saldo actual.
	// Si tu intención es que Amount represente el MOVIMIENTO y no el SALDO, esto es un bug.
	t.Amount = account.Amount

	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errors.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = $1"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errors.NewNotUnexpectedError("Unexpected database error")
	}
	return &account, nil
}
func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
