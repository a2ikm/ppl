package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
)

func GetPoem(url string) {
  doc, _ := goquery.NewDocument(url)
  doc.Find(".content h1, .content .content-body").Each(func(_ int, s *goquery.Selection) {
    text := s.Text()
    fmt.Println(text)
  })
}

func main() {
  url := "https://www.pplog.net/u/ikm"
  GetPoem(url)
}
