package serv

import (
	// "github.com/gorilla/mux"
	// "go.uber.org/zap"
	"fmt"
	"net/http"
	"time"
)

func (s *server) accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		fmt.Printf("\n\t14")
		next.ServeHTTP(w, r)

		fmt.Printf("\n\t16")

		s.logger.Infow(
			"access",
			"method", r.Method,
			"remote_addr", r.RemoteAddr,
			"url", r.URL.Path,
			"processing_time", time.Since(start).String(),
		)
	})
}

// TODO refactor!!!
func (s *server) panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("panicMiddleware", r.URL.Path)
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("\n\t36")
				fmt.Println("recovered", err)
				// http.Error(w, "Internal server error", 500)
				fmt.Printf("\n\t39")
			}
		}()
		fmt.Printf("\n\t42")
		next.ServeHTTP(w, r)
		fmt.Printf("\n\t44")
	})
}

// // Define our struct
// type authenticationMiddleware struct {
// 	tokenUsers map[string]string
// }

// // Initialize it somewhere
// func (amw *authenticationMiddleware) Populate() {
// 	amw.tokenUsers["00000000"] = "user0"
// 	amw.tokenUsers["aaaaaaaa"] = "userA"
// 	amw.tokenUsers["05f717e5"] = "randomUser"
// 	amw.tokenUsers["deadbeef"] = "user0"
// }

// // Middleware function, which will be called for each request
// func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         token := r.Header.Get("X-Session-Token")

//         if user, found := amw.tokenUsers[token]; found {
//         	// We found the token in our map
//         	log.Printf("Authenticated user %s\n", user)
//         	// Pass down the request to the next middleware (or final handler)
//         	next.ServeHTTP(w, r)
//         } else {
//         	// Write an error and stop the handler chain
//         	http.Error(w, "Forbidden", http.StatusForbidden)
//         }
//     })
// }
