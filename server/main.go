package main

import (
	"io"
	"log"
	"os"
	"text/template"

	"e-mar404/website/server/blog"

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

type PageData struct {
}

func NewPageData() PageData {
  return PageData {
  }
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())

  e.Static("/css", "css")
  e.Static("/images", "images")
  e.Static("blog", "blog")

  e.Renderer = NewTemplate() 

  page := NewPageData()

  e.GET("/", func (c echo.Context) error {
    return c.Render(200, "home", page)
  })

  e.GET("/blog", func (c echo.Context) error {
    html, err := blog.BlogHtml()
    if err != nil {
      log.Fatal(err)
    }

    return c.Render(200, "blog", string(html)) 
  })

  port := ":42069"
  envPort := os.Getenv("PORT")

  if envPort != "" {
      port = envPort
  }

  e.Logger.Fatal(e.Start(port))
}
