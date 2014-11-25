package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func check(e error, filename string) {
  if e != nil {
    fmt.Println("No such file:", filename)
  }
}

func main() {
  filename := os.Args[1]
  dat, err := ioutil.ReadFile(filename)
  check(err, filename)
  fmt.Print(string(dat))
}
