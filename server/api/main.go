package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"https://github.com/deliotomaz/go-first-app/config"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
)

func main(){
	
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()
}