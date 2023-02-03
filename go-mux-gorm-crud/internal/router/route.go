package router

import (
	"LHXHL/go-mux-gorm-crud/internal/controller"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/userinfo", controller.GetAllUsers).Methods("GET")
	r.HandleFunc("/userinfo/{id}", controller.GetUserById).Methods("GET")
	r.HandleFunc("/userinfo/delete/{id}", controller.DeleteUser).Methods("GET")
	r.HandleFunc("/userinfo/update/{id}", controller.UpdateUser).Methods("POST")
	r.HandleFunc("/userinfo/create", controller.CreateUser).Methods("POST")

}
