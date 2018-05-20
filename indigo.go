package main

import (
  "fmt"
  "os"
)

const version string = "0.1"
const help string = "\nIndigo v" + version + "\n\n\nArgument List:\n\n  new FILENAME          Start a new process\n  end FILENAME          End a process\n  console FILENAME      View a process's input console\n"

func main() {
    if len(os.Args) < 2 {
        fmt.Println(help)
        return
    }

    arg := os.Args[1]
    fmt.Println(arg)
}
