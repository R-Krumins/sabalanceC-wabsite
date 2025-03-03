package routers

import (
	"net/http"
)

func PageRouter() http.Handler {
	fs := http.FileServer(http.Dir("./static"))
	return fs
}
