package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"../infra"
	"../service"
	"github.com/gorilla/mux"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	//starting mongo pool
	session, err := mgo.Dial(infra.MONGODB_HOST)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	mPool := mgosession.NewPool(nil, session, infra.MONGODB_CONNECTION_POOL)
	defer mPool.Close()

	router := mux.NewRouter().StrictSlash(true)

	userRepository := infra.NewUserRepository(mPool, session, infra.MONGODB_DATABASE)
	userService := service.NewUserAppService(userRepository)

	widgetRepository := infra.NewWidgetRepository(mPool, session, infra.MONGODB_DATABASE)
	widgetService := service.NewWidgetAppService(widgetRepository)

	BuildUserHttpHandlers(router, userService)
	BuildWidgetHttpHandlers(router, widgetService)
	http.Handle("/", router)
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(infra.API_PORT),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}

// package main

// import (
// 	"log"
// 	"net/http"
// )

// func main() {

// 	router := NewAppRouter()
// 	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
// 	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	// })

// 	log.Fatal(http.ListenAndServe(":3000", router))

// }
