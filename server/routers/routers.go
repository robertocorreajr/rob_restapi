package routers

import "net/http"

// Routers é todas as rotas da API
type Routers struct {
	URI                  string
	Method               string
	Function             func(http.ResponseWriter, *http.Request)
	RequestAutentication bool
}
