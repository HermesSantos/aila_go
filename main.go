package main

import (
  "fmt"
  "os/exec"
)

func main () {
  cmd := exec.Command("git", "diff")
  output, err := cmd.Output()

  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  fmt.Println(string(output))
}
