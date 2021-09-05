package mdw

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SecureHeaders - adds security headers
func SecureHeaders(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		w.Header().Del("X-Powered-By")
		// add other security related headers
		next(w, r, p)
	}
}
