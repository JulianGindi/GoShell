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

  // Create an instance of fileLocation
  // locations := new(FileLocation)

  // Creating a channel to send commands to command dispatcher
  commands := make(chan ShellCommand)
  // messages := make(chan string)

  // Create a go routine that handles dispatching commands
  go commandDispatcher(commands)

  Reader(commands)
}

func Reader(commands chan ShellCommand) {
  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print(">> ")
    text, _ := reader.ReadString('\n')
    splitText := s.Split(s.TrimSpace(text), " ")
    command := splitText[0]
    args := splitText[1:]
    currentCommand := ShellCommand{command, args}
    commands <- currentCommand
    fmt.Println(args)
  }
}

func cd(loc *FileLocation, newLocation string) {
  loc.currentPath = newLocation
  loc.pathHistory = append(loc.pathHistory, newLocation)
}

func (c *ShellCommand) executeCommand() {
  command := exec.Command(c.command)
  commandOut, err := command.Output()
  if err != nil {
    panic(err)
  }
  fmt.Println(string(commandOut))
}

func commandDispatcher(commands chan ShellCommand) {
  for {
    select {
    case command := <-commands:
      command.executeCommand()
    }
  }
}
