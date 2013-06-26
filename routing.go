package main

import (
  "code.google.com/p/gorilla/mux"
  "net/http"
  "html/template"
)

func welcome(res http.ResponseWriter, req *http.Request) {
  http.ServeFile(res, req, "public/index.html")
}

func minify(res http.ResponseWriter, req *http.Request) {
  url         := req.FormValue("url")
  persistance := new(Persistance)

  slug, err := persistance.minify(url)
  check(err)

  type Data struct {
    Slug string
  }

  t, err := template.ParseFiles("public/minify.html")
  data   := Data{slug}
  check(err)
  t.Execute(res, data)
}

func redirect(res http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  slug := vars["slug"]
  persistance := new(Persistance)
  url, err := persistance.getUrl(slug)
  check(err)

  if url != "" {
    http.Redirect(res, req, url, 301)
  } else {
    http.NotFound(res, req)
  }
}

func router() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/", welcome).Methods("GET")
  router.HandleFunc("/minify", minify).Methods("POST")
  router.HandleFunc("/{slug}", redirect).Methods("GET")
  return router
}
