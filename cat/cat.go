package cat

import (
  "fmt"
  "io/ioutil"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func Cat(filename string) {
  dat, err := ioutil.ReadFile(filename)
  check(err)
  fmt.Print(string(dat))
}
