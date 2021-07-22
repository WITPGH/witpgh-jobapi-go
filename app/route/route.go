package route

import (
	"net/http"
	"witpgh-jobapi-go/app/controllers/employer"
	"witpgh-jobapi-go/app/route/middleware/cors"
	hr "witpgh-jobapi-go/app/route/middleware/httprouterwrapper"
	"witpgh-jobapi-go/app/route/middleware/logrequest"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// Load returns the routes and middleware
func LoadRoutes() http.Handler {
	return middleware(routes())
}

func routes() *httprouter.Router {
	r := httprouter.New()

	//employer
	r.POST("/employers/account/create", hr.Handler(alice.New(cors.Handler).ThenFunc(employer.CreateEmployerAccount)))

	return r
}

// *****************************************************************************
// Middleware
// *****************************************************************************
func middleware(h http.Handler) http.Handler {
	// Log every request
	h = logrequest.Handler(h)

	// Cors for swagger-ui
	h = cors.Handler(h)

	// Clear handler for Gorilla Context
	h = context.ClearHandler(h)

	return h
}
