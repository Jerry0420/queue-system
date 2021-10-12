package main

import (
    "database/sql"    
    "fmt"
    "net/http" 
    _ "github.com/lib/pq"

	// "bytes"
	// "context"
	// "encoding/json"
	// "io"
	// "log"
    "github.com/gorilla/mux"
	"github.com/jerry0420/queue-system/config"
	"github.com/jerry0420/queue-system/logging"
	"github.com/jerry0420/queue-system/utils"
)

func main() {
    logger := logging.NewLogger([]string{"method", "url", "code", "sep", "requestID", "duration"}, false)
    
    serverConfig := config.NewConfig()
    dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", 
        serverConfig.POSTGRES_USER(), 
        serverConfig.POSTGRES_PASSWORD(), 
        serverConfig.POSTGRES_HOST(), 
        serverConfig.POSTGRES_PORT(), 
        serverConfig.POSTGRES_DB(),
        serverConfig.POSTGRES_SSL(),
    )
    
    db, err := sql.Open("postgres", dbConnectionString)
    if err != nil {
        logger.FATALf("db connection fail %v", err)
    }
    
    err = db.Ping()
    if err != nil {
        logger.FATALf("db ping fail %v", err)
    }
    
    defer func() {
		err := db.Close()
		if err != nil {
			logger.FATALf("db connection close fail %v", err)
		}
	}()

    router := mux.NewRouter()

    // unsupported route goes here.
    router.HandleFunc("/{rest_of_router}", func (w http.ResponseWriter, r *http.Request)  {
        utils.JsonResponse(w, nil, utils.ServerError40401)
    })

    err = http.ListenAndServe("0.0.0.0:8000", router)
    if err != nil {
        logger.FATALf("ListenAndServe http fail %v", err)
    }
}

// func middleware(next http.HandlerFunc) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         logger := logging.NewLogger([]string{"method", "url", "code", "sep", "requestID", "duration"}, false)        
//         ctx := context.WithValue(r.Context(), "requestID", "aaaaaaaaaa")
//         r = r.WithContext(ctx)
        
//         responseWrapper := &utils.ResponseWrapper{
//             ResponseWriter: w,
//             Buffer: &bytes.Buffer{},
//         }

//         next(responseWrapper, r)
        
//         var wrapperResponse *utils.ResponseFormat
//         json.Unmarshal(responseWrapper.Buffer.Bytes(), &wrapperResponse)
//         ctx = context.WithValue(r.Context(), "code", wrapperResponse.Code)
//         ctx = context.WithValue(ctx, "duration", 3)
        
//         io.Copy(w, responseWrapper.Buffer)
//         r = r.WithContext(ctx)
        
//         logger.INFOf(r.Context(), "hello world %d", 1234)
//     }
// }