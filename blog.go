package main

type Blog struct {
  ID          int
  Name        string
  Language    string
  URL         string
  AccentColor string `json:"accent_color"`
}
