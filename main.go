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

func main() {
  fmt.Println("Welcome to GoShell")

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

// func (c *ShellCommand) executeCommand() {

// }
