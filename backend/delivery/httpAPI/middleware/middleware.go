package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/jerry0420/queue-system/backend/delivery/httpAPI/presenter"
	"github.com/jerry0420/queue-system/backend/domain"
	"github.com/jerry0420/queue-system/backend/logging"
	"github.com/jerry0420/queue-system/backend/usecase"
)

type Middleware struct {
	usecase *usecase.Usecase
	logger  logging.LoggerTool
}

func NewMiddleware(router *mux.Router, logger logging.LoggerTool, usecase *usecase.Usecase) *Middleware {
	mw := &Middleware{usecase, logger}
	router.Use(mw.LoggingMiddleware)
	return mw
}

func (mw *Middleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: remove after dev...
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		start := time.Now()

		randomUUID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "requestID", randomUUID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		ctx = context.WithValue(r.Context(), "duration", time.Since(start).Truncate(1*time.Millisecond))
		if errorCode := w.Header().Get("Server-Code"); errorCode != "" {
			ctx = context.WithValue(ctx, "code", errorCode)
		} else {
			ctx = context.WithValue(ctx, "code", strconv.Itoa(200))
		}

		r = r.WithContext(ctx)
		mw.logger.INFOf(r.Context(), "response")
	})
}

func (mw *Middleware) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		normalToken := r.Header.Get("Authorization")
		tokenClaims, err := mw.usecase.VerifyNormalToken(r.Context(), normalToken)
		if err != nil {
			presenter.JsonResponse(w, nil, err)
			return
		}
		ctx := context.WithValue(r.Context(), domain.SignKeyTypes.NORMAL, tokenClaims)
		mw.logger.INFOf(ctx, "storeID: %d", tokenClaims.StoreID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// for customers....
func (mw *Middleware) SessionAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("Authorization")
		session, _, err := mw.usecase.GetSessionAndStoreBySessionId(r.Context(), sessionId)
		if err != nil {
			presenter.JsonResponse(w, nil, err)
			return
		}
		ctx := context.WithValue(r.Context(), domain.StoreSessionString, session)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
