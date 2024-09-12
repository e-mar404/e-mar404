package blog 

import (
	"fmt"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func fileContents(fileName string) ([]byte, error) {
  file, err := os.Open("./blogs/" + fileName)
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

  return buffer, nil
}

func mdToHTML(md []byte) []byte {
  extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
  p := parser.NewWithExtensions(extensions)
  doc := p.Parse(md)

  htmlFlags := html.CommonFlags | html.HrefTargetBlank
  opts := html.RendererOptions{ Flags: htmlFlags }
  renderer := html.NewRenderer(opts)

  return markdown.Render(doc, renderer)
}
