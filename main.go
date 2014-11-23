package main

import (
  "fmt"
  "bufio"
  "os"
  "os/exec"
  s "strings"
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

  // Creating a channel to send commands to the command dispatcher
  commands := make(chan ShellCommand)
  messages := make(chan string)

  // Create a go routine that handles dispatching commands
  go commandDispatcher(commands, messages)

  Reader(commands, messages)
}

func Reader(commands chan ShellCommand, messages chan string) {
  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print(">> ")
    text, _ := reader.ReadString('\n')
    splitText := s.Split(s.TrimSpace(text), " ")
    command := splitText[0]
    args := splitText[1:]
    currentCommand := ShellCommand{command, args}
    commands <- currentCommand
    <-messages
  }
}

func cd(loc *FileLocation, newLocation string) {
  loc.currentPath = newLocation
  loc.pathHistory = append(loc.pathHistory, newLocation)
}

func (c *ShellCommand) executeCommand(messages chan string) {
  command := exec.Command(c.command)
  commandOut, err := command.Output()
  if err != nil {
    fmt.Println("Invalid Command")
    return
  }
  fmt.Println(string(commandOut))
  messages <- "done"
}

func commandDispatcher(commands chan ShellCommand, messages chan string) {
  for {
    select {
    case command := <-commands:
      command.executeCommand(messages)
    }
  }
}
