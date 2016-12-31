package requestID

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
)

const (
	HeaderKey  = "X-Request-ID"
	ContextKey = "requestID"
)

// Request ID Middleware.
// Checks the X-Request-ID header. If not found,
// generates a new uuid, and inserts whichever
// on the context before calling the next function.
// Should generally be the outermost middleware, so that
// all other middlewares have a request id available.
func RequestID(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		reqID := r.Header.Get(HeaderKey)
		if reqID == "" {
			reqID = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), ContextKey, reqID)
		fn(w, r.WithContext(ctx), p)
	}
}
