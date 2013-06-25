package main

import (
  "code.google.com/p/gorilla/mux"
  "net/http"
)

func routerHandlerFunc(router *mux.Router) http.HandlerFunc {
  return func(res http.ResponseWriter, req *http.Request) {
    router.ServeHTTP(res, req)
  }
}

type statusCapturingResponseWriter struct {
  status int
  http.ResponseWriter
}

func (w statusCapturingResponseWriter) WriteHeader(s int) {
  w.status = s
  w.ResponseWriter.WriteHeader(s)
}
