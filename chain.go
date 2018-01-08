//Package chain aims to implement a chain of handlers to use with the http package in the standard library
package chain

import "net/http"

type Handler func(http.Handler) http.Handler

//HandlerChain contains a chain of http.Handlers to use
type HandlerChain struct {
	chain []Handler
}

//New appends given handlers to a new chain
func New(h ...Handler) HandlerChain {
	return HandlerChain{append(([]Handler)(nil), h...)}
}

//Final wraps the handlers into the others, and lastly the given handler f to be used
func (h HandlerChain) Final(f http.Handler) http.Handler {
	if f == nil {
		f = http.DefaultServeMux
	}

	for i := range h.chain {
		f = h.chain[len(h.chain)-1-i](f)
	}
	return f
}
