package pgDB

import (
	"bytes"
	"context"
	"strconv"

	"github.com/jerry0420/queue-system/backend/domain"
)

func (repo *pgDBRepository) CreateCustomers(ctx context.Context, session domain.StoreSession, oldStatus string, newStatus string, customers []domain.Customer) error {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeOut)
	defer cancel()

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return domain.ServerError50002
	}
	defer tx.Rollback()

	err = repo.UpdateSessionWithTx(ctx, tx, session, oldStatus, newStatus)
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return err
	}

	variableCounts := 1
	var query bytes.Buffer
	var queryRowParams []interface{}
	query.WriteString("INSERT INTO customers (name, phone, queue_id, status) VALUES ")
	for index, customer := range customers {
		query.WriteString("($")
		query.WriteString(strconv.Itoa(variableCounts))
		query.WriteString(", $")
		query.WriteString(strconv.Itoa(variableCounts + 1))
		query.WriteString(", $")
		query.WriteString(strconv.Itoa(variableCounts + 2))
		query.WriteString(", $")
		query.WriteString(strconv.Itoa(variableCounts + 3))
		query.WriteString(")")
		variableCounts = variableCounts + 4
		queryRowParams = append(queryRowParams, customer.Name, customer.Phone, customer.QueueID, customer.Status)
		if index != len(customers)-1 {
			query.WriteString(", ")
		}
	}
	query.WriteString(" RETURNING id,name,phone,queue_id,created_at")

	stmt, err := tx.PrepareContext(ctx, query.String())
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return domain.ServerError50002
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, queryRowParams...)
	customers = customers[:0] // clear customers slice

	for rows.Next() {
		customer := domain.Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.QueueID, &customer.CreatedAt)
		if err != nil {
			repo.logger.ERRORf("error %v", err)
			return domain.ServerError50002
		}
		customer.Status = domain.CustomerStatus.NORMAL
		customers = append(customers, customer)

	}
	defer rows.Close()

	err = tx.Commit()
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return domain.ServerError50002
	}
	return nil
}

func (repo *pgDBRepository) UpdateCustomer(ctx context.Context, oldStatus string, newStatus string, customer *domain.Customer) error {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeOut)
	defer cancel()

	query := `UPDATE customers SET status=$1 WHERE id=$2 and status=$3`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return domain.ServerError50002
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, newStatus, customer.ID, oldStatus)
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return domain.ServerError50002
	}
	num, err := result.RowsAffected()
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return domain.ServerError50002
	}
	if num == 0 {
		return domain.ServerError40405
	}
	return nil
}