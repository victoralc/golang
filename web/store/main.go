package main

import (
	"net/http"
	routes2 "victor/golang/web/store/routes"
)

func main() {
	routes2.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
