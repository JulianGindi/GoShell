package main

import (
  "fmt"
  "bufio"
  "os"
)

type ShellCommand struct {
  command string
  arguments []string
}

type FileLocation struct {
  currentPath string
  pathHistory []string
}

func main() {
  fmt.Println("Welcome to GoShell")

  // Create an instance of fileLocation
  locations := new(FileLocation)

  // Creating a channel to communicate with the active shell session
  messages := make(chan string)
  go Reader(messages)

  // Blocking until we get a message
  <-messages
}

func Reader(messages chan string) {
  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print(">> ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
  }
}

func cd(loc *FileLocation, newLocation string) {
  loc.currentPath = newLocation
  loc.pathHistory = append(loc.pathHistory, newLocation)
}

// func (c *ShellCommand) executeCommand() {

// }
