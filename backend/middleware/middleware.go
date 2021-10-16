package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

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
        ctx := context.WithValue(r.Context(), "requestID", "aaaaaaaaaa")
        r = r.WithContext(ctx)
        
        responseWrapper := &presenter.ResponseWrapper{ResponseWriter: w, Buffer: &bytes.Buffer{}}

		next.ServeHTTP(responseWrapper, r)
        
        var wrapperResponse *presenter.ResponseFormat
        json.Unmarshal(responseWrapper.Buffer.Bytes(), &wrapperResponse)
		if wrapperResponse != nil {
			// api routes will go here.
			ctx = context.WithValue(r.Context(), "code", wrapperResponse.Code)
		}
        ctx = context.WithValue(ctx, "duration", 3)
        
        io.Copy(w, responseWrapper.Buffer)
        r = r.WithContext(ctx)
        
        mw.logger.INFOf(r.Context(), "hello world %d", 1234)
    })
}