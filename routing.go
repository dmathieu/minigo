package main

import (
  "code.google.com/p/gorilla/mux"
  "net/http"
)

func welcome(res http.ResponseWriter, req *http.Request) {
  http.ServeFile(res, req, "public"+req.URL.Path)
}

func minify(res http.ResponseWriter, req *http.Request) {
  http.Redirect(res, req, "/", 301)
}

func router() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/", welcome).Methods("GET")
  router.HandleFunc("/minify", minify).Methods("POST")
  return router
}
