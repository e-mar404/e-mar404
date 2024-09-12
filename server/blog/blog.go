package blog

import (
	"math/rand"
	"os"
	"strings"
)

type BlogData struct {
  Id string
  Name string
}

type BlogDataList struct {
  Blogs []BlogData
}

func BlogList() (BlogDataList, error) {
  fileList, err := os.ReadDir("./blogs")
  if err != nil {
    return BlogDataList{}, err
  }
  
  blogDataList := BlogDataList{}

  for _, file := range fileList {
    id, found := strings.CutSuffix(file.Name(), ".md")
    if !found {
      id = string(rand.Intn(10))
    }

    blogDataList.Blogs = append(blogDataList.Blogs, BlogData{
      Id: id,
      Name: file.Name(),
    })
  }

  return blogDataList, nil
}

func BlogHtml(file string) ([]byte, error) {
  buffer, err := fileContents(file + ".md") 
  if err != nil {
    return nil, err
  }

  html:= mdToHTML(buffer)

  return html, nil
}

