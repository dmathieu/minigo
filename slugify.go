package main

import (
  "crypto/md5"
  "encoding/base64"
  "strings"
)

type Slugify struct {
  url string
}

func (s Slugify) slugify() (string) {
  h := md5.New()
  bytestring := []byte(s.url)
  h.Write(bytestring)
  encrypted := h.Sum(nil)
  value     := base64.URLEncoding.EncodeToString(encrypted)
  trimmed   := strings.Trim(value, "=")
  return trimmed
}
