package controller

import (
	"LHXHL/go-mux-gorm-crud/internal/model"
	"LHXHL/go-mux-gorm-crud/internal/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func respondWithJson(w http.ResponseWriter, ok int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ok)
	json.NewEncoder(w).Encode(res)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetAllUsers()
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, users)
	}
	respondWithJson(w, http.StatusInternalServerError, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, err)
	}
	user, err := model.GetUserById(id)
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJson(w, http.StatusInternalServerError, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	log.Println(r.Body)
	err := util.ParseBody(r, &user)
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, err)
	}
	log.Println(user)
	newUser, err1 := model.CreateUser(user)
	if err1 != nil {
		respondWithJson(w, http.StatusInternalServerError, err1)
		return

	}
	respondWithJson(w, http.StatusInternalServerError, newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, err)
		return
	}

	err1 := model.DeleteUser(id)
	if err1 != nil {
		respondWithJson(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJson(w, http.StatusInternalServerError, "delete success!")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithJson(w, http.StatusInternalServerError, err)
		return
	}
	newUser := model.User{}
	err1 := util.ParseBody(r, &newUser)
	if err1 != nil {
		respondWithJson(w, http.StatusInternalServerError, err)
		return
	}
	user, err2 := model.UpdateUser(id, newUser)
	if err2 != nil {
		respondWithJson(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJson(w, http.StatusInternalServerError, user)
	return
}
