package main

import (
  "strings"
  "bytes"
)

type Post struct {
  Title       string
  ContentHTML string `json:"content_html"`
  ExcerptHTML string `json:"excerpt_html"`
  URL         string
  ID          int
  Blog        Blog
  Author      Author `json:"user"`
}

func (p Post) Mustache() map[string]string {
  return map[string]string{
    "post_title":         p.Title,
    "post_excerpt_html":  p.SanitizedExcerptHTML(),
    "post_url":           p.URL,
    "post_content_html":  p.ContentHTML,
    "author_name":        p.Author.Name(),
    "author_given_name":  p.Author.GivenName,
    "author_family_name": p.Author.FamilyName,
    "author_bio_html":    p.Author.Bio,
    "author_website":     p.Author.Website,
    "author_twitter":     p.Author.Twitter,
  }
}

func (p Post) SanitizedExcerptHTML() string {
  output := ""
  s := p.ExcerptHTML

  // Shortcut strings with no tags in them
  if !strings.ContainsAny(s, "<>") {
    output = s
  } else {
    // First remove line breaks etc as these have no meaning outside html tags (except pre)
    // this means pre sections will lose formatting... but will result in less uninentional paras.
    s = strings.Replace(s, "\n", "", -1)

    // Then replace line breaks with newlines, to preserve that formatting
    s = strings.Replace(s, "</p>", "\n", -1)
    s = strings.Replace(s, "<br>", "\n", -1)
    s = strings.Replace(s, "</br>", "\n", -1)
    s = strings.Replace(s, "\n", " ", -1)

    // Walk through the string removing all tags
    b := bytes.NewBufferString("")
    inTag := false
    for _, r := range s {
      switch r {
      case '<':
        inTag = true
      case '>':
        inTag = false
      default:
        if !inTag {
          b.WriteRune(r)
        }
      }
    }
    output = b.String()
  }

  return output
}
