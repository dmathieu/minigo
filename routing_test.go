package main

import(
  "testing"
  "net/http"
  "net/http/httptest"
)

func TestRoutingHome(t *testing.T) {
  handler  := routerHandlerFunc(router())
  req, err := http.NewRequest("GET", "http://example.com/", nil)

  if err != nil {
    t.Fatal(err)
  }
  w := httptest.NewRecorder()
  handler(w, req)

  if w.Code != 200 {
    t.Fatal("HTTP status expected 200. Was ", w.Code);
  }
}

func TestRoutingCreateMini(t *testing.T) {
  handler  := routerHandlerFunc(router())
  req, err := http.NewRequest("POST", "http://example.com/minify", nil)

  if err != nil {
    t.Fatal(err)
  }
  w := httptest.NewRecorder()
  handler(w, req)

  if w.Code != 301 {
    t.Fatal("HTTP status expected 200. Was ", w.Code);
  }
}
