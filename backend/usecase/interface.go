package usecase

import (
	"context"
	"time"

	"github.com/jerry0420/queue-system/backend/domain"
)

type UseCaseInterface interface {
	// store.go
	CreateStore(ctx context.Context, store *domain.Store, queues []domain.Queue) error
	GetStoreByEmail(ctx context.Context, email string) (domain.Store, error)
	CheckStoreExpirationStatus(store domain.Store, err error) (domain.Store, error)
	GetStoreWIthQueuesAndCustomersById(ctx context.Context, storeId int) (domain.StoreWithQueues, error)
	VerifyPasswordLength(password string) error
	EncryptPassword(password string) (string, error)
	ValidatePassword(ctx context.Context, incomingPassword string, password string) error
	CloseStore(ctx context.Context, store domain.Store) error
	GenerateToken(ctx context.Context, store domain.Store, signKeyType string, expireTime time.Time) (encryptToken string, err error)
	VerifyToken(
		ctx context.Context,
		encryptToken string,
		signKeyType string,
		getSignKey func(context.Context, int, string) (domain.SignKey, error),
	) (tokenClaims domain.TokenClaims, err error)
	GetSignKeyByID(ctx context.Context, signKeyID int, signKeyType string) (domain.SignKey, error)
	RemoveSignKeyByID(ctx context.Context, signKeyID int, signKeyType string) (domain.SignKey, error)
	GenerateEmailContentOfForgetPassword(passwordToken string, store domain.Store) (subject string, content string)
	UpdateStore(ctx context.Context, store *domain.Store, fieldName string, newFieldValue string) error
	TopicNameOfUpdateCustomer(storeId int) string

	// queue.go

	// customer.go
	CreateCustomers(ctx context.Context, session domain.StoreSession, oldStatus string, newStatus string, customers []domain.Customer) error
	UpdateCustomer(ctx context.Context, oldStatus string, newStatus string, customer *domain.Customer) error

	// session.go
	CreateSession(ctx context.Context, store domain.Store) (domain.StoreSession, error)
	UpdateSession(ctx context.Context, session domain.StoreSession, oldStatus string, newStatus string) error
	TopicNameOfUpdateSession(storeId int) string
	GetSessionAndStoreBySessionId(ctx context.Context, sessionId string) (domain.StoreSession, domain.Store, error)

}
