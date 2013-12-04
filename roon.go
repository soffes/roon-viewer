package main

import (
	"github.com/codegangsta/martini"
	"github.com/hoisie/mustache"
)

func main() {
	m := martini.Classic()
  m.Get("/", func() string {
    return "Index"
  })

  m.Get("/:slug", func(params martini.Params) string {
		post, error := Get("sam", params["slug"])
		if error != nil {
			return "error"
		}

		return mustache.RenderFile("templates/post.html.mustache", post.Mustache())
	})

	m.NotFound(func() string {
	  return "Not found."
	})

  m.Run()
}
