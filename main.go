package main

import (
  "net/http"
  "fmt"
)

func main() {
  logs := make(chan string, 10000)
  go runLogging(logs)

  handler := routerHandlerFunc(router())
  handler = wrapLogging(handler, logs)
  port := mustGetenv("PORT")

  logs <- fmt.Sprintf("serve at=start port=%s", port)

  err := http.ListenAndServe(":"+port, handler)
  check(err)
}
