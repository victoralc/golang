package routes

import (
	"net/http"
	controllers2 "victor/golang/web/store/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers2.Index)
	http.HandleFunc("/new", controllers2.New)
	http.HandleFunc("/insert", controllers2.Insert)
	http.HandleFunc("/delete", controllers2.Delete)
	http.HandleFunc("/edit", controllers2.Edit)
	http.HandleFunc("/update", controllers2.Update)
}