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

	"github.com/jerry0420/queue-system/backend/logging"
	"github.com/jerry0420/queue-system/backend/presenter"
)

type middleware struct {
	logger logging.LoggerTool
}

func NewMiddleware(router *mux.Router, logger logging.LoggerTool) {
	mw := &middleware{logger}
	router.Use(mw.loggingMiddleware)
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
        ctx = context.WithValue(ctx, "duration", time.Since(start))
        
		r = r.WithContext(ctx)
        mw.logger.INFOf(r.Context(), "response")
    })
}