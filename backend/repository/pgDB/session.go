package pgDB

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jerry0420/queue-system/backend/domain"
)

func (repo *PgDBRepository) CreateSession(ctx context.Context, store domain.Store) (domain.StoreSession, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeOut)
	defer cancel()

	session := domain.StoreSession{StoreId: store.ID, StoreSessionStatus: domain.StoreSessionStatus.NORMAL}

	query := `INSERT INTO store_sessions (store_id) VALUES ($1) RETURNING id`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return session, domain.ServerError50002
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, store.ID)
	err = row.Scan(&session.ID)
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return session, domain.ServerError40904
	}
	return session, nil
}

func (repo *PgDBRepository) UpdateSessionStatus(ctx context.Context, session *domain.StoreSession, oldStatus string, newStatus string) error {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeOut)
	defer cancel()

	query := `UPDATE store_sessions SET status=$1 WHERE id=$2 and status=$3`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return domain.ServerError50002
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, newStatus, session.ID, oldStatus)
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
		return domain.ServerError40404
	}
	return nil
}

func (repo *PgDBRepository) UpdateSessionWithTx(ctx context.Context, tx *sql.Tx, session domain.StoreSession, oldStatus string, newStatus string) error {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeOut)
	defer cancel()
	query := `UPDATE store_sessions SET status=$1 WHERE id=$2 and status=$3`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		repo.logger.ERRORf("error %v", err)
		return domain.ServerError50002
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, newStatus, session.ID, oldStatus)
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
		return domain.ServerError40404
	}
	return nil
}

func (repo *PgDBRepository) GetSessionAndStoreBySessionId(ctx context.Context, sessionId string) (domain.StoreSession, domain.Store, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeOut)
	defer cancel()

	session := domain.StoreSession{}
	store := domain.Store{}

	query := `SELECT stores.id, stores.created_at, store_sessions.status 
				FROM store_sessions
				INNER JOIN stores ON stores.id = store_sessions.store_id WHERE store_sessions.id=$1`
	row := repo.db.QueryRowContext(ctx, query, sessionId)
	err := row.Scan(&store.ID, &store.CreatedAt, &session.StoreSessionStatus)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		repo.logger.ERRORf("error %v", err)
		return session, store, domain.ServerError40404
	case err != nil:
		repo.logger.ERRORf("error %v", err)
		return session, store, domain.ServerError50002
	}
	session.ID = sessionId
	session.StoreId = store.ID
	return session, store, nil
}
