package main

import (
  "fmt"
  "os"
  "strings"
  "github.com/PuerkitoBio/goquery"
)

type Poem struct {
  User  string
  Title string
  Body  string
}

func NewPoem(user string) *Poem {
  p := new(Poem)
  p.User = user
  p.Title, p.Body = p.getPoem()
  return p
}

func (p *Poem) getUrl() string {
  baseUrl := "https://www.pplog.net/u/"
  return strings.Join([]string{ baseUrl, p.User }, "")
}

func (p *Poem) getTitle(doc *goquery.Document) string {
  title := ""
  doc.Find(".content h1").First().Each(func(_ int, s *goquery.Selection) {
    title = strings.TrimSpace(s.Text())
  })
  return title
}

func (p *Poem) getBody(doc *goquery.Document) string {
  body := ""
  doc.Find(".content .content-body").First().Each(func(_ int, s *goquery.Selection) {
    body = strings.TrimSpace(s.Text())
  })
  return body
}

func (p *Poem) getPoem() (string, string) {
  doc, _ := goquery.NewDocument(p.getUrl())
  title  := p.getTitle(doc)
  body   := p.getBody(doc)

  return title, body
}

func main() {
  poem := NewPoem(os.Args[1])
  fmt.Printf("%s\n\n%s", poem.Title, poem.Body)
}
