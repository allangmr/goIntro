package main

import "fmt";

func main(){
  var publisher, writer, artist, title string
  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  var year, pageNumber int16
  year = 1997
  pageNumber = 14
  var grade float32
  grade = 6.5
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "year", year, "number of pages", pageNumber, "grade", grade)
  title = "Epic Vol.1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  year = 2013
  pageNumber = 160
  grade = 9.0
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "year", year, "number of pages", pageNumber, "grade", grade)
}