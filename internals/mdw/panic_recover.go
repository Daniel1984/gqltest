package mdw

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"
)

// PanicRecover - prevents server from crashing by adding recovery functionality
func PanicRecover(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer func() {
			// Use the builtin recover function to check if there has been a
			// panic or not. If there has...
			if err := recover(); err != nil {
				log.Printf("err stack:> %s\n", debug.Stack())
				// Set a "Connection: close" header on the response and return error
				log.Printf("err:> %s\n", err)
				w.Header().Set("Connection", "close")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Something went wrong. Try again later"))
			}
		}()
		next(w, r, p)
	}
}
