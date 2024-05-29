package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
  templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, page interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, page)
}

func NewTemplate() *Templates {
  return &Templates {
    templates: template.Must(template.ParseGlob("views/*.html")),
  }
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())

  e.Static("/css", "css")
  e.Static("/images", "images")

  e.Renderer = NewTemplate() 

  e.GET("/", func (c echo.Context) error {
    return c.NoContent(200)
  })

  e.Logger.Fatal(e.Start(":6969"))
}
