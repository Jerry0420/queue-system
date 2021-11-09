package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/jerry0420/queue-system/backend/domain"
	"github.com/jerry0420/queue-system/backend/logging"
	"github.com/jerry0420/queue-system/backend/presenter"
)

type middleware struct {
	storeUsecase domain.StoreUsecaseInterface
	logger logging.LoggerTool
}

func NewMiddleware(router *mux.Router, logger logging.LoggerTool, storeUsecase domain.StoreUsecaseInterface) {
	mw := &middleware{storeUsecase, logger}
	router.Use(mw.loggingMiddleware)
	router.Use(mw.authenticationMiddleware)
}

func (mw *middleware) loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		randomUUID := uuid.New().String()
        ctx := context.WithValue(r.Context(), "requestID", randomUUID)
        
        responseWrapper := &presenter.ResponseWrapper{ResponseWriter: w, Buffer: &bytes.Buffer{}}

		r = r.WithContext(ctx)
		next.ServeHTTP(responseWrapper, r)
        
        var wrappedResponse *presenter.ResponseFormat
        json.Unmarshal(responseWrapper.Buffer.Bytes(), &wrappedResponse)
		io.Copy(w, responseWrapper.Buffer)

		if wrappedResponse != nil {
			// api routes will go here.
			ctx = context.WithValue(r.Context(), "code", wrappedResponse.Code)
		}
        ctx = context.WithValue(ctx, "duration", time.Since(start).Truncate(1 * time.Millisecond))
        
		r = r.WithContext(ctx)
        mw.logger.INFOf(r.Context(), "response")
    })
}

func (mw *middleware) authenticationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {		
		encryptToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwiZW1haWwiOiJvaG9oSG9zcGl0YWxAZ21haWwuY29tIiwibmFtZSI6Im9ob2giLCJzaWdua2V5X2lkIjoxMDgsImV4cCI6MTYzOTAzODgwNCwiaWF0IjoxNjM2NDQ2ODA0fQ.NTF0W32G94H-rBkVphbtH4HieY9vq__xP4_aTUkddf8"
		_, _ = mw.storeUsecase.ValidateToken(r.Context(), encryptToken)
		next.ServeHTTP(w, r)
    })
}

