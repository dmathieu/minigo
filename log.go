package main


import (
  "net/http"
  "fmt"
  "time"
)

func runLogging(logs chan string) {
  for log := range logs {
    fmt.Println(log)
  }
}

func wrapLogging(f http.HandlerFunc, logs chan string) http.HandlerFunc {
  return func(res http.ResponseWriter, req *http.Request) {
    wres := statusCapturingResponseWriter{-1, res}
    start := time.Now()
    f(wres, req)
    method := req.Method
    path := req.URL.Path
    elapsed := float64(time.Since(start)) / 1000000.0
    logs <- fmt.Sprintf("request at=finish method=%s path=%s status=%d elapsed=%f",
    method, path, wres.status, elapsed)
  }
}
