package mdw

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetCommonHeaders(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		next(w, r, p)
	}
}
