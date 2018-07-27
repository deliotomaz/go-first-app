package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../domain"
	"../service"
	"../utils"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//Configs for middleware  CORS and JWT validation
var middleWare = negroni.New(
	negroni.HandlerFunc(utils.Cors),
	negroni.HandlerFunc(utils.ValidateMiddleware),
)

type UserHttpHandlers struct {
	userService service.UserAppService
}
type WidgetHttpHandlers struct {
	widgetService service.WidgetAppService
}

//BuildUserHttpHandlers ... Build user routes
func BuildUserHttpHandlers(r *mux.Router, userService service.UserAppService) {
	hnd := &UserHttpHandlers{
		userService: userService,
	}
	r.Handle("/users/count", middleWare.With(negroni.Wrap(http.HandlerFunc(hnd.CountUsers)))).Methods("GET", "OPTIONS")
	r.Handle("/users", middleWare.With(negroni.Wrap(http.HandlerFunc(hnd.ListUsers)))).Methods("GET", "OPTIONS")
	r.Handle("/users/{id}", middleWare.With(negroni.Wrap(http.HandlerFunc(hnd.GetUserById)))).Methods("GET", "OPTIONS")

}

func BuildWidgetHttpHandlers(r *mux.Router, widgetService service.WidgetAppService) {
	hnd := &WidgetHttpHandlers{
		widgetService: widgetService,
	}
	r.Handle("/widgets/count", middleWare.With(negroni.Wrap(http.HandlerFunc(hnd.CountWidgets)))).Methods("GET", "OPTIONS")
	r.Handle("/widgets", middleWare.With(negroni.Wrap(http.HandlerFunc(hnd.ListWidgets)))).Methods("GET", "OPTIONS")
	r.Handle("/widgets/{id}", middleWare.With(negroni.Wrap(http.HandlerFunc(hnd.GetWidgetById)))).Methods("GET", "OPTIONS")
	r.Handle("/widgets", middleWare.With(negroni.Wrap(http.HandlerFunc(hnd.AddWidgets)))).Methods("POST", "OPTIONS")
	r.Handle("/widgets/{id}", middleWare.With(negroni.Wrap(http.HandlerFunc(hnd.UpdateWidgets)))).Methods("PUT", "OPTIONS")

}

//Handlers for user endpoint request
func (hnd *UserHttpHandlers) CountUsers(w http.ResponseWriter, r *http.Request) {
	tot, err := hnd.userService.GetCount()
	if err == nil {
		utils.ResponseOK(w, map[string]int{"Total": tot})

	} else {
		utils.ResponseBadRequest(w)
	}
}
func (hnd *UserHttpHandlers) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := hnd.userService.ListUsers()
	if err == nil {
		json.NewEncoder(w).Encode(users)

	} else {
		utils.ResponseBadRequest(w)
	}
}
func (hnd *UserHttpHandlers) GetUserById(w http.ResponseWriter, r *http.Request) {

	queryString := mux.Vars(r)
	userIdStr := queryString["id"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err == nil {
		user, err := hnd.userService.FindUserById(userId)
		if err == nil {
			json.NewEncoder(w).Encode(user)
		} else {
			utils.ResponseNotFound(w)
		}

	} else {
		utils.ResponseNotFound(w)

	}

}

//Handlers for user endpoint request

//Handlers for widget endpoint request
func (hnd *WidgetHttpHandlers) CountWidgets(w http.ResponseWriter, r *http.Request) {
	tot, err := hnd.widgetService.GetCount()
	if err == nil {
		utils.ResponseOK(w, map[string]int{"Total": tot})

	} else {
		utils.ResponseBadRequest(w)
	}
}
func (hnd *WidgetHttpHandlers) ListWidgets(w http.ResponseWriter, r *http.Request) {
	widgets, err := hnd.widgetService.ListWidgets()
	if err == nil {
		json.NewEncoder(w).Encode(widgets)

	} else {
		utils.ResponseBadRequest(w)
	}
}
func (hnd *WidgetHttpHandlers) GetWidgetById(w http.ResponseWriter, r *http.Request) {

	queryString := mux.Vars(r)
	widgetIdStr := queryString["id"]
	widgetId, err := strconv.ParseInt(widgetIdStr, 10, 64)
	if err == nil {
		widget, err := hnd.widgetService.FindWidgetById(widgetId)
		if err == nil {
			json.NewEncoder(w).Encode(widget)
		} else {
			utils.ResponseNotFound(w)
		}

	} else {
		utils.ResponseNotFound(w)

	}

}

func (hnd *WidgetHttpHandlers) UpdateWidgets(w http.ResponseWriter, r *http.Request) {

	queryString := mux.Vars(r)
	widgetIdStr := queryString["id"]
	widgetId, err := strconv.ParseInt(widgetIdStr, 10, 64)

	if err == nil {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.ResponseBadRequest(w)
		}
		log.Println(string(body))
		var widgetObj domain.Widget
		err = json.Unmarshal(body, &widgetObj)
		if err != nil {
			utils.ResponseBadRequest(w)
		}

		widgetObj.ID = widgetId
		err = hnd.widgetService.AddUpdate(widgetObj)
		if err == nil {
			utils.ResponseOK(w, nil)
		} else {
			utils.ResponseBadRequest(w)
		}

	} else {
		utils.ResponseBadRequest(w)

	}

}
func (hnd *WidgetHttpHandlers) AddWidgets(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResponseBadRequest(w)
	}
	log.Println(string(body))
	var widgetObj domain.Widget
	err = json.Unmarshal(body, &widgetObj)
	if err != nil {
		utils.ResponseBadRequest(w)
	}

	widgetObj.ID = 0
	err = hnd.widgetService.AddUpdate(widgetObj)
	if err == nil {
		utils.ResponseOK(w, nil)
	} else {
		utils.ResponseBadRequest(w)
	}

}

//Handlers for widget endpoint request
