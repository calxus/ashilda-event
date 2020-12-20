package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gtadam/ashilda-common"
)

type Router struct {
	basePath string
	database models.Database
	mux		 *mux.Router
}

func NewRouter(bp string) *Router {
	return &Router {
		basePath: bp,
		database: *models.NewDatabase(),
		mux:	  mux.NewRouter().StrictSlash(true),
	}
}

func (rt *Router) getEvents(w http.ResponseWriter, r *http.Request) {
	statement := models.NewDatabaseSelect(TABLE)
	statement.AddColumn(ID_FIELD)
	statement.AddColumn(NAME_FIELD)
	rows, _ := rt.database.ExecuteSelect(statement)
	events := []Event{}
	for rows.Next() {
		event := Event{}
		event.populate(rows)
		events = append(events, event)
	}
	rows.Close()
	json, _ := json.Marshal(events)
	fmt.Fprintf(w, string(json))
}

func (rt *Router) getEvent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	statement := models.NewDatabaseSelect(TABLE)
	statement.AddColumn(ID_FIELD)
	statement.AddColumn(NAME_FIELD)
	statement.AddCondition(ID_FIELD, models.EQUALS, id)
	rows, _ := rt.database.ExecuteSelect(statement)
	event := Event{}
	rows.Next()
	event.populate(rows)
	rows.Close()
	if (event.Event_id == 0) {
		w.WriteHeader(404)
		return
	}
	json, _ := json.Marshal(event)
	fmt.Fprintf(w, string(json))
}

func (rt *Router) putEvent(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	event := Event{}
	json.Unmarshal(body, &event)
	statement := models.NewDatabaseUpdate(TABLE)
	statement.AddStatement(NAME_FIELD, event.Name)
	statement.AddCondition(ID_FIELD, models.EQUALS, strconv.Itoa(event.Event_id))
	rt.database.ExecuteUpdate(statement)
}

func (rt *Router) postEvent(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	event := Event{}
	json.Unmarshal(body, &event)
	statement := models.NewDatabaseInsert(TABLE)
	statement.AddEntry(NAME_FIELD, event.Name)
	rt.database.ExecuteInsert(statement)
}

func (rt *Router) deleteEvent(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	event := Event{}
	json.Unmarshal(body, &event)
	statement := models.NewDatabaseDelete(TABLE)
	statement.AddCondition(ID_FIELD, models.EQUALS, strconv.Itoa(event.Event_id))
	rt.database.ExecuteDelete(statement)
}

func (rt *Router) populateRoutes() {
	rt.database.Connect()
	rt.mux.HandleFunc(rt.basePath+"/events", rt.getEvents).Methods("GET")
	rt.mux.HandleFunc(rt.basePath+"/event/{id:[0-9]+}", rt.getEvent).Methods("GET")
	rt.mux.HandleFunc(rt.basePath+"/event", rt.putEvent).Methods("PUT")
	rt.mux.HandleFunc(rt.basePath+"/event", rt.postEvent).Methods("POST")
	rt.mux.HandleFunc(rt.basePath+"/event", rt.deleteEvent).Methods("DELETE")
}