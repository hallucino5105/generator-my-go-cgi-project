package garg

import (
  "os"

  "github.com/jessevdk/go-flags"
)

type MyOptions struct {
  Verbose bool `short:"v" long:"verbose" required:"false" description:"Show verbose debug information"`
}

var GlobalOptions *MyOptions = &MyOptions{}

func ParseArg() {
  _, err := flags.Parse(GlobalOptions)
  if err != nil {
    os.Exit(1)
  }
}
