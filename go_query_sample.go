package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "strings"
  "io"
  "io/ioutil"
  "golang.org/x/text/transform"
  "golang.org/x/text/encoding/japanese"
)

func transformEncoding( rawReader io.Reader, trans transform.Transformer) (string, error) {
    ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
    if err == nil {
        return string(ret), nil
    } else {
        return "", err
    }
}

// Convert a string encoding from ShiftJIS to UTF-8
func FromShiftJIS(str string) (string, error) {
    return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewDecoder())
}

func main(){
  doc, _ := goquery.NewDocument("https://twitter.com/tokoya_x")
  doc.Find(".tweet-text").Each(func (_ int, s *goquery.Selection) {
    fmt.Println(s.Text())
  })
}
