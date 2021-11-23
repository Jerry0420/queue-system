package usecase

import (
	"context"

	"github.com/jerry0420/queue-system/backend/domain"
)

func (uc *usecase) CreateSession(ctx context.Context, store domain.Store) (domain.StoreSession, error) {
	session, err := uc.pgDBRepository.CreateSession(ctx, store)
	return session, err
}

func (uc *usecase) UpdateSession(ctx context.Context, session *domain.StoreSession, oldStatus string, newStatus string) error {
	err := uc.pgDBRepository.UpdateSession(ctx, session, oldStatus, newStatus)
	return err
}
