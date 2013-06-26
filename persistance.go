package main

import (
  "github.com/lib/pq"
  "database/sql"
)

type Persistance struct { }

func (p Persistance) connect() (*sql.DB) {
  url := mustGetenv("DATABASE_URL")
  conn, err := pq.ParseURL(url)
  check(err)

  db, err := sql.Open("postgres", conn)
  check(err)

  return db
}

func (p Persistance) createTables(db *sql.DB) {
  _, err := db.Exec(`
  CREATE TABLE IF NOT EXISTS minified (url text, slug text);`)
  check(err)
}

func (p Persistance) minify(url string) (string, error) {
  db := p.connect()
  defer db.Close()

  p.createTables(db)

  tx, err := db.Begin()
  check(err)

  stmt, err := tx.Prepare("INSERT INTO minified (url, slug) VALUES ($1, $2)")
  check(err)
  defer stmt.Close()

  slugify := Slugify{url}
  slug    := slugify.slugify()

  _, err = stmt.Exec(url, slug)
  check(err)

  tx.Commit()

  return slug, nil
}

func (p Persistance) getUrl(slug string) (string, error) {
  db := p.connect()
  defer db.Close()

  p.createTables(db)

  var url string
  err := db.QueryRow("SELECT url FROM minified WHERE slug = $1", slug).Scan(&url)

  switch {
  case err == sql.ErrNoRows:
  case err != nil:
    check(err)
  default:
  }

  return url, nil
}
