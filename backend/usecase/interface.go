package usecase

import (
	"context"
	"time"

	"github.com/jerry0420/queue-system/backend/domain"
)

type CustomerUseCaseInterface interface {
	CreateCustomers(ctx context.Context, session domain.StoreSession, oldStatus string, newStatus string, customers []domain.Customer) error
	UpdateCustomer(ctx context.Context, oldStatus string, newStatus string, customer *domain.Customer) error
}

type SessionUseCaseInterface interface {
	CreateSession(ctx context.Context, store domain.Store) (domain.StoreSession, error)
	UpdateSessionStatus(ctx context.Context, session *domain.StoreSession, oldStatus string, newStatus string) error
	TopicNameOfUpdateSession(storeId int) string
	GetSessionAndStoreBySessionId(ctx context.Context, sessionId string) (session domain.StoreSession, store domain.Store, err error)
}

type StoreUseCaseInterface interface {
	CreateStore(ctx context.Context, store *domain.Store, queues []domain.Queue) error
	SigninStore(ctx context.Context, email string, password string) (
		store domain.Store,
		token string,
		refreshTokenExpiresAt time.Time,
		err error,
	)
	RefreshToken(ctx context.Context, encryptedRefreshToken string) (
		store domain.Store,
		normalToken string,
		sessionToken string,
		tokenExpiresAt time.Time,
		err error,
	)
	CloseStore(ctx context.Context, store domain.Store) error
	ChunkStoresSlice(items [][][]string, chunkSize int) (chunks [][][][]string)
	CloseStoreRoutine(ctx context.Context) error
	ForgetPassword(ctx context.Context, email string) (store domain.Store, err error)
	UpdatePassword(ctx context.Context, passwordToken string, newPassword string) (store domain.Store, err error)
	UpdateStoreDescription(ctx context.Context, newDescription string, store *domain.Store) error
	VerifyPasswordLength(password string) error
	EncryptPassword(password string) (string, error)
	ValidatePassword(passwordInDb string, incomingPassword string) error
	GenerateToken(ctx context.Context, store domain.Store, signKeyType string, expireTime time.Time) (encryptToken string, err error)
	VerifyToken(ctx context.Context, encryptToken string, signKeyType string, withSignkeyPreserved bool) (tokenClaims domain.TokenClaims, err error)
	GenerateEmailContentOfForgetPassword(passwordToken string, store domain.Store) (subject string, content string)
	GenerateEmailContentOfCloseStore(storeName string, storeCreatedAt string) (subject string, content string)
	GenerateCsvFileNameAndContent(storeCreatedAt time.Time, storeName string, content [][]string) (date string, csvFileName string, csvContent []byte)
	TopicNameOfUpdateCustomer(storeId int) string
	GetStoreWithQueuesAndCustomersById(ctx context.Context, storeId int) (domain.StoreWithQueues, error)
	VerifyNormalToken(ctx context.Context, normalToken string) (tokenClaims domain.TokenClaims, err error)
	VerifySessionToken(ctx context.Context, sessionToken string) (store domain.Store, err error)
}
