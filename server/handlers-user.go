package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"./utils"
	"github.com/gorilla/mux"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Name: "Delio"},
	}
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request) {

	queryString := mux.Vars(r)
	userIdStr := queryString["id"]
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err == nil {
		utils.ResponseOK(w, &User{Name: "Delio", ID: userId})

	} else {
		utils.ResponseBadRequest(w)

	}

}
