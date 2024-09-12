package blog

import (
  "os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

  "fmt"
)

func mdToHTML(md []byte) []byte {
  extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
  p := parser.NewWithExtensions(extensions)
  doc := p.Parse(md)

  htmlFlags := html.CommonFlags | html.HrefTargetBlank
  opts := html.RendererOptions{ Flags: htmlFlags }
  renderer := html.NewRenderer(opts)

  return markdown.Render(doc, renderer)
}

func BlogHtml() ([]byte, error) {
  file, err := os.Open("./blogs/test.md")
  if err != nil {
    fmt.Println(err)
    return nil, err
  }
  defer file.Close()
 
  fileInfo, err := file.Stat()
  if err != nil {
    fmt.Println(err)
    return nil, err
  }

  fileSize := fileInfo.Size()
  buffer := make([]byte, fileSize)

  _, readErr := file.Read(buffer)
  if readErr != nil {
    fmt.Println(err)
    return nil, err
  }
  
  html:= mdToHTML(buffer)

  fmt.Printf("--- Markdown:\n%s\n\n --- HTML:\n%s\n\n", buffer, html)

  return html, nil
}

