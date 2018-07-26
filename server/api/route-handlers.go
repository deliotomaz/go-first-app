package main

import (
	"net/http"

	"../service"
	"../utils"
	"github.com/gorilla/mux"
)

type UserHttpHandlers struct {
	userService service.UserAppService
}

func BuildUserHttpHandlers(r *mux.Router, userService service.UserAppService) {
	hnd := &UserHttpHandlers{
		userService: userService,
	}
	r.Handle("/users/count", http.HandlerFunc(hnd.CountUsers)).Methods("GET", "OPTIONS")

}
func (hnd *UserHttpHandlers) CountUsers(w http.ResponseWriter, r *http.Request) {
	tot, err := hnd.userService.GetCount()
	if err == nil {
		utils.ResponseOK(w, map[string]int{"Total": tot})

	} else {
		utils.ResponseBadRequest(w)
	}
}

// func (hnd *UserHttpHandlers) ListUsers(w http.ResponseWriter, r *http.Request) {
// 	users := domain.Users{
// 		User{Name: "Delio"},
// 	}
// 	json.NewEncoder(w).Encode(users)
// }
// func GetUser(w http.ResponseWriter, r *http.Request) {

// 	queryString := mux.Vars(r)
// 	userIdStr := queryString["id"]
// 	userId, err := strconv.ParseUint(userIdStr, 10, 64)
// 	if err == nil {
// 		utils.ResponseOK(w, &User{Name: "Delio", ID: userId})

// 	} else {
// 		utils.ResponseBadRequest(w)

// 	}

// }
