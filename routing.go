package main

import (
  "code.google.com/p/gorilla/mux"
  "net/http"
)

func welcome(res http.ResponseWriter, req *http.Request) {
  http.ServeFile(res, req, "public"+req.URL.Path)
}

func router() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/", welcome).Methods("GET")
  return router
}
