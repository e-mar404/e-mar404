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

func main() {
  e := echo.New()
  e.Use(middleware.Logger())

  e.Static("/css", "css")
  e.Static("/images", "images")
  e.Static("/blog", "blog")

  e.Renderer = NewTemplate() 

  e.GET("/", func (c echo.Context) error {
    return c.Render(200, "home", nil)
  })

  e.GET("/blogs", func (c echo.Context) error {
    blogList, err := blog.BlogList()
    if err != nil {
      log.Print(err)

      return c.Render(400, "blog-list", "something went wrong")
    }

    return c.Render(200, "blog-list", blogList) 
  })

  e.GET("/blogs/:blog-id", func (c echo.Context) error {
    file := c.Param("blog-id")
    
    log.Print(file)

    html, err := blog.BlogHtml(file)
    if err != nil {
      log.Print(err)

      return c.Render(400, "blog-view", "something went wrong")
    }

    return c.Render(200, "blog-view", string(html))
  })

  port := ":42069"
  envPort := os.Getenv("PORT")

  if envPort != "" {
      port = envPort
  }

  e.Logger.Fatal(e.Start(port))
}
