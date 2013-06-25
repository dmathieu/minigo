package main

import (
  "os"
)

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func mustGetenv(k string) string {
  v := os.Getenv(k)
  if v == "" {
    panic("missing " + k)
  }
  return v
}
