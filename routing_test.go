package main

import(
  "testing"
  "net/http"
  "net/http/httptest"
  "strings"
  "fmt"
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
  req, err := http.NewRequest("POST", "http://example.com/minify", strings.NewReader("url=http://example.com"))
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded;")

  if err != nil {
    t.Fatal(err)
  }
  w := httptest.NewRecorder()
  handler(w, req)

  if w.Code != 200 {
    t.Fatal("HTTP status expected 200. Was ", w.Code);
  }
}

func TestRoutingRedirect(t *testing.T) {
  // Create a redirect
  persistance := new(Persistance)
  slug, err   := persistance.minify("http://google.com")
  url         := fmt.Sprintf("http://example.com/%s", slug)

  handler  := routerHandlerFunc(router())
  req, err := http.NewRequest("GET", url, nil)

  if err != nil {
    t.Fatal(err)
  }
  w := httptest.NewRecorder()
  handler(w, req)

  if w.Code != 301 {
    t.Fatal("HTTP status expected 301. Was ", w.Code);
  }
}

func TestRoutingRedirectMissing(t *testing.T) {
  handler  := routerHandlerFunc(router())
  req, err := http.NewRequest("GET", "http://example.com/foobarplouf", nil)

  if err != nil {
    t.Fatal(err)
  }
  w := httptest.NewRecorder()
  handler(w, req)

  if w.Code != 404 {
    t.Fatal("HTTP status expected 404. Was ", w.Code);
  }
}
