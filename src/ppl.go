package main

import (
  "fmt"
  "os"
  "strings"
  "github.com/PuerkitoBio/goquery"
)

func GetPoem(url string) {
  doc, _ := goquery.NewDocument(url)
  doc.Find(".content h1, .content .content-body").Each(func(_ int, s *goquery.Selection) {
    text := s.Text()
    fmt.Println(text)
  })
}

func GetUrl(user string) string {
  baseUrl := "https://www.pplog.net/u/"
  return strings.Join([]string{ baseUrl, user }, "")
}

func main() {
  url := GetUrl(os.Args[1])
  GetPoem(url)
}
