package routers

import "net/http"

var routerBook = []Routers{
	{
		URI:    "/api/book",
		Method: http.MethodGet,
		Function: func(rw http.ResponseWriter, r *http.Request) {

		},
		RequestAutentication: false,
	},
	{
		URI:    "/api/book/{bookId}",
		Method: http.MethodGet,
		Function: func(rw http.ResponseWriter, r *http.Request) {

		},
		RequestAutentication: false,
	},
}
