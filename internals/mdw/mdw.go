package mdw

import (
	"github.com/julienschmidt/httprouter"
)

// Middleware type
type Middleware func(httprouter.Handle) httprouter.Handle

// Chain runs list of middleware and executes in sequential order.
// After all middleware is executed, handler is called
func Chain(f httprouter.Handle, m ...Middleware) httprouter.Handle {
	// if our chain is done, run the handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise run recursively over nested handlers
	return m[0](Chain(f, m[1:]...))
}
